package report

import (
	"encoding/json"
	"retirement53/analysis"
	"retirement53/model"
)

type Report struct {
	User       model.User
	Total      float64
	Categories map[string]float64
	Forecast   []float64
}

func Build(u model.User, rs []model.Record) Report {
	return Report{User: u, Total: analysis.Total(rs), Categories: analysis.ByCategory(rs), Forecast: analysis.Forecast(analysis.Total(rs), 5)}
}
func Encode(r Report) ([]byte, error) { return json.MarshalIndent(r, "", "  ") }
func Risk(r Report) string {
	if r.Total < 10000 {
		return "high"
	}
	if r.Total < 50000 {
		return "medium"
	}
	return "low"
}
