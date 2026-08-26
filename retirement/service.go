package retirement

import (
	"fmt"
	"retirement53/model"
	"retirement53/storage"
	"time"
)

type Service struct {
	store *storage.Store
	now   func() time.Time
}

func NewService(s *storage.Store) *Service { return &Service{store: s, now: time.Now} }
func (s *Service) RegisterUser(u model.User) error {
	if e := model.ValidateUser(u); e != nil {
		return e
	}
	return s.store.SaveUser(u)
}
func (s *Service) AddRecord(r model.Record) error {
	if e := model.ValidateRecord(r); e != nil {
		return e
	}
	if _, e := s.store.LoadUser(r.UserID); e != nil {
		return fmt.Errorf("user: %w", e)
	}
	return s.store.SaveRecord(r)
}
func (s *Service) OpenCase(id, user string, expected int) error {
	c := model.RetirementCase{ID: id, UserID: user, Status: "open", Expected: expected, UpdatedAt: s.now()}
	if e := model.ValidateCase(c); e != nil {
		return e
	}
	return s.store.SaveCase(c)
}
func (s *Service) Confirm(id, actor string) error {
	c, e := s.store.LoadCase(id)
	if e != nil {
		return e
	}
	if c.Status == "closed" {
		return fmt.Errorf("case closed")
	}
	c.Confirmations++
	if c.Confirmations >= c.Expected {
		c.Status = "confirmed"
	}
	c.UpdatedAt = s.now()
	time.Sleep(2 * time.Millisecond)
	return s.store.SaveCase(c)
}
func (s *Service) CloseCase(id string) error {
	c, e := s.store.LoadCase(id)
	if e != nil {
		return e
	}
	if c.Confirmations < c.Expected {
		return fmt.Errorf("insufficient confirmations")
	}
	c.Status = "closed"
	c.UpdatedAt = s.now()
	return s.store.SaveCase(c)
}
func (s *Service) Case(id string) (model.RetirementCase, error) { return s.store.LoadCase(id) }
