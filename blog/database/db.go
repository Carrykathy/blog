package database

import (
	"blog/models"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// 全局变量，并导出（首字母大写）
var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // 启用SQL日志
	})

	if err != nil {
		panic("failed to connect database: " + err.Error())
	}

	// 确保自动迁移正确执行
	err = DB.AutoMigrate(&models.User{}, &models.Post{})
	if err != nil {
		panic("failed to migrate database: " + err.Error())
	}

	fmt.Println("Database initialized successfully")
}
