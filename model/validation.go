package model

import "fmt"

func ValidateUser(u User) error {
	if u.ID == "" || u.Name == "" {
		return fmt.Errorf("user identity required")
	}
	if u.BirthYear < 1900 || u.RetirementAge < 40 {
		return fmt.Errorf("invalid retirement profile")
	}
	return nil
}
func ValidateRecord(r Record) error {
	if r.ID == "" || r.UserID == "" {
		return fmt.Errorf("record identity required")
	}
	if r.Amount < 0 {
		return fmt.Errorf("amount cannot be negative")
	}
	if r.Category == "" {
		return fmt.Errorf("category required")
	}
	return nil
}
func ValidateCase(c RetirementCase) error {
	if c.ID == "" || c.UserID == "" {
		return fmt.Errorf("case identity required")
	}
	if c.Expected < 1 {
		return fmt.Errorf("expected confirmations required")
	}
	return nil
}
