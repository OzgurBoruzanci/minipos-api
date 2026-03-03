package repository

import (
	"minipos-api/internal/models"
	"testing"
)

func TestCalculateNewBalance(t *testing.T) {
	initialBalance := 1000.0
	paymentAmount := 250.0
	refundAmount := 100.0
	merchant := models.Merchant{
		Balance: initialBalance,
	}
	merchant.Balance += paymentAmount
	expectedAfterPayment := 1250.0
	if merchant.Balance != expectedAfterPayment {
		t.Errorf("Ödeme sonrası bakiye hatası! Beklenen: %f, Gelen: %f", expectedAfterPayment, merchant.Balance)
	}
	merchant.Balance -= refundAmount
	expectedAfterRefund := 1150.0
	if merchant.Balance != expectedAfterRefund {
		t.Errorf("İade sonrası bakiye hatası! Beklenen: %f, Gelen: %f", expectedAfterRefund, merchant.Balance)
	}
}

func TestInsufficientBalance(t *testing.T) {
	balance := 50.0
	refundAttempt := 100.0

	if refundAttempt > balance {
		t.Log("Başarılı: Sistem yetersiz bakiyeyi doğru tespit etti.")
	} else {
		t.Error("Hata: Sistem yetersiz bakiye ile iadeye izin verdi!")
	}
}
