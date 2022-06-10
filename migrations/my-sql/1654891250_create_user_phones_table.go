package my_sql_mogrations

import (
	"github.com/carprice-tech/migorm"
	"github.com/jinzhu/gorm"
	"time"
)

func init() {
	migorm.RegisterMigration(&migrationCreateUserPhonesTable{})
}

type migrationCreateUserPhonesTable struct{}

type createUserPhonesTable struct {
	ID        uint       `gorm:"primary_key"`
	UserID    uint       `gorm:"not null; index:idx_user_id_deleted_at"`
	Phone     string     `gorm:"not null"`
	CreatedAt time.Time  `gorm:"not null;type:datetime DEFAULT CURRENT_TIMESTAMP"`
	UpdatedAt time.Time  `gorm:"not null;type:datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
	DeletedAt *time.Time `gorm:"type:datetime DEFAULT NULL; index:idx_user_id_deleted_at"`
}

func (createUserPhonesTable) TableName() string {
	return "user_phones"
}

func (m *migrationCreateUserPhonesTable) Up(db *gorm.DB, log migorm.Logger) error {
	return db.AutoMigrate(createUserPhonesTable{}).Error
}

func (m *migrationCreateUserPhonesTable) Down(db *gorm.DB, log migorm.Logger) error {
	return db.DropTableIfExists("user_phones").Error
}
