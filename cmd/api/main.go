package main

import (
	"minipos-api/internal/handlers"
	"minipos-api/internal/middleware"
	"minipos-api/internal/repository"
	"minipos-api/pkg/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDB()
	merchantRepo := repository.NewMerchantRepository(database.DB)
	merchantHandler := handlers.NewMerchantHandler(merchantRepo)

	transactionRepo := repository.NewTransactionRepository(database.DB)
	transactionHandler := handlers.NewTransactionHandler(transactionRepo)

	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "MiniPOS API Ayakta ve Calisiyor",
		})
	})
	r.POST("/merchants", merchantHandler.CreateMerchant)
	r.GET("/merchants", merchantHandler.GetAllMerchants)
	r.GET("/merchants/:id", merchantHandler.GetMerchantByID)
	r.POST("/transactions", middleware.APIKeyAuth(merchantRepo), transactionHandler.CreateTransaction)
	r.GET("/merchants/:id/transactions", middleware.APIKeyAuth(merchantRepo), transactionHandler.GetTransactionsByMerchant)
	if err := r.Run(":8080"); err != nil {
		panic("Sunucu baslatilamadi: " + err.Error())
	}
}
