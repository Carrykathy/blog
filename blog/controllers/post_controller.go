package controllers

import (
	"blog/dtos"
	"blog/models"
	"blog/services"
	"blog/utils"
	"errors"
	"github.com/gin-gonic/gin"
)

// CreatePost @Summary 创建新文章
// @Description 根据提供的文章信息创建新的博客文章
// @Tags Posts
// @Accept json
// @Produce json
// @Param post body dtos.CreatePostDTO true "文章信息"
// @Success 200 {object} dtos.PostDTO
// @Failure 400 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Security ApiKeyAuth
func CreatePost(c *gin.Context) {
	var dto dtos.CreatePostDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		utils.BadRequest(c, err)
		return
	}

	// 从中间件上下文中获取用户ID
	user, _ := c.Get("user")
	userID := user.(map[string]interface{})["id"].(uint)

	post := &models.Post{
		Title:   dto.Title,
		Content: dto.Content,
		UserID:  userID, // 使用从上下文中获取的 UserID
	}

	createdPost, err := services.CreatePost(post)
	if err != nil {
		utils.InternalServerError(c, err)
		return
	}

	utils.Success(c, dtos.ToPostDTO(createdPost))
}

// @Summary 创建新文章
// @Description 根据提供的文章信息创建新的博客文章
// @Tags Posts
// @Accept json
// @Produce json
// @Param post body dtos.CreatePostDTO true "文章信息"
// @Success 200 {object} dtos.PostDTO
// @Failure 400 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Security ApiKeyAuth
func UpdatePost(c *gin.Context) {
	id := c.Param("id")
	currentUser, _ := c.Get("user")
	userMap := currentUser.(map[string]interface{})
	userID := userMap["id"].(uint)

	var post dtos.UpdatePostDTO
	if err := c.ShouldBindJSON(&post); err != nil {
		utils.BadRequest(c, err)
		return
	}

	existingPost, err := services.GetPostByID(id)
	if err != nil {
		utils.NotFound(c, errors.New("post not found"))
		return
	}

	if existingPost.UserID != userID {
		utils.Forbidden(c, errors.New("you can only update your own posts"))
		return
	}

	updatedPost := dtos.ToModelPostUpdate(&post)
	updatedPost.UserID = existingPost.UserID // 保留原作者 ID

	updatedPost, err = services.UpdatePost(id, updatedPost)
	if err != nil {
		utils.InternalServerError(c, err)
		return
	}

	utils.Success(c, dtos.ToPostDTO(updatedPost))
}

// @Summary 获取所有文章
// @Description 返回所有博客文章的列表
// @Tags Posts
// @Produce json
// @Success 200 {array} dtos.PostDTO
// @Failure 500 {object} utils.ErrorResponse
func GetPosts(c *gin.Context) {
	posts, err := services.GetAllPosts()
	if err != nil {
		utils.InternalServerError(c, err)
		return
	}

	utils.Success(c, dtos.ToPostDTOs(posts))
}

// @Summary 获取指定ID的文章详情
// @Description 根据文章ID返回其详细信息
// @Tags Posts
// @Produce json
// @Param id path string true "文章ID"
// @Success 200 {object} dtos.PostDTO
// @Failure 500 {object} utils.ErrorResponse
func GetPostById(c *gin.Context) {
	id := c.Param("id")
	post, err := services.GetPostByID(id)
	if err != nil {
		utils.InternalServerError(c, err)
		return
	}

	utils.Success(c, dtos.ToPostDTO(post))
}

// @Summary 删除指定ID的文章
// @Description 删除指定ID的博客文章（仅作者可删除）
// @Tags Posts
// @Produce json
// @Param id path string true "文章ID"
// @Success 200 {object} gin.H
// @Failure 500 {object} utils.ErrorResponse
// @Security ApiKeyAuth
func DeletePost(c *gin.Context) {
	id := c.Param("id")
	if err := services.DeletePost(id); err != nil {
		utils.InternalServerError(c, err)
		return
	}

	utils.Success(c, gin.H{"message": "Post deleted successfully"})
}

// @Summary 搜索文章
// @Description 根据关键词搜索匹配的博客文章
// @Tags Posts
// @Produce json
// @Param q query string true "搜索关键词"
// @Success 200 {array} dtos.PostDTO
// @Failure 400 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
func SearchPosts(c *gin.Context) {
	keyword := c.Query("q")
	if keyword == "" {
		utils.BadRequest(c, errors.New("search keyword is required"))
		return
	}

	posts, err := services.SearchPosts(keyword)
	if err != nil {
		utils.InternalServerError(c, err)
		return
	}

	utils.Success(c, dtos.ToPostDTOs(posts))
}

// 添加新的响应函数
func NotFound(c *gin.Context, err error) {
	c.JSON(404, utils.ErrorResponse{
		Error:   "Not Found",
		Details: err.Error(),
	})
}

func Forbidden(c *gin.Context, err error) {
	c.JSON(403, utils.ErrorResponse{
		Error:   "Forbidden",
		Details: err.Error(),
	})
}
