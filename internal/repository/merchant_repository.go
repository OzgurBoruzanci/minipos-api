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

func (repo *MerchantRepository) GetAllMerchants() ([]models.Merchant, error) {
	var merchants []models.Merchant
	err := repo.db.Find(&merchants).Error
	return merchants, err
}

func (repo *MerchantRepository) GetMerchantByID(id uint) (*models.Merchant, error) {
	var merchant models.Merchant
	err := repo.db.First(&merchant, id).Error
	return &merchant, err
}

func (repo *MerchantRepository) GetMerchantByAPIKey(apiKey string) (models.Merchant, error) {
	var merchant models.Merchant
	err := repo.db.Where("api_key = ?", apiKey).First(&merchant).Error
	return merchant, err
}

func (repo *TransactionRepository) GetTransactionByIdempotencyKey(key string) (*models.Transaction, error) {
	var transaction models.Transaction
	err := repo.db.Where("idempotency_key = ?", key).First(&transaction).Error
	return &transaction, err
}
