package repository

import (
	"mos-dobro/config"
	myslq "mos-dobro/internal/repository/my-slq"
)

type RepositoryManager interface {
	MySQL() myslq.Repository
}

type repositoryManager struct {
	cfg       config.Config
	mysqlRepo myslq.Repository
}

// NewRepositoryManager - функция создания менеджера репозиториев
func NewRepositoryManager(cfg config.Config) (RepositoryManager, error) {
	// Инициализируем MySQL репозиторий
	mysqlRepo, err := myslq.NewRepository(cfg)
	if err != nil {
		return nil, err
	}

	// Создаем менеджера репозиториев
	return &repositoryManager{
		cfg:       cfg,
		mysqlRepo: mysqlRepo,
	}, nil
}

// MySQL - метод получения MySQL репозитория
func (rm *repositoryManager) MySQL() myslq.Repository {
	return rm.mysqlRepo
}
