package controllers

import (
	"blog/dtos"
	"blog/services"
	"blog/utils"
	"github.com/gin-gonic/gin"
)

// @Summary 用户注册
// @Description 创建一个新用户账户
// @Tags Users
// @Accept json
// @Produce json
// @Param user body controllers.RegisterInput true "用户名和密码"
// @Success 200 {object} dtos.UserInfoDTO
// @Failure 400 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
func RegisterUser(c *gin.Context) {
	type RegisterInput struct {
		Username string `json:"username" binding:"required,min=3"`
		Password string `json:"password" binding:"required,min=6"`
	}

	var input RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.BadRequest(c, err)
		return
	}

	user, err := services.RegisterUser(input.Username, input.Password)
	if err != nil {
		utils.InternalServerError(c, err)
		return
	}

	utils.Success(c, dtos.ToUserInfoDTO(user))
}

// @Summary 用户登录
// @Description 使用用户名和密码登录并获取 JWT token
// @Tags Users
// @Accept json
// @Produce json
// @Param credentials body controllers.LoginInput true "用户名和密码"
// @Success 200 {object} gin.H{"user": dtos.UserInfoDTO, "token": string}
// @Failure 400 {object} utils.ErrorResponse
// @Failure 401 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Security ApiKeyAuth
func LoginUser(c *gin.Context) {
	type LoginInput struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.BadRequest(c, err)
		return
	}

	user, err := services.LoginUser(input.Username, input.Password)
	if err != nil {
		utils.Unauthorized(c, err)
		return
	}

	// 生成并返回 JWT token
	token, err := utils.GenerateToken(*user)
	if err != nil {
		utils.InternalServerError(c, err)
		return
	}

	utils.Success(c, gin.H{
		"user":  dtos.ToUserInfoDTO(user),
		"token": token,
	})
}
