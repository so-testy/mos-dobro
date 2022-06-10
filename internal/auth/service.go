package auth

import (
	"context"
	"mos-dobro/internal/common"
)

type Service interface {
	UserRegistration(ctx context.Context, login, password string, role common.UserRole) error
	AuthenticateByLoginPassword(ctx context.Context, login, password string) (string, error)
}
