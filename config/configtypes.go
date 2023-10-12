package config

type DatabaseConfig struct {
	URL      string `mapstructure:"url"`
	DBDriver string `mapstructure:"dbDriver"`
}

type ServerConfig struct {
	Port string `mapstructure:"port"`
}

type Config struct {
	Database DatabaseConfig
	Server   ServerConfig
}
