package analysis

import (
	"retirement53/model"
	"strings"
	"time"
)

func FilterUser(rs []model.Record, id string) []model.Record {
	out := []model.Record{}
	for _, r := range rs {
		if r.UserID == id {
			out = append(out, r)
		}
	}
	return out
}
func FilterCategory(rs []model.Record, cat string) []model.Record {
	out := []model.Record{}
	for _, r := range rs {
		if strings.EqualFold(r.Category, cat) {
			out = append(out, r)
		}
	}
	return out
}
func Between(rs []model.Record, start, end time.Time) []model.Record {
	out := []model.Record{}
	for _, r := range rs {
		if !r.CreatedAt.Before(start) && r.CreatedAt.Before(end) {
			out = append(out, r)
		}
	}
	return out
}
func SortByAmount(rs []model.Record) []model.Record {
	out := append([]model.Record{}, rs...)
	for i := 0; i < len(out); i++ {
		for j := i + 1; j < len(out); j++ {
			if out[j].Amount > out[i].Amount {
				out[i], out[j] = out[j], out[i]
			}
		}
	}
	return out
}
