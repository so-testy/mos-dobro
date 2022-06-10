package my_sql_migrator

import (
	"github.com/carprice-tech/migorm"
	"github.com/jinzhu/gorm"
	"mos-dobro/config"
	"path"
	"runtime"

	_ "github.com/go-sql-driver/mysql"
	_ "mos-dobro/migrations/my-sql"
)

func RunMigrations(cfg config.Config) error {
	// Подключаемся к БД
	db, err := gorm.Open("mysql", cfg.DB[config.DatabaseMySQL].GetDSN())
	if err != nil {
		return err
	}
	defer db.Close()

	// Создаем мигратор
	migrater := migorm.NewMigrater(db)

	// Определяем источник миграций
	_, file, _, _ := runtime.Caller(0)
	curDir := path.Dir(file)
	migrater.Conf().MigrationsDir = curDir + "/../migrations"

	// Запускаем миграции
	migorm.Run(migrater)

	return nil
}
