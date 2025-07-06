package routers

import (
	"blog/controllers"
	"blog/middleware"
	"github.com/gin-gonic/gin"
)

// RegisterRoutes 注册所有 API 路由
// 包含用户相关接口和文章相关接口
func RegisterRoutes(router *gin.Engine) {
	// 用户相关接口
	userGroup := router.Group("/user")
	{
		userGroup.POST("/register", controllers.RegisterUser)
		userGroup.POST("/login", controllers.LoginUser)
	}

	// 文章相关接口（需认证）
	postGroup := router.Group("/posts").Use(middleware.AuthMiddleware())
	{
		postGroup.POST("/", controllers.CreatePost)
		postGroup.GET("/", controllers.GetPosts)
		postGroup.GET("/:id", controllers.GetPostById)
		postGroup.PUT("/:id", controllers.UpdatePost)
		postGroup.DELETE("/:id", controllers.DeletePost)
		postGroup.GET("/search", controllers.SearchPosts)
	}
}
