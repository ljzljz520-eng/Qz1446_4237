package storage

import "retirement53/model"

func (s *Store) SaveRecord(v model.Record) error { return s.put("records", v.ID, v) }
func (s *Store) LoadRecord(id string) (model.Record, error) {
	var v model.Record
	e := s.get("records", id, &v)
	return v, e
}
func (s *Store) SaveUser(v model.User) error { return s.put("users", v.ID, v) }
func (s *Store) LoadUser(id string) (model.User, error) {
	var v model.User
	e := s.get("users", id, &v)
	return v, e
}
func (s *Store) SaveEvent(v model.Event) error         { return s.put("events", v.ID, v) }
func (s *Store) SaveAudit(v model.Audit) error         { return s.put("audits", v.ID, v) }
func (s *Store) SaveCase(v model.RetirementCase) error { return s.put("cases", v.ID, v) }
func (s *Store) LoadCase(id string) (model.RetirementCase, error) {
	var v model.RetirementCase
	e := s.get("cases", id, &v)
	return v, e
}
