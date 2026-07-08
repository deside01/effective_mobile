package config

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	HTTP HTTPConfig
	DB   DBConfig
}

type HTTPConfig struct {
	Addr         string
	ReadTimeout  int
	WriteTimeout int
	IdleTimeout  int
}

type DBConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	Name     string
	SSL      string
	Timezone string
}

func getenv(envName string) string {
	return strings.TrimSpace(os.Getenv(envName))
}

func Load(envName string) (*Config, error) {
	err := godotenv.Load(envName)
	if err != nil {
		log.Fatal(err)
	}

	httpCfg, err := loadHTTPConfig()
	if err != nil {
		return nil, fmt.Errorf("http config: %w", err)
	}

	dbCfg, err := loadDBConfig()
	if err != nil {
		return nil, fmt.Errorf("db config: %w", err)
	}

	return &Config{
		HTTP: httpCfg,
		DB:   dbCfg,
	}, nil
}
