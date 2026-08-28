package model

import "testing"

func TestUserEligibility(t *testing.T) {
	u := NewUser("u", "A", 1970, 53)
	if !u.Eligible(2023) {
		t.Fatal()
	}
}
