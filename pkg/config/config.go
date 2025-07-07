// pkg/config/config.go

package config

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/joho/godotenv"
)

// Config 定义应用程序所需的全局配置项
type Config struct {
	// JWT 密钥，用于签名和解析 Token
	JWTSecret string

	// JWT 过期时间（小时）
	JWTExpireHours int

	// 最大刷新时间，用于控制 Refresh Token 的最大有效期
	JWTMaxRefresh time.Duration

	// Redis 地址（如 localhost:6379），可用于缓存 Token 或其他用途
	RedisAddr string

	// 数据库连接字符串（SQLite/MySQL/PostgreSQL）
	DatabaseDSN string
}

var (
	appConfig *Config
	once      sync.Once
)

func LoadConfig() *Config {
	once.Do(func() {
		// 加载.env文件（如果存在）
		_ = godotenv.Load()

		appConfig = &Config{
			JWTSecret:      getEnv("JWT_SECRET", "default-secret-key"),
			JWTExpireHours: getEnvAsInt("JWT_EXPIRE_HOURS", 24),
			DatabaseDSN:    getEnv("DATABASE_DSN", "test.db"),
		}

		log.Println("Configuration loaded successfully")
	})
	return appConfig
}

// getEnv 获取环境变量，如果不存在则返回默认值
func getEnv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// getEnvAsInt 获取环境变量并转换为整数，失败返回默认值
func getEnvAsInt(name string, defaultValue int) int {
	valueStr := os.Getenv(name)
	var value int
	_, err := fmt.Sscanf(valueStr, "%d", &value)
	if err != nil {
		return defaultValue
	}
	return value
}
