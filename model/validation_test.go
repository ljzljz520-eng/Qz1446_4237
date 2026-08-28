package model

import "testing"

func TestValidation(t *testing.T) {
	if ValidateUser(User{}) == nil {
		t.Fatal()
	}
	if ValidateRecord(Record{ID: "r", UserID: "u", Category: "x", Amount: -1}) == nil {
		t.Fatal()
	}
}
