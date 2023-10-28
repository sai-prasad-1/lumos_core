package repository

import (
	"github.com/akhil-is-watching/post_service/models"
	"gorm.io/gorm"
)

type LoomRepository struct {
	db *gorm.DB
}

func NewLoomRepository(db *gorm.DB) *LoomRepository {
	return &LoomRepository{db}
}

func (repo *LoomRepository) Create(p *models.Loom) error {
	if err := repo.db.Create(p).Error; err != nil {
		return err
	}
	return nil
}

func (repo *LoomRepository) Get(LoomID string) (models.Loom, error) {
	var loom models.Loom
	if err := repo.db.Where("id = ?", LoomID).First(&loom).Error; err != nil {
		return loom, err
	}

	return loom, nil
}

func (repo *LoomRepository) All() ([]models.Loom, error) {
	var looms []models.Loom
	if err := repo.db.Find(&looms).Error; err != nil {
		return looms, err
	}

	return looms, nil
}

func (repo *LoomRepository) AllLoomsByUser(Username string) ([]models.Loom, error) {
	var looms []models.Loom
	if err := repo.db.Where("user_id = ?", Username).Find(&looms).Error; err != nil {
		return looms, err
	}
	return looms, nil
}

func (repo *LoomRepository) GetFeedForUser(Username string) ([]models.Loom, error) {
	var looms []models.Loom
	if err := repo.db.Joins("JOIN followings ON looms.user_id = followings.following_id").
		Where("followings.user_id = (SELECT id FROM users WHERE username = ?)", Username).
		Order("looms.created_at desc").
		Find(&looms).Error; err != nil {
		return looms, err
	}

	return looms, nil
}
