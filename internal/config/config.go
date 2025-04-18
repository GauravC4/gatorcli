package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/user"
)

const CONFIG_PATH = "./.gatorconfig.json"

type Config struct {
	DbUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() Config {
	var cfg Config

	file, err := os.Open(CONFIG_PATH)
	if err != nil {
		log.Fatal("error opening config file : ", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&cfg)

	if err != nil {
		log.Fatal("error decoding config file : ", err)
	}

	return cfg
}

func (cfg *Config) SetUser() {
	username := "user"

	currentUser, err := user.Current()

	if err != nil {
		fmt.Println("did no find username, assigning 'user' as default")
	} else {
		username = currentUser.Username
	}

	cfg.CurrentUserName = username

	err = write(cfg)
	if err != nil {
		log.Fatal("error writing to config file : ", err)
	}
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
