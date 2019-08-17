// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package kv provides some useful abstractions over local key-value storage.
//
//   go get github.com/google/note-maps/kv/...
//
// The model implemented by kv maps entities, which are like identifiers,
// to component values, which can be any Go type. Entity is an alias for
// uint64, and components are defined by kvschema, a code generator. The code
// generator looks for types that define the Encoder and Decoder interfaces
// from this package and produces strongly typed code for storing and
// retrieving instances of those types as values.
//
// Package kv also supports indexing. If a component value type, in addition to
// implementing Encoder and Decoder, also has one or more index methods, the
// generated code will also support looking up entities or loading entities in
// order according to each index. An index method must: have a name that starts
// with "Index", receive no arguments, and return a slice of a type that also
// implements Encoder and Decoder.
//
// Examples are included in the "examples" subdirectory.
//
// If `go generate` doesn't produce a kvschema.go file, or the resulting
// kvschema.go file does not include support for all the types you've defined,
// try `kvschema -v` to find out why.
package kv

import (
	"encoding/binary"
	"encoding/json"
	"log"
	"sort"
	"sync/atomic"
	"time"
)

// DB represents the functions a database connection should implement to be
// convenient for code that uses this package.
type DB interface {
	// NewTxn returns a new TxnCommitDiscarder that optionally supports updates.
	NewTxn(update bool) TxnCommitDiscarder

	// Close releases any resources held by this DB and closes the connection.
	Close() error
}

// Txn represents the functions a key-value store transaction must implement in
// order to be used as a backing transaction in this package.
type Txn interface {
	// Alloc should never return the same Entity value twice until the space of
	// possible Entity values is exhausted.
	//
	// Alloc cannot be implemented through Get and Set operations on the Txn
	// interface itself because independent concurrent transactions require
	// mutually unique Entity values, and the Txn interface maybe implemented
	// by a transaction type.
	Alloc() (Entity, error)

	// Set stores key and value in the underlying key-value store.
	Set(key, value []byte) error

	// Get finds the value associated with key in the underlying key-value store
	// and passes it to f.
	//
	// If the key does not an exist, this is not an error: Get may or may not
	// pass an empty slice to f.
	//
	// In any case, if f returns an error, then Get must also return an error.
	Get(key []byte, f func([]byte) error) error

	// PrefixIterator returns an iterator over all key-value pairs with keys
	// matching the given prefix.
	//
	// The initial state of the PrefixIterator is not valid: use or Next() or
	// Seek() to move the iterator to a valid key-value pair.
	//
	// The resulting iterator considers all valid keys as relative to the given
	// prefix, so for prefix {1,2} an underlying key {1,2,3,4} will be visible
	// through this iterator as merely {3,4}.
	PrefixIterator(prefix []byte) Iterator
}

// Iterator supports iteration over key-value pairs.
type Iterator interface {
	// Iterators must be discarded when no longer in use.
	Discarder

	// Seek moves the iterator to the key-value pair that matches the given key.
	//
	// If there is no such key-value pair, Seek moves to the item with first key
	// after the given key.
	Seek(key []byte)

	// Next moves to the iterator to the next key-value pair.
	Next()

	// Valid returns true if the iterator is at a valid key-value pair.
	Valid() bool

	// Key returns the key of the iterator's current key-value pair.
	//
	// May panic if Valid() returns false.
	Key() []byte

	// Value calls f with the value of the iterator's current key-value pair.
	//
	// May panic if Valid() returns false.
	Value(f func([]byte) error) error
}

// Discarder provides the Discard method.
type Discarder interface {
	Discard()
}

// TxnDiscarder combines the Txn and Discarder interfaces.
//
// A typical use of TxnDiscarder might be for a transaction that does not
// support mutations.
type TxnDiscarder interface {
	Txn
	Discarder
}

// TxnCommitDiscarder adds a Commit method to the Txn and Discarder interfaces.
//
// A typical use of TxnCommitDiscarder is to track a transaction that supports
// mutations.
type TxnCommitDiscarder interface {
	TxnDiscarder

	// Commit commits the mutations performed through this Txn to the backing store.
	Commit() error
}

// IndexCursor describes a location within an index, to track the position of
// an iterator within that index.
type IndexCursor struct {
	Key    []byte
	Offset int
}

// Partitioned combines a Txn with a partition identified by an Entity in order
// to provide useful common functions.
type Partitioned struct {
	Txn
	Partition Entity
}

// AllComponentEntities returns the first n entities that have values
// associated with c, beginning with the first entity greater than or equal to
// *start.
//
// A nil start value will be interpreted as a pointer to zero.
//
// A value of n less than or equal to zero will be interpretted as the largest
// possible value.
func (t Partitioned) AllComponentEntities(c Component, start *Entity, n int) (es []Entity, err error) {
	prefix := make(Prefix, 8+2)
	t.Partition.EncodeAt(prefix)
	c.EncodeAt(prefix[8:])
	iter := t.PrefixIterator(prefix[:10])
	defer iter.Discard()
	var actualStart Entity
	if start == nil || *start == 0 {
		actualStart = 1
	} else {
		actualStart = *start
	}
	bstart := actualStart.Encode()
	for iter.Seek(bstart); iter.Valid() && (n <= 0 || len(es) < n); iter.Next() {
		var e Entity
		e.Decode(iter.Key())
		es = append(es, e)
	}
	if start != nil && len(es) > 0 {
		*start = es[len(es)-1] + 1
	}
	return
}

// EntitiesByComponentIndex returns entities with c values ordered by their ix
// values.
//
// Reading begins at cursor, and ends when the length of the returned Entity
// slice is less than n. When reading is not complete, cursor is updated such
// that using it in a subequent call to EntitiesByComponentIndex would return
// the next n entities.
func (s Partitioned) EntitiesByComponentIndex(c, ix Component, cursor *IndexCursor, n int) (es []Entity, err error) {
	key := make(Prefix, 8+2+8+2)
	s.Partition.EncodeAt(key)
	c.EncodeAt(key[8:])
	Entity(0).EncodeAt(key[10:])
	ix.EncodeAt(key[18:])
	iter := s.PrefixIterator(key)
	defer iter.Discard()
	iter.Seek(cursor.Key)
	if !iter.Valid() {
		return
	}
	var buf EntitySlice
	if err = iter.Value(buf.Decode); err != nil {
		return
	}
	if cursor.Offset < len(buf) {
		es = append(es, buf[cursor.Offset:]...)
		if len(es) >= n {
			cursor.Offset += n
			if len(es) > n {
				es = es[:n]
			}
			return
		}
	}
	for iter.Next(); iter.Valid(); iter.Next() {
		if err = iter.Value(buf.Decode); err != nil {
			return
		}
		es = append(es, buf...)
		cursor.Key = append(cursor.Key[0:0], iter.Key()...)
		if len(es) >= n {
			cursor.Offset = len(buf) - (len(es) - n)
			if len(es) > n {
				es = es[:n]
			}
			return
		}
	}
	cursor.Offset = len(buf)
	return
}

// Allocer provides a generic implementation of the Txn.Alloc function.
//
// A single Allocer must be shared by all Txn values derived from the same
// database. If each Txn uses a unique Allocer, then they may produce duplicate
// values, which would violate the Txn.Alloc contract.
type Allocer struct {
	db   DB
	key  []byte
	last uint64
}

// NewAllocer returns a new Allocer.
//
// There should be no more than one Allocer per key per DB at any time.
func NewAllocer(db DB, key []byte) *Allocer {
	a := &Allocer{db: db, key: key}
	txn := db.NewTxn(false)
	defer txn.Discard()
	var last Entity
	txn.Get(key, func(bs []byte) error {
		if len(bs) > 0 {
			return last.Decode(bs)
		}
		return nil
	})
	a.last = uint64(last)
	now := uint64(time.Now().UnixNano())
	if a.last < now {
		a.last = now
	}
	return a
}

// Alloc uses an atomic counter to produce unique uint64 values.
func (a Allocer) Alloc() (Entity, error) {
	return Entity(atomic.AddUint64(&a.last, 1)), nil
}

// Save stores the last value returned by Allocer so that later calls to
// NewAllocer for the same DB and key will not produce any values that
// duplicate values returned by this one.
func (a Allocer) Save() error {
	txn := a.db.NewTxn(true)
	defer txn.Discard()
	if err := txn.Set(a.key, Entity(a.last).Encode()); err != nil {
		return err
	}
	return txn.Commit()
}

// Encoder is an interface implemented by any type that is to be stored in the
// key or value of a key-value pair.
type Encoder interface {
	Encode() []byte
}

// Decoder is an interface implemented by any type that is to be retrieved from
// the key or value of a key-value pair.
type Decoder interface {
	Decode(src []byte) error
}

// Prefix is a convenience type for constructing keys through concatenation.
type Prefix []byte

// ConcatEntity creates a new Prefix that contains p followed by e.
func (p Prefix) ConcatEntity(e Entity) Prefix {
	b := make([]byte, len(p)+8)
	copy(b, p)
	e.EncodeAt(b[len(p):])
	return b
}

// ConcatEntityComponent creates a new Prefix that contains p followed by e and
// c.
func (p Prefix) ConcatEntityComponent(e Entity, c Component) Prefix {
	b := make([]byte, len(p)+8+2)
	copy(b, p)
	e.EncodeAt(b[len(p):])
	c.EncodeAt(b[len(p)+8:])
	return b
}

// ConcatEntityComponentBytes creates a new Prefix that contains p followed by
// e, c, and bs.
func (p Prefix) ConcatEntityComponentBytes(e Entity, c Component, bs []byte) Prefix {
	b := make([]byte, len(p)+8+2+len(bs))
	copy(b, p)
	e.EncodeAt(b[len(p):])
	c.EncodeAt(b[len(p)+8:])
	copy(b[len(p)+8+2:], bs)
	return b
}

// AppendComponent appends c to p and returns the result.
func (p Prefix) AppendComponent(c Component) Prefix {
	return append(p, c.Encode()...)
}

// Component is a hard-coded and globally unique identifier for a component
// type.
//
// Components are typically hard-coded constants.
type Component uint16

// EncodeAt encodes e into the first two bytes of dst and panics if len(dst) <
// 2.
func (c Component) EncodeAt(dst []byte) {
	binary.BigEndian.PutUint16(dst, uint16(c))
}

// Encode encodes c into a new slice of two bytes.
func (c Component) Encode() []byte {
	var bs [2]byte
	c.EncodeAt(bs[:])
	return bs[:]
}

// Entity is an identifier that can be associated with Go values via
// Components, and
//
// Entities are typically created through Txn.Alloc().
type Entity uint64

// EncodeAt encodes e into the first eight bytes of dst and panics if len(dst)
// < 8.
func (e Entity) EncodeAt(dst []byte) {
	binary.BigEndian.PutUint64(dst, uint64(e))
}

// Encode encodes e into a new slice of eight bytes.
func (e Entity) Encode() []byte {
	var bs [8]byte
	e.EncodeAt(bs[:])
	return bs[:]
}

// Decode decodes the first eight bytes of src into e.
func (e *Entity) Decode(src []byte) error {
	*e = Entity(binary.BigEndian.Uint64(src))
	return nil
}

// EntitySlice implements sorting and searching for slices of Entity as well
// as sort order preserving insertion and removal operations.
type EntitySlice []Entity

// Len returns len(es).
//
// Len exists only to implement sort.Interface.
func (es EntitySlice) Len() int { return len(es) }

// Less returns true if and only if es[a] < es[b].
//
// Less exists only to implement sort.Interface.
func (es EntitySlice) Less(a, b int) bool { return es[a] < es[b] }

// Swap swaps the values of es[a] and es[b].
//
// Swap exists only to implement sort.Interface.
func (es EntitySlice) Swap(a, b int) { es[a], es[b] = es[b], es[a] }

// Sort sorts the values of es in ascending order.
func (es EntitySlice) Sort() { sort.Sort(es) }

// Equal returns true if and only if the contents of es match the contents of
// o.
func (es EntitySlice) Equal(o EntitySlice) bool {
	if len(es) != len(o) {
		return false
	}
	for i := range es {
		if es[i] != o[i] {
			return false
		}
	}
	return true
}

// Search returns the index of the first element of es that is greater than or
// equal to e.
//
// In other words, if e is an element of es, then es[es.Search(e)] == e.
// However, if all elements in es are less than e, then es.Search(e) == len(e).
//
// If es is not sorted, the results are undefined.
func (es EntitySlice) Search(e Entity) int {
	return sort.Search(len(es), func(i int) bool { return es[i] >= e })
}

// Insert adds e to es if it is not already included without disrupting the
// sorted ordering of es, and returns true if and only if e was not already
// present.
//
// If es is not already sorted, the results are undefined.
func (es *EntitySlice) Insert(e Entity) bool {
	if i := es.Search(e); i < len(*es) {
		if (*es)[i] == e {
			return false
		}
		*es = append((*es)[:i+1], (*es)[i:]...)
		(*es)[i] = e
	} else {
		*es = append(*es, e)
	}
	return true
}

// Remove removes e from es if it is present without disrupting the sorted
// ordering of es, and returns true if and only if e was there to be removed.
//
// If es is not already sorted, the results are undefined.
func (es *EntitySlice) Remove(e Entity) bool {
	if i := es.Search(e); i < len(*es) && (*es)[i] == e {
		*es = append((*es)[:i], (*es)[i+1:]...)
		return true
	}
	return false
}

// Encode encodes es into a new slice of bytes.
func (es EntitySlice) Encode() []byte {
	bs := make([]byte, 8*len(es))
	for i, e := range es {
		e.EncodeAt(bs[i*8:])
	}
	return bs
}

// Decode decodes src into es.
func (es *EntitySlice) Decode(src []byte) error {
	ln := len(src) / 8
	if len(*es) < ln {
		*es = make([]Entity, ln)
	}
	for i := 0; i < ln; i++ {
		(*es)[i].Decode(src[i*8:])
	}
	return nil
}

// String is an alias for string that implements the Encoder and Decoder
// interfaces.
type String string

// Encode encodes s into a new slice of bytes.
func (s String) Encode() []byte { return []byte(s) }

// Decode decodes src into s.
func (s *String) Decode(src []byte) error {
	*s = String(src)
	return nil
}

// StringSlice is a slice of strings that implements the Encoder and Decoder
// interfaces.
type StringSlice []String

// Encode encodes s into a new slice of bytes.
func (ss StringSlice) Encode() []byte {
	bs, err := json.Marshal(ss)
	if err != nil {
		log.Println(err)
	}
	return bs
}

// Decode decodes src into s.
func (ss *StringSlice) Decode(src []byte) error {
	return json.Unmarshal(src, ss)
}
