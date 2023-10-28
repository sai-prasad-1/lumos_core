package repository

import (
	"github.com/akhil-is-watching/post_service/models"
	"gorm.io/gorm"
)

type FollowingRepository struct {
	db *gorm.DB
}

func NewFollowingRepository(db *gorm.DB) *FollowingRepository {
	return &FollowingRepository{db}
}

func (repo *FollowingRepository) Create(p *models.Following) error {
	if err := repo.db.Create(p).Error; err != nil {
		return err
	}
	return nil
}

func (repo *FollowingRepository) All() ([]models.Following, error) {
	var following []models.Following
	if err := repo.db.Find(&following).Error; err != nil {
		return following, err
	}

	return following, nil
}

func (repo *FollowingRepository) AllFollowingByUser(Username string) ([]models.Following, error) {
	var following []models.Following
	if err := repo.db.Where("user_id = ?", Username).Find(&following).Error; err != nil {
		return following, err
	}
	return following, nil
}

func (repo *FollowingRepository) CountFollowingByUser(Username string) (int, error) {
	var following []models.Following
	if err := repo.db.Where("user_id = ?", Username).Find(&following).Error; err != nil {
		return 0, err
	}

	return len(following), nil
}
