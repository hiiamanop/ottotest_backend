package entity

import "time"

type Customer struct {
	ID           uint   `gorm:"primaryKey"`
	Name         string `gorm:"size:100;not null"`
	Email        string `gorm:"size:100;unique;not null"`
	PointBalance int    `gorm:"not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
