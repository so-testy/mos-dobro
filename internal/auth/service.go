package auth

import (
	"context"
	"mos-dobro/internal/common"
)

type Service interface {
	UserRegistration(ctx context.Context, login, password string, role common.UserRole) (string, error)
	AuthenticateByPhone(ctx context.Context, phone string) error
	GetTokenByPhoneCode(ctx context.Context, smsCode string) (string, error)
}
