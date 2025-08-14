package models

type User struct {
	Id        int    `gorm:"primary_key"`
	Username  string `gorm:"unique;not null;index;"`
	Password  string `gorm:"not null;"`
	Email     string `gorm:"unique;not null;index;"`
	Posts     []Post `gorm:"foreignkey:AuthorUsername"`
	CreatedAt uint64 `gorm:"autoCreateTime"`
	UpdatedAt uint64 `gorm:"autoUpdateTime"`
}
