package handler

import (
	"kilobite/helper"
	"kilobite/transaction"
	"kilobite/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

// parameter di uri
// tangkap parameter mapping input struct
// panggil service, input struct sebagai parameter
// service, berbekal campaign id bisa panggil repoxx
// repo mencari data transaction suatu campaign

type transactionHandler struct {
	service transaction.Service
}

func NewTransactionHandler(service transaction.Service) *transactionHandler {
	return &transactionHandler{service}
}

func (h *transactionHandler) GetCampaignTransactions(c *gin.Context) {
	var input transaction.GetCampaignTransactionsInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to get campaign's transaction", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)

	input.User = currentUser

	transactions, err := h.service.GetTransactionsByCampaignID(input)
	if err != nil {
		response := helper.APIResponse("Failed to get campaign's transactions", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Campaign's transactions", http.StatusOK, "success", transaction.FormatCampaignTransactions(transactions))
	c.JSON(http.StatusOK, response)
}

//GetUserTransactions
//Handler
//ambil nilai user dari wt/middleware
//service
//repo => ambil data transactions (preload campaign)

func (h *transactionHandler) GetUserTransactions(c *gin.Context) {

	// Mengambil data user siapa yg menggunakan request
	currentUser := c.MustGet("currentUser").(user.User)
	userID := currentUser.ID

	transactions, err := h.service.GetTransactionsByUserID(userID)

	if err != nil {
		response := helper.APIResponse("Failed to get user's transactions", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("User's transactions", http.StatusOK, "success", transaction.FormatUserTransactions(transactions))
	c.JSON(http.StatusOK, response)
}

func (h *transactionHandler) CreateTransaction(c *gin.Context) {
	var input transaction.CreateTransactionInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to create transaction", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)

	input.User = currentUser

	NewTransaction, err := h.service.CreateTransaction(input)

	if err != nil {
		response := helper.APIResponse("Failed to create transaction", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Success to create transaction", http.StatusOK, "success", NewTransaction)
	c.JSON(http.StatusOK, response)
}

// ada input dari user
// handler tangkap input terus mapping ke input struct
// Panggil service buat transaksi, manggil sistem midtrans
// Panggil repository create new transaction data
