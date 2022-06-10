package myslq

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"mos-dobro/config"
)

type RepositoryImpl struct {
	db *gorm.DB
}

// NewRepository - функция создания нового MySQL репозитория
func NewRepository(cfg config.Config) (Repository, error) {
	// Подключаемся к БД
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       cfg.DB[config.DatabaseMySQL].GetDSN(),
		DefaultStringSize:         255,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    false,
		DontSupportRenameColumn:   true,
		DontSupportForShareClause: true,
	}))
	if err != nil {
		return nil, err
	}

	return &RepositoryImpl{db: db}, nil
}
