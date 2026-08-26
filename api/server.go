package api

import (
	"encoding/json"
	"net/http"
	"retirement53/model"
	"retirement53/retirement"
)

type Server struct{ svc *retirement.Service }

func New(s *retirement.Service) *Server { return &Server{svc: s} }
func (s *Server) Handler() http.Handler {
	m := http.NewServeMux()
	m.HandleFunc("/health", s.health)
	m.HandleFunc("/users", s.users)
	m.HandleFunc("/cases/confirm", s.confirm)
	return m
}
func (s *Server) health(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}
func (s *Server) users(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	var u model.User
	if json.NewDecoder(r.Body).Decode(&u) != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if e := s.svc.RegisterUser(u); e != nil {
		http.Error(w, e.Error(), 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
func (s *Server) confirm(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if e := s.svc.Confirm(id, "api"); e != nil {
		http.Error(w, e.Error(), 400)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
