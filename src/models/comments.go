package models

type Comment struct {
	Id             int    `gorm:"primary_key"`
	Content        string `gorm:"not null"`
	AuthorUsername int    `gorm:"not null"`
	PostSlug       string `gorm:"not null"`
	CreatedAt      uint64 `gorm:"autoCreateTime"`
	UpdatedAt      uint64 `gorm:"autoUpdateTime"`
}
