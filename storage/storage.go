package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Config struct {
	Accounts map[string]string `json:"accounts"`
}

const ownerReadWrite = 0o600

func SaveAccount(name, secret string) error {
	cfg, err := loadConfig()
	if err != nil {
		return err
	}
	if cfg.Accounts == nil {
		cfg.Accounts = make(map[string]string)
	}

	cfg.Accounts[name] = secret
	return saveConfig(cfg)
}

func LoadAccount(name string) (string, error) {
	cfg, err := loadConfig()
	if err != nil {
		return "", err
	}

	secret, ok := cfg.Accounts[name]
	if !ok {
		return "", fmt.Errorf("account not found")
	}

	return secret, nil
}

func getPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(home, ".totp_2fa.json"), nil
}

func loadConfig() (Config, error) {
	path, err := getPath()
	if err != nil {
		return Config{}, err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return Config{}, nil
		}
		return Config{}, err
	}

	var cfg Config
	err = json.Unmarshal(data, &cfg)
	return cfg, err
}

func saveConfig(cfg Config) error {
	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}

	path, err := getPath()
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, ownerReadWrite)
}
