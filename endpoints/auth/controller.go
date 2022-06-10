package auth_controller

import (
	"fmt"
	"github.com/gorilla/mux"
	"mos-dobro/config"
	"mos-dobro/internal/auth"
	"net/http"
)

type Controller struct {
	cfg    config.Config
	router *mux.Router
	authS  auth.Service
}

// NewAuthController - функция создания нового Auth контроллера
func NewAuthController(cfg config.Config, router *mux.Router, authS auth.Service) *Controller {
	return &Controller{
		cfg:    cfg,
		router: router,
		authS:  authS,
	}
}

// Init - метод инициализирующий Auth контроллер
func (c *Controller) Init() {
	c.router.HandleFunc(c.getAuthPath("userRegistration"), c.UserRegistration).Methods(http.MethodPost)
	c.router.HandleFunc(c.getAuthPath("authenticateByPhone"), c.AuthenticateByPhone).Methods(http.MethodPost)
	c.router.HandleFunc(c.getAuthPath("getTokenByPhoneCode"), c.GetTokenByPhoneCode).Methods(http.MethodGet)
}

// getAuthPath - метод получения пути Auth роута
func (c *Controller) getAuthPath(path string) string {
	return fmt.Sprintf("/my-sql/%s", path)
}
