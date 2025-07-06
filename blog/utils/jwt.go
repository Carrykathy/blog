// utils/jwt.go

package utils

import (
	"blog/models"
	"blog/pkg/config"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// 从环境变量中读取 JWT 密钥
var jwtSecret = []byte(config.LoadConfig().JWTSecret)

type Claims struct {
	UserID               uint   `json:"user_id"`
	Username             string `json:"username"`
	jwt.RegisteredClaims        // 使用新版结构
}

// 简化初始化逻辑
func GenerateToken(user models.User) (string, error) {
	cfg := config.LoadConfig() // 统一使用config

	claims := &Claims{
		UserID:   user.ID,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(cfg.JWTExpireHours) * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(cfg.JWTSecret))
}

func ParseToken(tokenString string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
