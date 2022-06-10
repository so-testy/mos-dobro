package auth

import (
	"context"
	"regexp"
)

// AuthenticateByPhone - метод аутентификации пользователя по номеру телефона
func (s ServiceImpl) AuthenticateByPhone(ctx context.Context, phone string) error {
	// Валидируем данные
	// TODO: Перенести в хелперс
	isValid, err := regexp.MatchString(`(7)[0-9]{10}$`, phone)
	if err != nil {
		return err
	}
	if !isValid {
		// TODO: Невалидный номер телефона
	}

	// Проверяем, есть ли у нас пользователь с таким номером телефона
	userID, err := s.repo.GetUserIDByPhone(phone)
	if err != nil {
		return err
	}
	if userID == nil || *userID == 0 {
		// TODO: Такой пользователь не существует
	}

	// Отправляем смс на номер телефона пользователя

	return nil
}
