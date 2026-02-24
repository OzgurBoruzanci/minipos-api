package database

import (
	"log"
	"minipos-api/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "host=localhost user=postgres password=gizlisifrem dbname=minipos port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Veritabanına bağlanılamadı: %v", err)
	}
	log.Println("Veritabanı bağlantısı başarılı!")
	err = DB.AutoMigrate(&models.Merchant{})
	if err != nil {
		log.Fatalf("Tablolar oluşturulamadı: %v", err)
	}
	log.Println("Veritabanı tabloları (Migration) başarıyla oluşturuldu!")
}
