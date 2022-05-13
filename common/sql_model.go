package common

import "time"

type SQLModel struct {
	ID        int        `json:"-" gorm:"column:id;AUTO_INCREMENT;not_null;"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at;"`
	Status    int        `json:"-" gorm:"column:status;default:1;"`
}

