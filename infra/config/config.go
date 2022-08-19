package config

type AppConfig struct {
	Database DatabaseConfig
}

type DatabaseConfig struct {
	DSN    string
	Driver string
}
