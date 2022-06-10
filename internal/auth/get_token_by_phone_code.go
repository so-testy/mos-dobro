package auth

import "context"

// GetTokenByPhoneCode - метод получения токена по номеру телефона пользователя
func (s ServiceImpl) GetTokenByPhoneCode(ctx context.Context, smsCode string) (string, error) {

	return "", nil
}
