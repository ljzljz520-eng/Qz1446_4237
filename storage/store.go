package storage

import (
	"encoding/json"
	"go.etcd.io/bbolt"
	"os"
)

type Store struct{ db *bbolt.DB }

var buckets = []string{"records", "users", "events", "audits", "cases"}

func Open(path string) (*Store, error) {
	db, e := bbolt.Open(path, 0600, nil)
	if e != nil {
		return nil, e
	}
	s := &Store{db: db}
	e = db.Update(func(tx *bbolt.Tx) error {
		for _, b := range buckets {
			if _, x := tx.CreateBucketIfNotExists([]byte(b)); x != nil {
				return x
			}
		}
		return nil
	})
	if e != nil {
		db.Close()
		return nil, e
	}
	return s, nil
}
func (s *Store) Close() error {
	if s == nil || s.db == nil {
		return nil
	}
	return s.db.Close()
}
func (s *Store) put(bucket, key string, v any) error {
	data, e := json.Marshal(v)
	if e != nil {
		return e
	}
	return s.db.Update(func(tx *bbolt.Tx) error { return tx.Bucket([]byte(bucket)).Put([]byte(key), data) })
}
func (s *Store) get(bucket, key string, out any) error {
	return s.db.View(func(tx *bbolt.Tx) error {
		v := tx.Bucket([]byte(bucket)).Get([]byte(key))
		if v == nil {
			return os.ErrNotExist
		}
		return json.Unmarshal(v, out)
	})
}
