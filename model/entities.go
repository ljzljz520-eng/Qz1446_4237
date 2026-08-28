package model

import "time"

type Record struct {
	ID        string
	UserID    string
	Amount    float64
	Category  string
	CreatedAt time.Time
}
type User struct {
	ID            string
	Name          string
	BirthYear     int
	RetirementAge int
}
type Event struct {
	ID      string
	UserID  string
	Kind    string
	At      time.Time
	Details string
}
type Audit struct {
	ID     string
	UserID string
	Actor  string
	Action string
	At     time.Time
}
type RetirementCase struct {
	ID            string
	UserID        string
	Status        string
	Confirmations int
	Expected      int
	UpdatedAt     time.Time
}

func NewUser(id, name string, birth, age int) User {
	return User{ID: id, Name: name, BirthYear: birth, RetirementAge: age}
}
func (u User) Eligible(year int) bool { return year-u.BirthYear >= u.RetirementAge }
