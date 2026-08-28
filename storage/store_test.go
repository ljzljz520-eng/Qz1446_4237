package storage

import (
	"path/filepath"
	"retirement53/model"
	"testing"
)

func TestPersistenceSurvivesReopen(t *testing.T) {
	p := filepath.Join(t.TempDir(), "x.db")
	s, e := Open(p)
	if e != nil {
		t.Fatal(e)
	}
	s.SaveUser(model.User{ID: "u", Name: "A", BirthYear: 1970, RetirementAge: 53})
	s.Close()
	s, e = Open(p)
	if e != nil {
		t.Fatal(e)
	}
	defer s.Close()
	if _, e = s.LoadUser("u"); e != nil {
		t.Fatal(e)
	}
}
