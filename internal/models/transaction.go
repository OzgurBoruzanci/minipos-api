package models

import "time"

type Transaction struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	MerchantID uint      `json:"merchant_id" gorm:"not null"`
	Amount     float64   `json:"amount" gorm:"not null"`
	Type       string    `json:"type" gorm:"not null"`
	Status     string    `json:"status" gorm:"default:'SUCCESS'"`
	CreatedAt  time.Time `json:"created_at"`
}
