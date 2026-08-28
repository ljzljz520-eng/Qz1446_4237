package retirement

import (
	"path/filepath"
	"retirement53/model"
	"retirement53/storage"
	"sync"
	"testing"
)

func TestWorkflowOne(t *testing.T) {
	s, _ := storage.Open(filepath.Join(t.TempDir(), "x"))
	defer s.Close()
	v := NewService(s)
	if e := v.RegisterAndRecord(model.User{ID: "u", Name: "N", BirthYear: 1970, RetirementAge: 53}, model.Record{ID: "r", UserID: "u", Category: "income", Amount: 20}); e != nil {
		t.Fatal(e)
	}
}
func TestWorkflowTwo(t *testing.T) {
	s, _ := storage.Open(filepath.Join(t.TempDir(), "x"))
	defer s.Close()
	v := NewService(s)
	v.RegisterUser(model.User{ID: "u", Name: "N", BirthYear: 1970, RetirementAge: 53})
	v.OpenCase("c", "u", 3)
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() { defer wg.Done(); _ = v.Confirm("c", "a") }()
	}
	wg.Wait()
	c, _ := v.Case("c")
	if c.Confirmations != 3 {
		t.Fatalf("confirmations=%d", c.Confirmations)
	}
}
func TestWorkflowThree(t *testing.T) {
	s, _ := storage.Open(filepath.Join(t.TempDir(), "x"))
	defer s.Close()
	v := NewService(s)
	v.RegisterUser(model.User{ID: "u", Name: "N", BirthYear: 1970, RetirementAge: 53})
	if e := v.AddRecord(model.Record{ID: "r", UserID: "u", Category: "expense", Amount: 10}); e != nil {
		t.Fatal(e)
	}
}
