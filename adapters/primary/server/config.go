package server

import (
	"fmt"
	"time"
)

type Config struct {
	Address           string
	Domain            string
	WriteTimeout      time.Duration
	ReadTimeout       time.Duration
	MaxMessageBytes   int64
	MaxRecipients     int
	AllowInsecureAuth bool
}

func NewConfig(port string) *Config {
	addr := fmt.Sprintf("0.0.0.0:%s", port)
	return &Config{
		Address:           addr,
		Domain:            "simplified",
		WriteTimeout:      10 * time.Second,
		ReadTimeout:       10 * time.Second,
		MaxMessageBytes:   1024 * 1024,
		MaxRecipients:     50,
		AllowInsecureAuth: true,
	}
}
