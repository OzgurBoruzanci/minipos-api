package handlers

import (
	"net/http"

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
