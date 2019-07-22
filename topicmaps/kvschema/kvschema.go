// Code generated by "kvschema"; DO NOT EDIT.

package kvschema

import (
	"github.com/google/note-maps/kv"
)

// Store provides entities, components, and indexes backed by a key-value
// store.
//
// Usage:
//
//   n, err := Store{Store: store}.NameComponent(0).Scan([]kv.Entity{7, 42})
//
type Store struct {
	kv.Store
	parent kv.Entity
}

func (s Store) Parent(e kv.Entity) *Store {
	s.parent = e
	return &s
}

// SetName sets the Name associated with e to v.
//
// Corresponding indexes are updated.
func (s *Store) SetName(e kv.Entity, v *Name) error {
	key := make(kv.Prefix, 8+2+8)
	s.parent.EncodeAt(key)
	NamePrefix.EncodeAt(key[8:])
	e.EncodeAt(key[10:])
	return s.Set(key, v.Encode())
}

// GetNameSlice returns a Name for each entity in es.
//
// If the underlying storage returns an empty value with no error for keys that
// do not exist, and Name.Decode() can decode an empty byte slice, then a
// query for entities that are not associated with a Name should return no
// errors.
func (s *Store) GetNameSlice(es []kv.Entity) ([]Name, error) {
	result := make([]Name, len(es))
	key := make(kv.Prefix, 8+2+8)
	s.parent.EncodeAt(key)
	NamePrefix.EncodeAt(key[8:])
	for i, e := range es {
		e.EncodeAt(key[10:])
		err := s.Get(key, (&result[i]).Decode)
		if err != nil {
			return nil, err
		}
	}
	return result, nil
}

// SetOccurrence sets the Occurrence associated with e to v.
//
// Corresponding indexes are updated.
func (s *Store) SetOccurrence(e kv.Entity, v *Occurrence) error {
	key := make(kv.Prefix, 8+2+8)
	s.parent.EncodeAt(key)
	OccurrencePrefix.EncodeAt(key[8:])
	e.EncodeAt(key[10:])
	return s.Set(key, v.Encode())
}

// GetOccurrenceSlice returns a Occurrence for each entity in es.
//
// If the underlying storage returns an empty value with no error for keys that
// do not exist, and Occurrence.Decode() can decode an empty byte slice, then a
// query for entities that are not associated with a Occurrence should return no
// errors.
func (s *Store) GetOccurrenceSlice(es []kv.Entity) ([]Occurrence, error) {
	result := make([]Occurrence, len(es))
	key := make(kv.Prefix, 8+2+8)
	s.parent.EncodeAt(key)
	OccurrencePrefix.EncodeAt(key[8:])
	for i, e := range es {
		e.EncodeAt(key[10:])
		err := s.Get(key, (&result[i]).Decode)
		if err != nil {
			return nil, err
		}
	}
	return result, nil
}

// SetTopicMapInfo sets the TopicMapInfo associated with e to v.
//
// Corresponding indexes are updated.
func (s *Store) SetTopicMapInfo(e kv.Entity, v *TopicMapInfo) error {
	key := make(kv.Prefix, 8+2+8)
	s.parent.EncodeAt(key)
	TopicMapInfoPrefix.EncodeAt(key[8:])
	e.EncodeAt(key[10:])
	return s.Set(key, v.Encode())
}

// GetTopicMapInfoSlice returns a TopicMapInfo for each entity in es.
//
// If the underlying storage returns an empty value with no error for keys that
// do not exist, and TopicMapInfo.Decode() can decode an empty byte slice, then a
// query for entities that are not associated with a TopicMapInfo should return no
// errors.
func (s *Store) GetTopicMapInfoSlice(es []kv.Entity) ([]TopicMapInfo, error) {
	result := make([]TopicMapInfo, len(es))
	key := make(kv.Prefix, 8+2+8)
	s.parent.EncodeAt(key)
	TopicMapInfoPrefix.EncodeAt(key[8:])
	for i, e := range es {
		e.EncodeAt(key[10:])
		err := s.Get(key, (&result[i]).Decode)
		if err != nil {
			return nil, err
		}
	}
	return result, nil
}
