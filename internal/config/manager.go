package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const (
	//defaultPath string = "/etc/lumi"
	defaultPath string = "./configs"
)

const (
	CoreConfig TypeOfConfig = "core"
)

type Manager interface {
	// Core operations
	Load() (*AppConfig, error)
	Save(config AppConfig) error

	// File operations
	Exists() bool
	CreateDefault() error

	// Validation
	Validate(config AppConfig) error
}

type FileConfigManager struct {
	configPath string
	configType TypeOfConfig
}

func New(configType TypeOfConfig) Manager {
	configPath := filepath.Join(defaultPath, string(configType)+".json")

	return &FileConfigManager{
		configType: configType,
		configPath: configPath,
	}
}

func (m *FileConfigManager) Load() (*AppConfig, error) {
	if !m.Exists() {
		return nil, fmt.Errorf("config file does not exist: %s", m.configPath)
	}

	data, err := os.ReadFile(m.configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read the config file: %w", err)
	}

	var config AppConfig
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to parse JSON config: %w", err)
	}

	if err := m.Validate(config); err != nil {
		return nil, fmt.Errorf("config validation failed: %w", err)
	}

	return &config, nil
}

func (m *FileConfigManager) Save(config AppConfig) error {
	if err := m.Validate(config); err != nil {
		return fmt.Errorf("config validation failed: %w", err)
	}

	dir := filepath.Dir(m.configPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create config directory: %w", err)
	}

	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal config to JSON: %w", err)
	}

	if err := os.WriteFile(m.configPath, data, 0644); err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}

	return nil
}

func (m *FileConfigManager) Validate(config AppConfig) error {

	if config.Server.Port <= 0 {
		return fmt.Errorf("invalid server port: %d", config.Server.Port)
	}

	if config.Server.Host == "" {
		return fmt.Errorf("server host cannot be empty")
	}

	return nil
}

func (m *FileConfigManager) Exists() bool {
	_, err := os.Stat(m.configPath)
	return err == nil
}

func (m *FileConfigManager) CreateDefault() error {
	return m.Save(DefaultCoreConfig)
}
