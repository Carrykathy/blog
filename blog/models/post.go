package models

type Post struct {
	ID      uint   `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	UserID  uint   `json:"user_id"`
	User    User   `gorm:"foreignKey:UserID" json:"user"`
}
