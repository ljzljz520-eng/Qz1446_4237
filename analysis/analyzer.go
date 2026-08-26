package analysis

import (
	"retirement53/model"
	"sort"
)

func Total(records []model.Record) float64 {
	var n float64
	for _, r := range records {
		n += r.Amount
	}
	return n
}
func ByCategory(records []model.Record) map[string]float64 {
	m := map[string]float64{}
	for _, r := range records {
		m[r.Category] += r.Amount
	}
	return m
}
func Balance(income, expense float64) float64 { return income - expense }
func Forecast(monthly float64, years int) []float64 {
	out := make([]float64, 0, years)
	v := monthly
	for i := 0; i < years; i++ {
		out = append(out, v*12)
		v *= 1.02
	}
	return out
}
func TopCategories(m map[string]float64, n int) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool { return m[keys[i]] > m[keys[j]] })
	if n > len(keys) {
		n = len(keys)
	}
	return keys[:n]
}
