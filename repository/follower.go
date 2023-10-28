package repository

import (
	"github.com/akhil-is-watching/post_service/models"
	"gorm.io/gorm"
)

type FollowerRepository struct {
	db *gorm.DB
}

func NewFollowerRepository(db *gorm.DB) *FollowerRepository {
	return &FollowerRepository{db}
}

func (repo *FollowerRepository) Create(p *models.Follower) error {
	if err := repo.db.Create(p).Error; err != nil {
		return err
	}
	return nil
}

func (repo *FollowerRepository) All() ([]models.Follower, error) {
	var followers []models.Follower
	if err := repo.db.Find(&followers).Error; err != nil {
		return followers, err
	}

	return followers, nil
}

func (repo *FollowerRepository) AllFollowerByUser(Username string) ([]models.Follower, error) {
	var followers []models.Follower
	if err := repo.db.Where("user_id = ?", Username).Find(&followers).Error; err != nil {
		return followers, err
	}
	return followers, nil
}

func (repo *FollowerRepository) CountFollowerByUser(Username string) (int, error) {
	var followers []models.Follower
	if err := repo.db.Where("user_id = ?", Username).Find(&followers).Error; err != nil {
		return 0, err
	}

	return len(followers), nil
}
