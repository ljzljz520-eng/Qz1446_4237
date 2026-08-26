package analysis

import "retirement53/model"

type Metrics struct {
	Income, Expense, Savings, Rate float64
	Count                          int
}

func CalculateMetrics(rs []model.Record) Metrics {
	m := Metrics{Count: len(rs)}
	for _, r := range rs {
		if r.Amount >= 0 {
			if r.Category == "income" || r.Category == "pension" {
				m.Income += r.Amount
			} else {
				m.Expense += r.Amount
			}
		}
	}
	m.Savings = m.Income - m.Expense
	if m.Income > 0 {
		m.Rate = m.Savings / m.Income
	}
	return m
}
func MonthlyAverage(rs []model.Record) float64 {
	if len(rs) == 0 {
		return 0
	}
	return Total(rs) / float64(len(rs))
}
func ProjectInflation(v float64, years int, rate float64) []float64 {
	out := []float64{}
	for i := 0; i < years; i++ {
		out = append(out, v)
		v *= 1 + rate
	}
	return out
}
func NeedsReserve(m Metrics, months int) float64 {
	if m.Expense <= 0 {
		return 0
	}
	return m.Expense * float64(months)
}
