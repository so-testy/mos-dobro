package endpoints

import (
	"github.com/gorilla/mux"
	"mos-dobro/config"
	authController "mos-dobro/endpoints/auth"
	"mos-dobro/internal/auth"
)

type Controller interface {
	Init()
}

type ControllerManager interface {
	Auth() Controller
}

type controllerManager struct {
	authC Controller
}

// NewControllerManager - функция создания нового менеджера контроллеров
func NewControllerManager(cfg config.Config, router *mux.Router, auth auth.Service) ControllerManager {
	return &controllerManager{
		authC: authController.NewAuthController(cfg, router, auth),
	}
}

// Auth - метод получения Auth контроллера
func (cm *controllerManager) Auth() Controller {
	return cm.authC
}
