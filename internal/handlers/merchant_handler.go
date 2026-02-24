package handlers

import (
	"net/http"

	"minipos-api/internal/models"
	"minipos-api/internal/repository"

	"github.com/gin-gonic/gin"
)

type MerchantHandler struct {
	repo *repository.MerchantRepository
}

func NewMerchantHandler(repo *repository.MerchantRepository) *MerchantHandler {
	return &MerchantHandler{repo: repo}
}

func (h *MerchantHandler) CreateMerchant(c *gin.Context) {
	var merchant models.Merchant
	if err := c.ShouldBindJSON(&merchant); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Gecersiz veri formati"})
		return
	}
	if err := h.repo.CreateMerchant(&merchant); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Veritabanina kaydedilemedi: " + err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "Magaza basariyla olusturuldu",
		"data":    merchant,
	})
}
