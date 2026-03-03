package repository

import (
	"errors"
	"minipos-api/internal/models"

	"gorm.io/gorm"
)

type TransactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (repo *TransactionRepository) ProcessPayment(transaction *models.Transaction) error {

	return repo.db.Transaction(func(tx *gorm.DB) error {
		var merchant models.Merchant
		if err := tx.First(&merchant, transaction.MerchantID).Error; err != nil {
			return errors.New("magaza bulunamadi")
		}
		if transaction.Type == "PAYMENT" {
			merchant.Balance += transaction.Amount

		} else if transaction.Type == "REFUND" {
			if merchant.Balance < transaction.Amount {
				return errors.New("iade icin yetersiz bakiye")
			}
			merchant.Balance -= transaction.Amount
		} else {
			return errors.New("gecersiz islem tipi: sadece PAYMENT veya REFUND olabilir")
		}
		if err := tx.Save(&merchant).Error; err != nil {
			return err
		}
		if err := tx.Create(transaction).Error; err != nil {
			return err
		}
		return nil
	})
}

func (repo *TransactionRepository) GetTransactionsByMerchant(merchantID uint, page int, limit int) ([]models.Transaction, error) {
	var transactions []models.Transaction
	offset := (page - 1) * limit
	err := repo.db.Where("merchant_id = ?", merchantID).
		Order("created_at desc").
		Offset(offset).
		Limit(limit).
		Find(&transactions).Error

	return transactions, err
}
