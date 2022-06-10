package my_sql_mogrations

import (
	"github.com/carprice-tech/migorm"
	"github.com/jinzhu/gorm"
	"time"
)

func init() {
	migorm.RegisterMigration(&migrationCreateUserTable{})
}

type migrationCreateUserTable struct{}

type createUserTable struct {
	ID         uint   `gorm:"primary_key"`
	FamilyName string `gorm:"not null"`
	GivenName  string `gorm:"not null"`
	MiddleName string
	Birthdate  string
	Gender     uint
	CreatedAt  time.Time  `gorm:"not null;type:datetime DEFAULT CURRENT_TIMESTAMP"`
	UpdatedAt  time.Time  `gorm:"not null;type:datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
	DeletedAt  *time.Time `gorm:"type:datetime DEFAULT NULL; index:idx_login_deleted_at"`
}

func (createUserTable) TableName() string {
	return "users"
}

func (m *migrationCreateUserTable) Up(db *gorm.DB, log migorm.Logger) error {
	return db.AutoMigrate(createUserTable{}).Error
}

func (m *migrationCreateUserTable) Down(db *gorm.DB, log migorm.Logger) error {
	return db.DropTableIfExists("users").Error
}
