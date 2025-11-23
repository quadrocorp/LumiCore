package config

// Author == quadrocorp
// Version == v1.0-alpha || Photon
// NOTE: Photon is the alpha version of this project, and implies an MVP.

// This file contains variables with default configurations.
// It is planned to be hardcoded as a method of a failsafe:
// - If the config file is destroyed/lost/unavailable the config manager
// will automatically recreate configuration using the predefined default
// ensuring that Lumi would always start no matter what happens. (of course
// if file used to contain list of Docker containers, the default one will not.)

// Note: might need of another way to save information about containers, perhaps using
// multiple sources could prevent permanent loss in case of configuration file corruption.

var DefaultCoreConfig AppConfig = AppConfig{
	Server: ServerConfig{
		Host: "localhost",
		Port: 8080,
	},
	Database: DatabaseConfig{
		Host:     "localhost",
		Port:     5432,
		Name:     "lumicore",
		Username: "postgres",
		Password: "password",
	},
}
