package utils

import (
	"github.com/gin-gonic/gin"
)

// SuccessResponse 通用成功响应结构
type SuccessResponse struct {
	Data interface{} `json:"data"`
}

// ErrorResponse 通用错误响应结构
type ErrorResponse struct {
	Error   string `json:"error"`
	Details string `json:"details,omitempty"`
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(200, SuccessResponse{Data: data})
}

func BadRequest(c *gin.Context, err error) {
	c.JSON(400, ErrorResponse{
		Error:   "Bad Request",
		Details: err.Error(),
	})
}

func Unauthorized(c *gin.Context, err error) {
	c.JSON(401, ErrorResponse{
		Error:   "Unauthorized",
		Details: err.Error(),
	})
}

func InternalServerError(c *gin.Context, err error) {
	c.JSON(500, ErrorResponse{
		Error:   "Internal Server Error",
		Details: err.Error(),
	})
}

// 添加新的响应函数
func NotFound(c *gin.Context, err error) {
	c.JSON(404, ErrorResponse{
		Error:   "Not Found",
		Details: err.Error(),
	})
}

func Forbidden(c *gin.Context, err error) {
	c.JSON(403, ErrorResponse{
		Error:   "Forbidden",
		Details: err.Error(),
	})
}
