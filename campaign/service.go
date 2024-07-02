package campaign

type Service interface {
	GetCampaigns(userID int) ([]Campaign, error)
	GetCampaignByID(input GetCampaignDetailInput) (Campaign, error)
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
