package storage

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct {
	Secret string `json:"secret"`
}

const ownerReadWrite = 0o600

func getPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(home, ".totp_2fa.json"), nil
}

func SaveSecret(secret string) error {
	cfg := Config{Secret: secret}
	data, err := json.Marshal(cfg)
	if err != nil {
		return err
	}

	path, err := getPath()
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, ownerReadWrite)
}

func LoadSecret() (string, error) {
	path, err := getPath()
	if err != nil {
		return "", err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return "", err
	}

	return cfg.Secret, nil
}
