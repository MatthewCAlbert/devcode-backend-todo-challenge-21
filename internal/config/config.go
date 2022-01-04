package config

import (
	"os"
)

type Config struct {
	AppEnv        string
	ServerPort    string
	MySqlHost     string
	MySqlPort     string
	MySqlDatabase string
	MySqlUser     string
	MySqlPassword string
}

func Load() (*Config, error) {
	c := Config{}

	c.AppEnv = os.Getenv("APP_ENV")
	c.ServerPort = "3030"

	c.MySqlPort = "3306"
	c.MySqlHost = os.Getenv("MYSQL_HOST")
	c.MySqlDatabase = os.Getenv("MYSQL_DBNAME")
	c.MySqlUser = os.Getenv("MYSQL_USER")
	c.MySqlPassword = os.Getenv("MYSQL_PASSWORD")

	return &c, nil
}
