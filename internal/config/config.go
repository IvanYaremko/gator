package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

const (
	fileName = ".gatorconfig.json"
)

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func (c *Config) SetUser(username string) error {
	c.CurrentUserName = username
	err := write(*c)
	if err != nil {
		return fmt.Errorf("setuser error: %w", err)
	}

	return nil
}

func ReadConfig() (Config, error) {
	filePath, err := getConfigFilePath()
	if err != nil {
		return Config{}, fmt.Errorf("error retrieving home directory: %w", err)
	}

	file, err := os.Open(filePath)
	if err != nil {
		return Config{}, fmt.Errorf("error retrieving config file: %w", err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return Config{}, fmt.Errorf("error io read all of file: %w", err)
	}

	config := Config{}
	if err := json.Unmarshal(data, &config); err != nil {
		return Config{}, fmt.Errorf("error unmarshalling config file: %w", err)
	}

	return config, nil
}

func getConfigFilePath() (string, error) {
	dir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("error retrieving home directory: %w", err)
	}

	return fmt.Sprintf("%s/%s", dir, fileName), nil
}

func write(cfg Config) error {
	bytes, err := json.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("func write marshal: %w", err)
	}
	filepath, err := getConfigFilePath()
	if err != nil {
		return fmt.Errorf("func write get file path: %w", err)
	}
	if err := os.WriteFile(filepath, bytes, 0666); err != nil {
		return fmt.Errorf("func write file error: %w", err)
	}

	return nil
}
