package api

import (
	"net/http/httptest"
	"path/filepath"
	"retirement53/retirement"
	"retirement53/storage"
	"testing"
)

func TestHealth(t *testing.T) {
	s, _ := storage.Open(filepath.Join(t.TempDir(), "x"))
	defer s.Close()
	w := httptest.NewRecorder()
	New(retirement.NewService(s)).Handler().ServeHTTP(w, httptest.NewRequest("GET", "/health", nil))
	if w.Code != 200 {
		t.Fatal(w.Code)
	}
}
