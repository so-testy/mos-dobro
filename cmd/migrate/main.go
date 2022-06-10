package main

import (
	mySqlMigrator "mos-dobro/cmd/migrate/my-sql"
	"mos-dobro/config"
)

func main() {
	// Получаем конфиг настройки
	cfg := config.GetConfig()

	// Запускаем MySQL миграции
	if err := mySqlMigrator.RunMigrations(cfg); err != nil {
		panic(err)
	}
}
