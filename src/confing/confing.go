package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	DataBase struct {
		User     string
		Password string
		DBName   string
		Host     string
		Port     int
		SSLMode  string
	}
	JWT struct {
		Secret string
		Expire int
	}
	Server struct {
		Port int
	}
}

var AppConfig Config

func LoadConfig() {

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".") // Ищем файл в корне проекта

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Ошибка загрузки конфига:", err)
	}

	err := viper.Unmarshal(&AppConfig) // Маппим в структуру
	if err != nil {
		log.Fatal("Ошибка парсинга конфига:", err)
	}
}
