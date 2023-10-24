package util

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	DATABASE_TYPE 		string `mapstructure:"DATABASE_TYPE"`
	DATABASE_HOST 		string `mapstructure:"DATABASE_HOST"`
	DATABASE_PORT 		string `mapstructure:"DATABASE_PORT"`
	DATABASE_USER 		string `mapstructure:"DATABASE_USER"`
	DATABASE_PASS 		string `mapstructure:"DATABASE_PASS"`
	DATABASE_NAME 		string `mapstructure:"DATABASE_NAME"`
	DATABASE_SSL_MODE 	string `mapstructure:"DATABASE_SSL_MODE"`
	ServerAddress 		string `mapstructure:"SERVER_ADDRESS"`
	SessionType 		string `mapstructure:"SESSION_TYPE"`
	// cookies
	CookieName 	 		string `mapstructure:"COOKIE_NAME"`
	CookieLifetime 		string `mapstructure:"COOKIE_LIFETIME"`
	CookiePersist  		string `mapstructure:"COOKIE_PERSIST"`
	CookieSecure 		string `mapstructure:"COOKIE_SECURE"`
	CookieDomain 		string `mapstructure:"COOKIE_DOMAIN"`
	Port 				string `mapstructure:"PORT"`
	RootPath 			string
}

//loadconfig reads config from file or env variables
func LoadConfig() (*Config, error) {
	var config *Config
	path, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}
	config.RootPath = path
	return config, nil
}

func BuildDSN(c *Config) string {
	var dsn string

	switch c.DATABASE_TYPE {
	case "postgres", "postgresql":
		dsn = fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=%s timezone=UTC connect_timeout=5",
			c.DATABASE_HOST,
			c.DATABASE_PORT,
			c.DATABASE_USER,
			c.DATABASE_NAME, 
			c.DATABASE_SSL_MODE, 
		)

		if c.DATABASE_PASS != "" {
			dsn = fmt.Sprintf("%s password=%s", dsn, c.DATABASE_PASS)
		}

	default:
	}

	return dsn
}
