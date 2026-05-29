package configs

import (
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yaml"
)

type Configuration struct {
	ServerPort string
	ServerHost string
	ServerMode string
	DBPort     uint16
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
}

func InitConfig() Configuration {
	var serverConfig Configuration

	config.WithOptions(config.ParseEnv)
	config.AddDriver(yaml.Driver)

	err := config.LoadFiles("configs/config.yaml")
	if err != nil {
		panic(err)
	}

	serverConfig.ServerPort = config.String("server.port", "80")
	serverConfig.ServerHost = config.String("server.host", "127.0.0.1")
	serverConfig.ServerMode = config.String("server.mode", "development")
	serverConfig.DBPort = uint16(config.Int("db.port", 5432))
	serverConfig.DBHost = config.String("db.host", "127.0.0.1")
	serverConfig.DBUser = config.String("db.user", "postgres")
	serverConfig.DBPassword = config.String("db.password", "mysecretpassword")
	serverConfig.DBName = config.String("db.name", "postgres")
	return serverConfig
}
