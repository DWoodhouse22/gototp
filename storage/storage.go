package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

type Account struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Group  string `json:"group"`
	Secret string `json:"secret"`
}

type Config struct {
	Accounts     map[string]Account `json:"accounts"`
	CurrentGroup string             `json:"current_group"`
}

const ownerReadWrite = 0o600

func SaveAccount(name, secret, group string) error {
	cfg, err := loadConfig()
	if err != nil {
		return err
	}
	if cfg.Accounts == nil {
		cfg.Accounts = make(map[string]Account)
	}

	for _, act := range cfg.Accounts {
		if act.Name == name && act.Group == group {
			return fmt.Errorf("account '%s' already exists in group '%s'", name, group)
		}
	}

	id := uuid.NewString()
	act := Account{
		ID:     id,
		Name:   name,
		Secret: secret,
		Group:  group,
	}
	cfg.Accounts[id] = act
	return saveConfig(cfg)
}

func FindAccounts(name, group string) ([]Account, error) {
	cfg, err := loadConfig()
	if err != nil {
		return nil, err
	}

	var out []Account
	for _, act := range cfg.Accounts {
		if act.Name != name {
			continue
		}
		if group != "" && act.Group != group {
			continue
		}
		out = append(out, act)
	}
	return out, nil
}

func SetCurrentGroup(group string) error {
	cfg, err := loadConfig()
	if err != nil {
		return err
	}

	cfg.CurrentGroup = group
	return saveConfig(cfg)
}

func GetCurrentGroup() (string, error) {
	cfg, err := loadConfig()
	if err != nil {
		return "", err
	}

	return cfg.CurrentGroup, nil
}

func ListAccounts() ([]Account, error) {
	cfg, err := loadConfig()
	if err != nil {
		return nil, err
	}

	accounts := make([]Account, 0, len(cfg.Accounts))
	for _, act := range cfg.Accounts {
		accounts = append(accounts, act)
	}
	return accounts, nil
}

func RemoveAccount(a Account) error {
	cfg, err := loadConfig()
	if err != nil {
		return err
	}

	if _, ok := cfg.Accounts[a.ID]; !ok {
		return fmt.Errorf("account does not exist")
	}

	delete(cfg.Accounts, a.ID)
	return nil
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
