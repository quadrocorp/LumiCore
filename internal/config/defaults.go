package config

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
