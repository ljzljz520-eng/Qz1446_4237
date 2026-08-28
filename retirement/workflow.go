package retirement

import (
	"fmt"
	"retirement53/model"
)

func (s *Service) RegisterAndRecord(u model.User, r model.Record) error {
	if e := s.RegisterUser(u); e != nil {
		return e
	}
	return s.AddRecord(r)
}
func (s *Service) ProcessCase(id string, actors []string) error {
	for _, a := range actors {
		if e := s.Confirm(id, a); e != nil {
			return e
		}
	}
	return s.CloseCase(id)
}
func (s *Service) Summary(id string) (string, error) {
	c, e := s.Case(id)
	if e != nil {
		return "", e
	}
	return fmt.Sprintf("%s:%d/%d", c.Status, c.Confirmations, c.Expected), nil
}
