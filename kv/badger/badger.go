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

// Package badger providers a Badger-backed implementation of kv.Txn.
package badger

import (
	"fmt"
	"io"
	"log"

	"github.com/dgraph-io/badger"
	"github.com/google/note-maps/kv"
)

var (
	entitySequenceKey = []byte{0}
)

// DB holds some kv-specific state in addition to mixing in a badger.DB.
type DB struct {
	*badger.DB
	a *kv.Allocer
}

// DefaultOptions returns a recommended default Options value for a database
// rooted at dir.
func DefaultOptions(dir string) badger.Options {
	return badger.DefaultOptions(dir)
}

// Open creates a new DB with the given options.
func Open(opt badger.Options) (*DB, error) {
	bdb, err := badger.Open(opt)
	if err != nil {
		return nil, err
	}

	db := &DB{DB: bdb}
	db.a = kv.NewAllocer(db, []byte{0})
	return db, nil
}

func (db *DB) Dump(w io.Writer) {
	txn := db.NewTransaction(false)
	defer txn.Discard()
	opts := badger.DefaultIteratorOptions
	iter := txn.NewIterator(opts)
	defer iter.Close()
	count := 0
	for iter.Seek([]byte{0}); iter.Valid(); iter.Next() {
		item := iter.Item()
		key := item.Key()
		value, err := item.ValueCopy(nil)
		if err != nil {
			fmt.Fprintf(w, "%x\t%v", key, err)
			break
		}
		fmt.Fprintf(w, "%x\t%x\n", key, value)
		count++
	}
	fmt.Fprintf(w, "%v keys\n", count)
}

// Close releases unallocated Entity values and closes the database.
func (db *DB) Close() error {
	if db == nil {
		return nil
	}
	if db.a != nil {
		db.a.Save()
	}
	return db.DB.Close()
}

// NewTxn creates a new kv.Txn.
func (db *DB) NewTxn(update bool) kv.TxnCommitDiscarder {
	btxn := db.DB.NewTransaction(update)
	return txn{db: db, tx: btxn}
}

type txn struct {
	db *DB
	tx *badger.Txn
}

func (s txn) Alloc() (kv.Entity, error) {
	return s.db.a.Alloc()
}

func (s txn) Set(key, value []byte) error { return s.tx.Set(key, value) }

func (s txn) Get(key []byte, f func([]byte) error) error {
	item, err := s.tx.Get(key)
	if err == badger.ErrKeyNotFound {
		return f(nil)
	} else if err != nil {
		return err
	} else {
		return item.Value(f)
	}
}

func (s txn) PrefixIterator(prefix []byte) kv.Iterator {
	opts := badger.DefaultIteratorOptions
	opts.Prefix = prefix
	return iterator{
		s.tx.NewIterator(opts),
		prefix,
	}
}

func (s txn) Commit() error {
	if err := s.tx.Commit(); err != nil {
		return err
	}
	if err := s.db.Sync(); err != nil {
		// An error from Sync is important, but it does not indicate that the
		// commit failed.
		log.Println("txn.Commit: db.Sync:", err)
	}
	return nil
}

func (s txn) Discard() {
	s.tx.Discard()
}

type iterator struct {
	*badger.Iterator
	prefix []byte
}

func (i iterator) Seek(key []byte) { i.Iterator.Seek(append(i.prefix, key...)) }

func (i iterator) Key() []byte { return i.Item().Key()[len(i.prefix):] }

func (i iterator) Value(f func([]byte) error) error { return i.Item().Value(f) }

func (i iterator) Discard() { i.Close() }
