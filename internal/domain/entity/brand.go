package entity

import "time"

type Brand struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"size:100;not null"`
	Description string `gorm:"size:255"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
