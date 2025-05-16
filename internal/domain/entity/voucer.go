package entity

import "time"

type Voucher struct {
	ID          uint   `gorm:"primaryKey"`
	BrandID     uint   `gorm:"not null"`
	Name        string `gorm:"size:100;not null"`
	Description string `gorm:"size:255"`
	PointCost   int    `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
