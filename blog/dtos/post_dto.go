package dtos

import "blog/models"

// @swagger:model
type CreatePostDTO struct {
	// 文章标题
	// Required: true
	Title string `json:"title" binding:"required"`
	// 文章内容
	// Required: true
	Content string `json:"content" binding:"required"`
}

// @swagger:model
type UpdatePostDTO struct {
	// 可选：文章标题
	Title string `json:"title"`
	// 可选：文章内容
	Content string `json:"content"`
}

// @swagger:model
type PostDTO struct {
	// 帖子唯一ID
	ID uint `json:"id"`
	// 帖子标题
	Title string `json:"title"`
	// 帖子内容
	Content string `json:"content"`
	// 作者用户ID
	UserID uint `json:"user_id"`
	// 作者信息
	User UserDTO `json:"user"`
}

// @swagger:model
type UserDTO struct {
	// 用户唯一ID
	ID uint `json:"id"`
	// 用户名
	Username string `json:"username"`
}

func ToModelPost(dto *CreatePostDTO) *models.Post {
	return &models.Post{
		Title:   dto.Title,
		Content: dto.Content,
	}
}

func ToModelPostUpdate(dto *UpdatePostDTO) *models.Post {
	return &models.Post{
		Title:   dto.Title,
		Content: dto.Content,
	}
}

func ToPostDTO(post *models.Post) *PostDTO {
	return &PostDTO{
		ID:      post.ID,
		Title:   post.Title,
		Content: post.Content,
		UserID:  post.UserID,
		User: UserDTO{
			ID:       post.User.ID,
			Username: post.User.Username,
		},
	}
}

func ToPostDTOs(posts []models.Post) []PostDTO {
	var dtos []PostDTO
	for _, post := range posts {
		dtos = append(dtos, *ToPostDTO(&post))
	}
	return dtos
}
