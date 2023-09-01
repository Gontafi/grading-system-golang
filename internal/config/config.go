package config

import (
	"github.com/spf13/viper"
	"log"
	"os"
	"time"
)

type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
}

type RedisConfig struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

type ServerConfig struct {
	Port         string        `yaml:"port"`
	ReadTimeout  time.Duration `yaml:"readTimeout"`
	WriteTimeout time.Duration `yaml:"writeTimeout"`
	IdleTimeout  time.Duration `yaml:"idle_timeout"`
}

type Config struct {
	Database DatabaseConfig `yaml:"database"`
	Redis    RedisConfig    `yaml:"redis"`
	Server   ServerConfig   `yaml:"server"`
}

func LoadConfig(configPath string) (config *Config, err error) {
	if configPath == "" {
		log.Println("Config path is not set")
		return &Config{}, nil
	}
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exists : %s", err)
	}
	viper.AddConfigPath(configPath)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	err = viper.ReadInConfig()
	if err != nil {
		log.Println("failed to read config, Error:", err)
		return &Config{}, err
	}

	if err := viper.Unmarshal(&config); err != nil {
		return &Config{}, err
	}

	return config, nil
}
