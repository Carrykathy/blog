package dtos

import "blog/models"

// @swagger:model
type UserInfoDTO struct {
	// 用户唯一标识
	ID uint `json:"id"`
	// 用户名
	Username string `json:"username"`
}

func ToUserInfoDTO(user *models.User) UserInfoDTO {
	return UserInfoDTO{
		ID:       user.ID,
		Username: user.Username,
	}
}
