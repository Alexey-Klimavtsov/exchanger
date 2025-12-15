package config

import "time"

type Config struct {
	Port        string
	CacheTTL    time.Duration
	RateLimit   int
	RateWindow  time.Duration
}

func Load() Config {
	return Config{
		Port:       ":8080",
		CacheTTL:   10 * time.Minute,
		RateLimit:  10,
		RateWindow: time.Minute,
	}
}
