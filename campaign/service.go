package campaign

import (
	"fmt"

	"github.com/gosimple/slug"
)

type Service interface {
	GetCampaigns(userID int) ([]Campaign, error)
	GetCampaignByID(input GetCampaignDetailInput) (Campaign, error)
	CreateCampaign(input CreateCampaignInput) (Campaign, error)
	UpdateCampaign(inputID GetCampaignDetailInput, inputData CreateCampaignInput) (Campaign, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetCampaigns(userID int) ([]Campaign, error) {
	if userID != 0 {
		campaign, err := s.repository.FindByUserID(userID)

		//Melakukan pengecekan error
		if err != nil {
			return campaign, err
		}

		return campaign, nil
	}

	campaign, err := s.repository.FindAll()

	//Melakukan pengecekan error
	if err != nil {
		return campaign, err
	}

	return campaign, nil
}

// Membuat fungsi untuk GetCampaignById
func (s *service) GetCampaignByID(input GetCampaignDetailInput) (Campaign, error) {
	campaign, err := s.repository.FindByID(input.ID)

	//Pengecekan Error
	if err != nil {
		return campaign, err
	}

	return campaign, nil
}

// Membuat fungsi untuk CreateCampaign
func (s *service) CreateCampaign(input CreateCampaignInput) (Campaign, error) {
	//Langkah dan maksud koding ini
	//1. Dari inputan user di mapping ke CreateCampaignInput
	//2. Dari CreateCampaignInput di mapping ke Campaign Object
	campaign := Campaign{}     // Perintah untuk membuat Object campaign
	campaign.Name = input.Name // Mapping 1 Per 1
	campaign.ShortDescription = input.Description
	campaign.Description = input.Description
	campaign.Perks = input.Perks
	campaign.GoalAmount = input.GoalAmount
	campaign.UserID = input.User.ID
	//Pembuatan Slug (tulisan jadi ada - contoh bayu-ajie-pamungkas)

	slugCandidate := fmt.Sprintf("%s %d", input.Name, input.User.ID)
	campaign.Slug = slug.Make(slugCandidate)

	newCampaign, err := s.repository.Save(campaign)
	if err != nil {
		return newCampaign, err
	}

	return newCampaign, nil
}

// Implementasi untuk Update
func (s *service) UpdateCampaign(inputID GetCampaignDetailInput, inputData CreateCampaignInput) (Campaign, error) {
	campaign, err := s.repository.FindByID(inputID.ID)
	if err != nil {
		return campaign, err
	}

	campaign.Name = inputData.Name
	campaign.ShortDescription = inputData.ShortDescription
	campaign.Description = inputData.Description
	campaign.Perks = inputData.Perks
	campaign.GoalAmount = inputData.GoalAmount

	updatedCampaign, err := s.repository.Update(campaign)
	if err != nil {
		return updatedCampaign, err
	}

	return updatedCampaign, nil
}
