package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"mos-dobro/config"
	"mos-dobro/endpoints"
	"mos-dobro/internal/auth"
	myslq "mos-dobro/internal/repository/my-slq"
	"net/http"
	"time"
)

func main() {
	// Получаем конфиг настройки
	cfg := config.GetConfig()

	// Инициализируем репозиторий
	repo, err := myslq.NewRepository(cfg)
	if err != nil {
		panic(err)
	}

	// Создаем роутер
	router := mux.NewRouter()

	// Создаем Auth сервис
	authS, err := auth.NewService(cfg, repo)
	if err != nil {
		panic(err)
	}

	// Создаем контроллеры
	controller := endpoints.NewControllerManager(cfg, router, authS)

	// Инициализируем необходимые нам контроллеры
	controller.Auth().Init()

	// Запускаем сервер
	srv := &http.Server{
		Handler:      router,
		Addr:         cfg.Srvc[config.ServiceDefault].GetAddress(),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("server start: ", srv.Addr)
	log.Fatal(srv.ListenAndServe())
}
