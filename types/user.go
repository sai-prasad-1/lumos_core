package types

import "github.com/akhil-is-watching/post_service/models"

type UserCreate struct {
	ID          string `json:"username"`
	Name        string `json:"name"`
	College     string `json:"college"`
	Description string `json:"description"`
	Email       string `json:"email"`
	Password    string `json:"password"`
}

func (u UserCreate) Convert() *models.User {
	return &models.User{
		ID:                u.ID,
		Email:             u.Email,
		Name:              u.Name,
		College:           u.College,
		Description:       u.Description,
		PasswordHash:      u.Password,
		WalletAddress:     "0x95222290DD7278Aa3Ddd389Cc1E1d165CC4BAfe5",
		WalletPrivkeyHash: "0x95222290DD7278Aa3Ddd389Cc1E1d165CC4BAfe5",
	}
}
