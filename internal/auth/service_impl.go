package auth

import (
	"mos-dobro/config"
	myslq "mos-dobro/internal/repository/my-slq"
)

type ServiceImpl struct {
	cfg  config.Config
	repo myslq.Repository
}

// NewService - функция создания нового Auth сервиса
func NewService(cfg config.Config, repo myslq.Repository) (Service, error) {
	return &ServiceImpl{
		cfg:  cfg,
		repo: repo,
	}, nil
}
