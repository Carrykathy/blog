package models

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"unique;size:255" json:"username"` // 增加长度限制
	Password string `gorm:"size:255" json:"-"`               // 确保足够存储哈希
}
