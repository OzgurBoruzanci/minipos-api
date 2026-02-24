package repository

import (
	"minipos-api/internal/models"

	"gorm.io/gorm"
)

type MerchantRepository struct {
	db *gorm.DB
}

func NewMerchantRepository(db *gorm.DB) *MerchantRepository {
	return &MerchantRepository{db: db}
}

func (repo *MerchantRepository) CreateMerchant(merchant *models.Merchant) error {
	return repo.db.Create(merchant).Error
}
