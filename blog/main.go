// @title Blog API 文档
// @version 1.0
// @description 基于 Gin 框架的博客系统 API 接口文档
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
// @schemes http https
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @description 使用 Bearer Token 进行身份验证，格式为 `Bearer <token>`

package main

import (
	"blog/database"
	"blog/pkg/config"
	"blog/routers"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
)

func main() {
	// 初始化配置
	cfg := config.LoadConfig()
	log.Printf("Database DSN: %s", cfg.DatabaseDSN)

	// 初始化数据库
	database.InitDB()

	// 初始化路由
	r := gin.Default()
	routers.RegisterRoutes(r)

	// Swagger 路由
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 启动服务
	log.Printf("Server is running on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
