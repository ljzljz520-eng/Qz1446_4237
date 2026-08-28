package report

import (
	"fmt"
	"strings"
)

func Text(r Report) string {
	parts := []string{fmt.Sprintf("用户: %s", r.User.Name), fmt.Sprintf("总额: %.2f", r.Total), fmt.Sprintf("风险: %s", Risk(r))}
	for k, v := range r.Categories {
		parts = append(parts, fmt.Sprintf("%s=%.2f", k, v))
	}
	return strings.Join(parts, "\n")
}
func CSV(r Report) string {
	out := "category,amount\n"
	for k, v := range r.Categories {
		out += fmt.Sprintf("%s,%.2f\n", k, v)
	}
	return out
}
func IsSustainable(r Report) bool { return r.Total > 0 && Risk(r) != "high" }
func Yearly(r Report, year int) float64 {
	if year < 0 || year >= len(r.Forecast) {
		return 0
	}
	return r.Forecast[year]
}
