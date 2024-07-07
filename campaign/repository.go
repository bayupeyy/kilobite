package campaign

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Campaign, error)
	FindByUserID(userID int) ([]Campaign, error)
	FindByID(ID int) (Campaign, error)
	Save(campaign Campaign) (Campaign, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) FindAll() ([]Campaign, error) {
	var campaigns []Campaign

	err := r.db.Find(&campaigns).Error
	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

//Fungsi untuk Find by user
func (r *repository) FindByUserID(userID int) ([]Campaign, error) {
	var campaigns []Campaign
	err := r.db.Where("user_id = ?", userID).Find(&campaigns).Error
	if err != nil {
		return campaigns, err
	}
	return campaigns, nil
}

//Fungsi untuk mencari ID
func (r *repository) FindByID(ID int) (Campaign, error) {
	var campaign Campaign
	err := r.db.Preload("User").Preload("CampaignImages").Where("id = ?", ID).Find(&campaign).Error

	if err != nil {
		return campaign, err
	}

	return campaign, nil
}

//Implementasi kontrak Save
func (r *repository) Save(campaign Campaign) (Campaign, error) {
	err := r.db.Create(&campaign).Error
	if err != nil {
		return campaign, err
	}

	return campaign, nil
}
