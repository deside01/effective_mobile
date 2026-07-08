package config

import (
	"fmt"
	"strconv"
)

func loadHTTPConfig() (httpCfg HTTPConfig, err error) {
	addr := getenv("ADDR")
	if addr == "" {
		return httpCfg, fmt.Errorf("ADDR is required")
	}

	readTt, err := parseTimeout("READ_TIMEOUT")
	if err != nil {
		return httpCfg, err
	}

	writeTt, err := parseTimeout("WRITE_TIMEOUT")
	if err != nil {
		return httpCfg, err
	}

	idleTt, err := parseTimeout("IDLE_TIMEOUT")
	if err != nil {
		return httpCfg, err
	}

	return HTTPConfig{
		Addr:         addr,
		ReadTimeout:  readTt,
		WriteTimeout: writeTt,
		IdleTimeout:  idleTt,
	}, nil
}

func parseTimeout(envName string) (int, error) {
	val, err := strconv.Atoi(getenv(envName))
	if err != nil {
		return 0, fmt.Errorf("parsing %s: %w", envName, err)
	}

	if val <= 0 {
		return 0, fmt.Errorf("%s must be positive, got %d", envName, val)
	}
	return val, nil
}
