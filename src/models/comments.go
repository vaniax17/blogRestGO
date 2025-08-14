package models

type Comment struct {
	Id             int    `gorm:"primary_key"`
	Content        string `gorm:"not null"`
	AuthorUsername string `gorm:"not null"`
	PostSlug       string `gorm:"not null;index;"`
	Slug           string `gorm:"not null;index;"`
	CreatedAt      uint64 `gorm:"autoCreateTime"`
	UpdatedAt      uint64 `gorm:"autoUpdateTime"`
}
