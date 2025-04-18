package config

import (
	"encoding/json"
	"os"
)

const CONFIG_PATH = "./.gatorconfig.json"

type Config struct {
	DbUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() (Config, error) {
	var cfg Config

	file, err := os.Open(CONFIG_PATH)
	if err != nil {
		return cfg, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&cfg)

	if err != nil {
		return cfg, err
	}

	return cfg, nil
}

func (cfg *Config) SetUser(userName string) error {
	cfg.CurrentUserName = userName

	err := write(cfg)
	if err != nil {
		return err
	}

	return nil
}

func write(cfg *Config) error {
	file, err := os.Create(CONFIG_PATH)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	err = encoder.Encode(cfg)
	if err != nil {
		return err
	}

	return nil
}
