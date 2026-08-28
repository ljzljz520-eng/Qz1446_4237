package storage

import (
	"encoding/json"
	"go.etcd.io/bbolt"
	"retirement53/model"
)

func (s *Store) list(bucket string, out func([]byte) error) error {
	return s.db.View(func(tx *bbolt.Tx) error {
		return tx.Bucket([]byte(bucket)).ForEach(func(_, v []byte) error {
			if v == nil {
				return nil
			}
			return out(v)
		})
	})
}
func (s *Store) ListRecords(user string) ([]model.Record, error) {
	out := []model.Record{}
	e := s.list("records", func(v []byte) error {
		var r model.Record
		if e := json.Unmarshal(v, &r); e != nil {
			return e
		}
		if user == "" || r.UserID == user {
			out = append(out, r)
		}
		return nil
	})
	return out, e
}
func (s *Store) ListUsers() ([]model.User, error) {
	out := []model.User{}
	e := s.list("users", func(v []byte) error {
		var u model.User
		if e := json.Unmarshal(v, &u); e != nil {
			return e
		}
		out = append(out, u)
		return nil
	})
	return out, e
}
func (s *Store) DeleteRecord(id string) error {
	return s.db.Update(func(tx *bbolt.Tx) error { return tx.Bucket([]byte("records")).Delete([]byte(id)) })
}
func (s *Store) Count(bucket string) (int, error) {
	n := 0
	e := s.db.View(func(tx *bbolt.Tx) error {
		return tx.Bucket([]byte(bucket)).ForEach(func(_, v []byte) error {
			if v != nil {
				n++
			}
			return nil
		})
	})
	return n, e
}
