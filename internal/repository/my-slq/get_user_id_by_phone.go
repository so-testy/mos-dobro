package myslq

import (
	"errors"
	"gorm.io/gorm"
	"mos-dobro/pkg/helpers"
)

// GetUserIDByPhone - метод получения ID пользователя по номеру телефона
func (r RepositoryImpl) GetUserIDByPhone(phone string) (*uint, error) {
	var userPhone *UserPhone
	if err := r.db.Select("user_id").Where("phone = ?", phone).Take(&userPhone).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		return nil, err
	}
	return helpers.UintToPtr(userPhone.UserID), nil
}
