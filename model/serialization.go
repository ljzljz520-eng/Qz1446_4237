package model

import (
	"encoding/json"
	"time"
)

func EncodeRecord(r Record) ([]byte, error) { return json.Marshal(r) }
func DecodeRecord(b []byte) (Record, error) { var r Record; e := json.Unmarshal(b, &r); return r, e }
func EncodeUser(u User) ([]byte, error)     { return json.Marshal(u) }
func DecodeUser(b []byte) (User, error)     { var u User; e := json.Unmarshal(b, &u); return u, e }
func At(t time.Time) string                 { return t.UTC().Format(time.RFC3339) }
