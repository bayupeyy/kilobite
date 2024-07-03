package campaign

import "strings"

type CampaignFormatter struct {
	ID               int    `json:"id"`
	UserID           int    `json:"user_id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	ImageURL         string `json:"image_url"`
	GoalAmount       int    `json:"goal_amount"`
	CurrentAmount    int    `json:"current_amount"`
	Slug             string `json:"slug"`
}

func FormatCampaign(campaign Campaign) CampaignFormatter {
	CampaignFormatter := CampaignFormatter{}
	CampaignFormatter.ID = campaign.ID
	CampaignFormatter.UserID = campaign.UserID
	CampaignFormatter.Name = campaign.Name
	CampaignFormatter.ShortDescription = campaign.ShortDescription
	CampaignFormatter.GoalAmount = campaign.GoalAmount
	CampaignFormatter.CurrentAmount = campaign.CurrentAmount
	CampaignFormatter.Slug = campaign.Slug
	CampaignFormatter.ImageURL = ""

	if len(campaign.CampaignImages) > 0 {
		CampaignFormatter.ImageURL = campaign.CampaignImages[0].FileName

	}

	return CampaignFormatter
}

func FormatCampaigns(campaigns []Campaign) []CampaignFormatter {
	campaignsFormatter := []CampaignFormatter{}
	for _, campaign := range campaigns {
		campaignFormatter := FormatCampaign(campaign)
		campaignsFormatter = append(campaignsFormatter, campaignFormatter)
	}

	return campaignsFormatter
}

type CampaignDetailFormatter struct {
	ID               int      `json:"id"`
	Name             string   `json:"name"`
	ShortDescription string   `json:"short_description"`
	Description      string   `json:"description"`
	ImageURL         string   `json:"image_url"`
	GoalAmount       int      `json:"goal_amount"`
	CurrentAmount    int      `json:"current_amount"`
	UserID           int      `json:"user_id"`
	Slug             string   `json:"slug"`
	Perks            []string `json:"perks"`
}

// Fungsi untuk detail Campaign Detail
func FormatCampaignDetail(campaign Campaign) CampaignDetailFormatter {
	campaignDetailFormatter := CampaignDetailFormatter{}
	campaignDetailFormatter.ID = campaign.ID
	campaignDetailFormatter.Name = campaign.Name
	campaignDetailFormatter.ShortDescription = campaign.ShortDescription
	campaignDetailFormatter.Description = campaign.Description
	campaignDetailFormatter.ImageURL = ""
	campaignDetailFormatter.GoalAmount = campaign.GoalAmount
	campaignDetailFormatter.CurrentAmount = campaign.CurrentAmount
	campaignDetailFormatter.Slug = campaign.Slug
	campaignDetailFormatter.UserID = campaign.UserID

	if len(campaign.CampaignImages) > 0 {
		campaignDetailFormatter.ImageURL = campaign.CampaignImages[0].FileName
	}

	var perks []string

	//Untuk memecah isi perks setiap ada koma
	for _, perk := range strings.Split(campaign.Perks, ",") {
		perks = append(perks, strings.TrimSpace(perk)) // strings.TrimSpace digunakan untuk menghilangkan spasi di depan
	}
	return campaignDetailFormatter
}
