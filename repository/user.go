package repository

import (
	"errors"
	"strings"

	"github.com/akhil-is-watching/post_service/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (repo *UserRepository) Create(u *models.User) error {
	if err := repo.db.Create(u).Error; err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique") {
			return errors.New("email already exists")
		}
		return err
	}
	return nil
}

func (repo *UserRepository) Get(Username string) (*models.User, error) {
	var user *models.User
	if err := repo.db.Where("id = ?", Username).First(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}
