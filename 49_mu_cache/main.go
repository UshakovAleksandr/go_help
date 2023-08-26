package main

import (
	"fmt"
	"log"
	"net/http"

	"49_mu_cache/api"
	"49_mu_cache/db"
)

func main() {
	// Инициализация БД в памяти.
	storage := db.New()
	// Создание объекта App, использующего БД в памяти.
	app := api.New(storage)
	// Запуск сетевой службы и HTTP-сервера
	// на всех локальных IP-адресах на порту 80.
	fmt.Println("Starting server...")
	err := http.ListenAndServe(":8080", app.Router())
	if err != nil {
		log.Fatal(err)
	}
}
