package my_sql_mogrations

import (
	"github.com/carprice-tech/migorm"
	"github.com/jinzhu/gorm"
	"time"
)

func init() {
	migorm.RegisterMigration(&migrationCreateUserEmailsTable{})
}

type migrationCreateUserEmailsTable struct{}

type createUserEmailsTable struct {
	ID        uint       `gorm:"primary_key"`
	UserID    string     `gorm:"not null; index:idx_user_id_deleted_at"`
	Email     string     `gorm:"not null"`
	CreatedAt time.Time  `gorm:"not null;type:datetime DEFAULT CURRENT_TIMESTAMP"`
	UpdatedAt time.Time  `gorm:"not null;type:datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
	DeletedAt *time.Time `gorm:"type:datetime DEFAULT NULL; index:idx_user_id_deleted_at"`
}

func (createUserEmailsTable) TableName() string {
	return "user_emails"
}

func (m *migrationCreateUserEmailsTable) Up(db *gorm.DB, log migorm.Logger) error {
	return db.AutoMigrate(createUserEmailsTable{}).Error
}

func (m *migrationCreateUserEmailsTable) Down(db *gorm.DB, log migorm.Logger) error {
	return db.DropTableIfExists("user_emails").Error
}
