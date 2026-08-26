package main

import (
	"retirement53/analysis"
	"retirement53/model"
	"testing"
)

func TestEndToEnd(t *testing.T) {
	u := model.NewUser("u", "N", 1970, 53)
	if !u.Eligible(2023) {
		t.Fatal()
	}
	if analysis.Balance(100, 40) != 60 {
		t.Fatal()
	}
}
