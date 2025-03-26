// config/config.go
package config

import (
	"github.com/spf13/viper"
)

type Config struct {
    Server   ServerConfig
    Database DatabaseConfig
    JWT      JWTConfig
}

type ServerConfig struct {
    Port string
    Mode string // development, production
}

type DatabaseConfig struct {
    Host     string
    Port     string
    User     string
    Password string
    Name     string
}

type JWTConfig struct {
    Secret     string
    Expiration int // in hours
}

// LoadConfig reads configuration from file or environment variables
func LoadConfig(path string) (*Config, error) {
    viper.AddConfigPath(path)
    viper.SetConfigName("config")
    viper.SetConfigType("yaml")

    viper.AutomaticEnv() // override with env variables

    // Set defaults
    viper.SetDefault("server.port", "8080")
    viper.SetDefault("server.mode", "development")
    viper.SetDefault("jwt.expiration", 24)
    
    err := viper.ReadInConfig()
    if err != nil {
        return nil, err
    }

    var config Config
    err = viper.Unmarshal(&config)
    return &config, err
}