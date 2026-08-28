package report

import (
	"retirement53/model"
	"testing"
)

func TestReport(t *testing.T) {
	r := Build(model.User{Name: "N"}, []model.Record{{Amount: 10, Category: "income"}})
	if Risk(r) != "high" {
		t.Fatal()
	}
	if len(Text(r)) == 0 {
		t.Fatal()
	}
}
