package config

import (
	"github.com/spf13/viper"
)

const (
	defaultServerPort = "8080"
	defaultMySqlPort  = "3306"
)

type Config struct {
	ServerPort    string
	MySqlHost     string
	MySqlPort     string
	MySqlDatabase string
	MySqlUser     string
	MySqlPassword string
}

func Load(env string) (*Config, error) {
	c := Config{}

	v := viper.New()
	v.AddConfigPath("./configs/")
	v.SetConfigType("yaml")
	if env == "production" {
		v.SetConfigName("prod")
	} else {
		v.SetConfigName("dev")
	}
	v.AutomaticEnv()

	v.SetDefault("server_port", defaultServerPort)
	v.SetDefault("MYSQL_PORT", defaultMySqlPort)

	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	c.ServerPort = v.GetString("server_port")
	c.MySqlHost = v.GetString("MYSQL_HOST")
	c.MySqlPort = v.GetString("MYSQL_PORT")
	c.MySqlDatabase = v.GetString("MYSQL_DBNAME")
	c.MySqlUser = v.GetString("MYSQL_USER")
	c.MySqlPassword = v.GetString("MYSQL_PASSWORD")

	return &c, nil
}
