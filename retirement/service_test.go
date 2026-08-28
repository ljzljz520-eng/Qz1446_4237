package retirement

import (
	"path/filepath"
	"retirement53/model"
	"retirement53/storage"
	"testing"
)

func TestRecordFlow53(t *testing.T) {
	s, _ := storage.Open(filepath.Join(t.TempDir(), "x"))
	defer s.Close()
	svc := NewService(s)
	if e := svc.RegisterUser(model.User{ID: "u", Name: "退休53", BirthYear: 1970, RetirementAge: 53}); e != nil {
		t.Fatal(e)
	}
	if e := svc.AddRecord(model.Record{ID: "r", UserID: "u", Category: "pension", Amount: 100}); e != nil {
		t.Fatal(e)
	}
}
