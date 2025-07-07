// services/post_service.go

package services

import (
	"blog/database"
	"blog/models"
	"fmt"
)

// 修改 CreatePost 返回完整对象
func CreatePost(post *models.Post) (*models.Post, error) {
	result := database.DB.Create(post)
	if result.Error != nil {
		return nil, result.Error
	}

	// 重新查询以获取关联的User信息
	return GetPostByID(fmt.Sprintf("%d", post.ID))
}

// 修改 UpdatePost 返回完整对象
func UpdatePost(id string, post *models.Post) (*models.Post, error) {
	result := database.DB.Model(&models.Post{}).Where("id = ?", id).Updates(post)
	if result.Error != nil {
		return nil, result.Error
	}

	return GetPostByID(id)
}

func GetAllPosts() ([]models.Post, error) {
	var posts []models.Post
	result := database.DB.Preload("User").Find(&posts)
	return posts, result.Error
}

func GetPostByID(id string) (*models.Post, error) {
	var post models.Post
	result := database.DB.Preload("User").First(&post, id)
	return &post, result.Error
}

func DeletePost(id string) error {
	var post models.Post
	result := database.DB.Delete(&post, id)
	return result.Error
}

func SearchPosts(keyword string) ([]models.Post, error) {
	var posts []models.Post
	keywordLower := "%" + keyword + "%"
	result := database.DB.Preload("User").
		Where("title LIKE ? OR content LIKE ?", keywordLower, keywordLower).
		Find(&posts)
	return posts, result.Error
}
