package config

// Author == quadrocorp
// Version == v1.0-alpha || Photon
// NOTE: Photon is the alpha version of this project, and implies an MVP.

// This file defines configuration types in LumiCore package.

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
