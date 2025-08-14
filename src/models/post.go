package models

type Post struct {
	Id             int    `gorm:"primary_key"`
	Title          string `gorm:"not null"`
	Content        string `gorm:"not null"`
	Slug           string `gorm:"not null;index;unique"`
	AuthorUsername string `gorm:"not null"`
	CreatedAt      uint64 `gorm:"autoCreateTime"`
	UpdatedAt      uint64 `gorm:"autoUpdateTime"`
}
