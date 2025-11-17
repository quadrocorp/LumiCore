package config

// Owned by quadrocorp.
// This file defines configuration types in LumiCore package.
// LumiCore is the core of Lumi, the Telegram Bot manager and Docker Orchestration tool
// with statistics features.

// File updated to the version: Photon
// NOTE: Photon is the alpha version of this project, and implies an MVP.

// AppConfig represents the main application configuration
type AppConfig struct {
	Server   ServerConfig   `json:"server"`
	Database DatabaseConfig `json:"database"`
}

type ServerConfig struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

type DatabaseConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type TypeOfConfig string

type Format string
