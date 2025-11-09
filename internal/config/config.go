package config

import (
	"encoding/json"
	"fmt"
	"os"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func (config *Config) SetUser(username string) error {
	config.CurrentUserName = username

	err := write(*config)
	if err != nil {
		return fmt.Errorf("couldn't write config to file: %w", err)
	}
	return nil
}

func Read() (Config, error) {
	path, err := getConfigFilePath()
	if err != nil {
		return Config{}, fmt.Errorf("error getting config file path: %w", err)
	}

	configData, err := os.ReadFile(path)
	if err != nil {
		return Config{}, fmt.Errorf("error reading config file: %w", err)
	}

	config := Config{}
	err = json.Unmarshal(configData, &config)
	if err != nil {
		return Config{}, fmt.Errorf("cannot unmarshal config data: %w", err)
	}

	return config, nil
}

func getConfigFilePath() (string, error) {
	dir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("cannot read home directory: %w", err)
	}
	return dir + "/" + configFileName, nil
}

func write(cfg Config) error {
	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return fmt.Errorf("couldn't marshal config data: %w", err)
	}

	filePath, err := getConfigFilePath()
	if err != nil {
		return fmt.Errorf("error getting config file path: %w", err)
	}

	err = os.WriteFile(filePath, data, 0644)
	if err != nil {
		return fmt.Errorf("couldn't write to file: %w", err)
	}

	return nil
}
