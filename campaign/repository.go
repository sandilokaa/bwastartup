package campaign

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Campaign, error)
	FindByUserID(userID int) ([]Campaign, error)
}

type repostiory struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repostiory {
	return &repostiory{db}
}

func (r *repostiory) FindAll() ([]Campaign, error) {
	var campaigns []Campaign

	err := r.db.Find(&campaigns).Error
	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

func (r *repostiory) FindByUserID(userID int) ([]Campaign, error) {
	var campaigns []Campaign

	err := r.db.Where("user_id = ?", userID).Preload("CampaignImages", "campaign_images.is_primary = 1").Find(&campaigns).Error
	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}
