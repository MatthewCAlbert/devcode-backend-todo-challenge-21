package config

import (
	"fmt"

	"github.com/spf13/viper"
)

const (
	defaultServerPort = "8080"
	defaultMySqlPort  = "3306"
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

	v := viper.New()
	v.AddConfigPath("./configs/")
	v.SetConfigType("yaml")
	v.SetDefault("APP_ENV", "development")

	c.AppEnv = v.GetString("APP_ENV")

	if c.AppEnv == "production" {
		v.SetConfigName("prod")
	} else {
		v.SetConfigName("dev")
	}
	v.AutomaticEnv()

	v.SetDefault("server_port", defaultServerPort)
	v.SetDefault("MYSQL_PORT", defaultMySqlPort)

	if err := v.ReadInConfig(); err != nil {
		fmt.Println(err)
		return nil, err
	}

	c.AppEnv = v.GetString("APP_ENV")
	c.ServerPort = v.GetString("server_port")
	c.MySqlHost = v.GetString("MYSQL_HOST")
	c.MySqlPort = v.GetString("MYSQL_PORT")
	c.MySqlDatabase = v.GetString("MYSQL_DBNAME")
	c.MySqlUser = v.GetString("MYSQL_USER")
	c.MySqlPassword = v.GetString("MYSQL_PASSWORD")

	return &c, nil
}
