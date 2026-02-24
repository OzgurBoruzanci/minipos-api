package models

type Merchant struct {
	ID       uint    `json:"id" gorm:"primaryKey"`
	Name     string  `json:"name" gorm:"not null"`
	ApiKey   string  `json:"api_key" gorm:"not null;unique"`
	Balance  float64 `json:"balance" gorm:"default:0.0"`
	IsActive bool    `json:"is_active" gorm:"default:true"`
}
