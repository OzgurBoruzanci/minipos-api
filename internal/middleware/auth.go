package middleware

import (
	"net/http"

	"minipos-api/internal/repository"

	"github.com/gin-gonic/gin"
)

func APIKeyAuth(repo *repository.MerchantRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.GetHeader("X-API-Key")
		if apiKey == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "API Key eksik! Islem reddedildi."})
			return
		}
		merchant, err := repo.GetMerchantByAPIKey(apiKey)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Gecersiz API Key! Yetkisiz erisim."})
			return
		}
		c.Set("merchant_id", merchant.ID)
		c.Next()
	}
}
