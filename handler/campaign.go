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
func (h *campaignHandler) GetCampaigns(c *gin.Context) {
	//Tangkap parameter di handler
	userID, _ := strconv.Atoi(c.Query("user_id"))

	campaigns, err := h.service.GetCampaigns(userID)
	if err != nil {
		response := helper.APIResponse("Error to get campaign", http.StatusBadRequest, "Error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("List of campaigns", http.StatusBadRequest, "success", campaign.FormatCampaigns(campaigns))
	c.JSON(http.StatusOK, response)
}

func (h *campaignHandler) GetCampaign(c *gin.Context) {
	// api/v1/campaign/2
	// handler : mapping id yang di url ke struct input => service, call formatter
	// Service : inputannya struct input => menangkap id di url, manggil repo
	// Repository  get campaign by id

	var input campaign.GetCampaignDetailInput

	err := c.ShouldBindUri(&input)
	//pengecekan error
	if err != nil {
		response := helper.APIResponse("Failed to get detail of campaign", http.StatusBadRequest, "Error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	campaignDetail, err := h.service.GetCampaignByID(input)
	//pengecekan error
	if err != nil {
		response := helper.APIResponse("Failed to get detail of campaign", http.StatusBadRequest, "Error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Campaign detail", http.StatusOK, "success", campaign.FormatCampaignDetail(campaignDetail))
	c.JSON(http.StatusOK, response)

}
