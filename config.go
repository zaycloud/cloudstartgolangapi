package main

import (
	"fmt"
	"os"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Database struct {
		File     string `yaml:"file" envconfig:"DB_FILE"`
		Username string `yaml:"sql-user" envconfig:"DB_USERNAME"`
		Password string `yaml:"sql-pass" envconfig:"DB_PASSWORD"`
		Database string `yaml:"sql-database" envconfig:"DB_DATABASE"`
		Server   string `yaml:"sql-server" envconfig:"DB_SERVER"`
		Port     int    `yaml:"sql-port" envconfig:"DB_PORT"`
	} `yaml:"database"`

	Redis struct {
		Server string `yaml:"server" envconfig:"REDIS_SERVER"`
	} `yaml:"redis"`

	SSO struct {
		ClientId        string `yaml:"clientId" envconfig:"SSO_CLIENTID"`
		ClientSecret    string `yaml:"clientSecret" envconfig:"SSO_CLIENTSECRET"`
		BaseAddress     string `yaml:"baseAddress" envconfig:"SSO_BASEADDRESS"`
		BackBaseAddress string `yaml:"backBaseAddress" envconfig:"SSO_BACKBASEADDRESS"`
	} `yaml:"sso"`
}

func readConfig(cfg *Config) {
	readFile(cfg)
	readEnv(cfg)
}

func readFile(cfg *Config) {
	fileName := "config.yml"
	s := os.Getenv("RUNENVIRONMENT")
	if len(s) > 0 {
		fileName = "config" + s + ".yml"
	}

	f, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Error opening config file: %v\n", err)
		return
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		fmt.Printf("Error decoding config file: %v\n", err)
	}
}

func readEnv(cfg *Config) {
	err := envconfig.Process("", cfg)
	if err != nil {
		fmt.Printf("Error processing environment variables: %v\n", err)
	}
}
