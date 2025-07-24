package controllers

import (
	"go-test/models"
	"go-test/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTransactionHistory(c *gin.Context) {
	histories, err := models.GetTransactionHistories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Failed to get transaction histories",
		})
		return
	}

	c.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "Success to get transaction histories",
		Results: histories,
	})
}

func CreateTransaction(c *gin.Context) {
	var t models.TransactionRequest

	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Invalid request",
		})
		return
	}

	userID := c.GetInt("userId")

	if t.TransactionType != "in" && t.TransactionType != "out" {
		c.JSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Invalid transaction type (must be 'in' or 'out')",
		})
		return
	}

	if t.TransactionType == "out" {
		currentStock, err := models.CheckStock(t.ProductID)
		if err != nil || currentStock < t.Quantity {
			c.JSON(http.StatusBadRequest, utils.Response{
				Success: false,
				Message: "Insufficient stock",
			})
			return
		}
	}

	createdTransaction, err := models.CreateTransaction(t, userID)
	if err != nil {
		log.Println("CreateTransaction Error:", err)
		c.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Failed create transaction",
		})
		return
	}

	c.JSON(http.StatusCreated, utils.Response{
		Success: true,
		Message: "Transaction successful",
		Results: createdTransaction,
	})
}
