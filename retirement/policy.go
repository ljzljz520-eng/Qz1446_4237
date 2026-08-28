package retirement

import (
	"fmt"
	"retirement53/model"
)

func CheckAge(u model.User, current int) error {
	if !u.Eligible(current) {
		return fmt.Errorf("not eligible")
	}
	return nil
}
func RequiredReviewers(u model.User) int {
	if u.RetirementAge <= 53 {
		return 2
	}
	return 1
}
func CanArchive(c model.RetirementCase) bool {
	return c.Confirmations >= c.Expected && c.Status != "closed"
}
func Transition(c *model.RetirementCase, target string) error {
	if target == "closed" && !CanArchive(*c) {
		return fmt.Errorf("cannot archive")
	}
	if target != "open" && target != "confirmed" && target != "closed" {
		return fmt.Errorf("invalid status")
	}
	c.Status = target
	return nil
}
