package handler

import (
	"kilobite/campaign"
	"kilobite/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//Langkah langkah
//1. Tangkap parameter di handler
//2. handler ke service
//3. service yang menentukan repository mana yang di-call
//4. repository : FindAll, GetByUserID
//5. DB

type campaignHandler struct {
	service campaign.Service
}

// Membuat object dari campaignhandler
func NewCampaignHandler(service campaign.Service) *campaignHandler {
	return &campaignHandler{service}
}

// api/v1/campaign
func (h *campaignHandler) GetCampaign(c *gin.Context) {
	//Tangkap parameter di handler
	userID, _ := strconv.Atoi(c.Query("user_id"))

	campaign, err := h.service.GetCampaigns(userID)

	//pengecekan error
	if err != nil {
		response := helper.APIResponse("Error to get campaign", http.StatusBadRequest, "Error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("List of campaigns", http.StatusBadRequest, "success", campaign)
	c.JSON(http.StatusOK, response)
}
