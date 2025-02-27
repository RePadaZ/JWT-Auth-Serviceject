package main

import (
	config "JWT-Auth-Serviceject/src/confing"
	"JWT-Auth-Serviceject/src/database"
	"JWT-Auth-Serviceject/src/routes"
	"log"
)

func main() {

	// Загружаем конфиг файл
	config.LoadConfig()

	// Подключаемся к БД
	database.Connect()

	// Устанавливаем роутеры
	r := routes.SetupRouter()
	err := r.Run(":8080")
	if err != nil {
		log.Fatal("Ошибка При запуске сервера:", err)
	}

}
