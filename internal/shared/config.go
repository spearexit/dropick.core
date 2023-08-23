package shared

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

type ConnectionInfos struct {
	User     string
	Password string
	Host     string
	Port     string
	Database string
}

type DatabaseConfig struct {
	Connection *ConnectionInfos
}

type ServerConfig struct {
	Port string
	Mode string
}

type ConfigStruct struct {
	Database *DatabaseConfig
	Server   *ServerConfig
}

var Config *ConfigStruct

func init() {
	if Config == nil {
		Config = &ConfigStruct{
			Database: &DatabaseConfig{
				Connection: &ConnectionInfos{
					User:     os.Getenv("DB_USER"),
					Password: os.Getenv("DB_PASSWORD"),
					Host:     os.Getenv("DB_HOST"),
					Port:     os.Getenv("DB_PORT"),
					Database: os.Getenv("DB_DATABASE"),
				},
			},
			Server: &ServerConfig{
				Port: os.Getenv("SERVER_PORT"),
				Mode: os.Getenv("SERVER_MODE"),
			},
		}
	}
}
