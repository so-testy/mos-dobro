package auth

import (
	"context"
	"mos-dobro/internal/common"
)

// UserRegistration - метод регистрирующий пользователя
func (s ServiceImpl) UserRegistration(ctx context.Context, login, password string, role common.UserRole) (string, error) {

	return "", nil
}
