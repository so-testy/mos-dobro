package myslq

type Repository interface {
	GetUserIDByPhone(phone string) (*uint, error)
}
