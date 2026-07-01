package model

import "github.com/hanzjefferson/go-cleanarch/internal/entity"

type UserResponseData struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

func UserResponseDTO(entity *entity.User) *UserResponseData {
	return &UserResponseData{
		ID: entity.ID,
		Email: entity.Email,
		Username: entity.Username,
	}
}