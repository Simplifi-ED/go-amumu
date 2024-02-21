package server

import (
	"os"
	"time"

	"gopkg.in/yaml.v3"
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

func NewConfig(configPath string) (*Config, error) {
	type alias struct {
		Address           string `yaml:"host"`
		Domain            string `yaml:"domain"`
		WriteTimeout      string `yaml:"writetimeout" default:"10"`
		ReadTimeout       string `yaml:"readtimeout" default:"10"`
		MaxMessageBytes   int64  `yaml:"maxMessagebytes" default:"1048576"`
		MaxRecipients     int    `yaml:"maxRecipients" default:"50"`
		AllowInsecureAuth bool   `yaml:"allowinsecureAuth" default:"true"`
	}
	smtpConfig := &alias{}

	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	d := yaml.NewDecoder(file)

	if err := d.Decode(&smtpConfig); err != nil {
		return nil, err
	}
	writetimeout, _ := time.ParseDuration(smtpConfig.WriteTimeout)
	readtimeout, _ := time.ParseDuration(smtpConfig.ReadTimeout)
	return &Config{
		Address:           smtpConfig.Address,
		Domain:            smtpConfig.Domain,
		WriteTimeout:      writetimeout * time.Second,
		ReadTimeout:       readtimeout * time.Second,
		MaxMessageBytes:   smtpConfig.MaxMessageBytes,
		MaxRecipients:     smtpConfig.MaxRecipients,
		AllowInsecureAuth: smtpConfig.AllowInsecureAuth,
	}, nil
}
