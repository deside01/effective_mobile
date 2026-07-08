package config

import (
	"fmt"
)

func (dbCfg *DBConfig) GetDSN() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=%s",
		dbCfg.Host,
		dbCfg.Port,
		dbCfg.User,
		dbCfg.Password,
		dbCfg.Name,
		dbCfg.SSL,
		dbCfg.Timezone,
	)
}

func loadDBConfig() (DBConfig, error) {
	dbCfg := DBConfig{
		User:     getenv("DB_USER"),
		Password: getenv("DB_PASSWORD"),
		Host:     getenv("DB_HOST"),
		Port:     getenv("DB_PORT"),
		Name:     getenv("DB_NAME"),
		SSL:      getenv("DB_SSL"),
		Timezone: getenv("DB_TIMEZONE"),
	}

	if err := validateDB(&dbCfg); err != nil {
		return DBConfig{}, err
	}

	return dbCfg, nil
}

func validateDB(cfg *DBConfig) error {
	if cfg.User == "" {
		return fmt.Errorf("DB_USER is required")
	}
	if cfg.Password == "" {
		return fmt.Errorf("DB_PASSWORD is required")
	}
	if cfg.Host == "" {
		return fmt.Errorf("DB_HOST is required")
	}
	if cfg.Port == "" {
		return fmt.Errorf("DB_PORT is required")
	}
	if cfg.Name == "" {
		return fmt.Errorf("DB_NAME is required")
	}

	if cfg.SSL == "" {
		cfg.SSL = "disable"
	}

	if cfg.Timezone == "" {
		cfg.Timezone = "UTC"
	}

	return nil
}
