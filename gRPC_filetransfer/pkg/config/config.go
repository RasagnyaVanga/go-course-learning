package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

// Config holds all configuration values for the server.
// It can be populated from a YAML file and/or environment variables.
type Config struct {
	Server       ServerConfig       `yaml:"server"`
	FilesStorage FilesStorageConfig `yaml:"files_storage"`
}

// ServerConfig contains settings for the gRPC server.
type ServerConfig struct {
	// Port is the address the gRPC server listens on.
	// Example: ":50051"
	Port string `yaml:"port" `
}

// FilesStorageConfig defines where uploaded files will be stored.
type FilesStorageConfig struct {
	// Location is the directory path where uploaded files will be saved.
	// Example: "./uploads"
	UploadsLocation   string `yaml:"uploads_location" `
	DownloadsLocation string `yaml:"downloads_location" `
}

// LoadConfig reads configuration from a YAML file and environment variables.
// YAML provides default values, and environment variables (if set) override them.
func LoadConfig() (*Config, error) {
	cfg := &Config{}

	// 1. Read configuration from YAML file
	err := cleanenv.ReadConfig("./pkg/config/config.yml", cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to read config.yml: %w", err)
	}

	// 2. Override any fields using environment variables (if they exist)
	if err := cleanenv.ReadEnv(cfg); err != nil {
		return nil, fmt.Errorf("failed to read environment variables: %w", err)
	}

	return cfg, nil
}
