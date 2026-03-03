package notifications

import (
	"fmt"
	"minipos-api/internal/models"
	"time"
)

var WebhookQueue = make(chan models.Transaction, 100)

func StartNotificationWorker() {
	go func() {
		fmt.Println("🚀 Arka plan bildirim işçisi (Worker) başlatıldı...")
		for tx := range WebhookQueue {
			fmt.Printf("📧 [Bildirim Gönderiliyor] Mağaza ID: %d, Tutar: %.2f TL...\n", tx.MerchantID, tx.Amount)
			time.Sleep(2 * time.Second)
			fmt.Printf("✅ [Bildirim Tamamlandı] İşlem ID: %d için bildirim başarıyla iletildi.\n", tx.ID)
		}
	}()
}
