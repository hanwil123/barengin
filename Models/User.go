package Models

import "time"

type AuthUsers struct {
	ID        uint64 `gorm:"primaryKey" json:"id"`
	Userid    string `gorm:"foreignKey" json:"userid"`
	Email     string `gorm:"unique"json:"email"`
	Password  []byte
	Name      string `json:"name"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
