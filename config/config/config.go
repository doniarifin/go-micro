package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	// App struct {
	// 	Port int
	// } `mapstructure:"app"`

	Database struct {
		Host   string
		Port   string
		DBName string
		User   string
		Pass   string
	} `mapstructure:"database"`

	RabbitMQ struct {
		URL   string
		Queue string
	} `mapstructure:"rabbitmq"`
}

func Load() *Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../../data")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}

	return &cfg
}
