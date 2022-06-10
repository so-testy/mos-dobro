package myslq

import "time"

type User struct {
	ID         uint `gorm:"primary_key"`
	FamilyName string
	GivenName  string
	MiddleName string
	Birthdate  string
	Gender     uint
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time
}

type UserPhone struct {
	ID        uint `gorm:"primary_key"`
	UserID    uint
	Phone     string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type UserEmail struct {
	ID        uint `gorm:"primary_key"`
	UserID    string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
