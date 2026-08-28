package config

import "os"

type Config struct {
	DBPath string
	Addr   string
}

func Load() Config {
	p := os.Getenv("RETIREMENT_DB")
	if p == "" {
		p = "retirement.db"
	}
	a := os.Getenv("RETIREMENT_ADDR")
	if a == "" {
		a = ":8080"
	}
	return Config{DBPath: p, Addr: a}
}
