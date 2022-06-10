package auth

import (
	"mos-dobro/config"
	"mos-dobro/internal/repository"
)

type ServiceImpl struct {
	cfg  config.Config
	repo repository.Repository
}

// NewService - функция создания нового Auth сервиса
func NewService(cfg config.Config, repo repository.Repository) (Service, error) {
	return &ServiceImpl{
		cfg:  cfg,
		repo: repo,
	}, nil
}
