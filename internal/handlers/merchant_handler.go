package handlers

import (
	"net/http"
	"strconv"

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

func (h *MerchantHandler) GetAllMerchants(c *gin.Context) {
	merchants, err := h.repo.GetAllMerchants()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Magazalar getirilemedi: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Magazalar basariyla getirildi",
		"data":    merchants,
	})
}

func (h *MerchantHandler) GetMerchantByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Gecersiz ID formatı"})
		return
	}
	merchant, err := h.repo.GetMerchantByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Magaza getirilemedi: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Magaza basariyla getirildi",
		"data":    merchant,
	})
}
