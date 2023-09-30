package models

import (
	"time"

	"gorm.io/gorm"
)

type Event struct {
	ID          int            `gorm:"primaryKey" json:"id"`
	Name        string         `gorm:"not null" json:"name"`
	Description string         `gorm:"not null" json:"description"`
	Date        time.Time      `gorm:"not null" json:"date"`
	CreatedAt   time.Time      `gorm:"autoCreateTime" json:"-"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime" json:"-"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}
