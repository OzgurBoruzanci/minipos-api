package handlers

import (
	"net/http"
	"strconv"

	"minipos-api/internal/models"
	"minipos-api/internal/repository"

	"github.com/gin-gonic/gin"
)

type TransactionHandler struct {
	repo *repository.TransactionRepository
}

func NewTransactionHandler(repo *repository.TransactionRepository) *TransactionHandler {
	return &TransactionHandler{repo: repo}
}

func (h *TransactionHandler) CreateTransaction(c *gin.Context) {
	var transaction models.Transaction

	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Gecersiz veri formati"})
		return
	}

	if err := h.repo.ProcessPayment(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{
		"message": "Islem basariyla gerceklesti",
		"data":    transaction,
	})
}

func (h *TransactionHandler) GetTransactionsByMerchant(c *gin.Context) {
	idParam := c.Param("id")

	merchantID, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Gecersiz ID formati"})
		return
	}
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}
	transactions, err := h.repo.GetTransactionsByMerchant(uint(merchantID), page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Islemler getirilemedi: " + err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{
		"page":  page,
		"limit": limit,
		"data":  transactions,
	})
}
