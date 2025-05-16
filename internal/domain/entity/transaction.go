package entity

import "time"

type Transaction struct {
	ID         uint   `gorm:"primaryKey"`
	CustomerID uint   `gorm:"not null"`
	TotalPoint int    `gorm:"not null"`
	Status     string `gorm:"size:20;not null;default:SUCCESS"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Items      []TransactionItem `gorm:"foreignKey:TransactionID"`
}

type TransactionItem struct {
	ID            uint `gorm:"primaryKey"`
	TransactionID uint `gorm:"not null"`
	VoucherID     uint `gorm:"not null"`
	Quantity      int  `gorm:"not null"`
	PointCost     int  `gorm:"not null"`
	CreatedAt     time.Time
}
