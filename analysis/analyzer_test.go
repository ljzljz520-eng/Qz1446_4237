package analysis

import (
	"retirement53/model"
	"testing"
)

func TestTotals(t *testing.T) {
	r := []model.Record{{Amount: 2}, {Amount: 3}}
	if Total(r) != 5 {
		t.Fatal()
	}
	if len(Forecast(1, 3)) != 3 {
		t.Fatal()
	}
}
