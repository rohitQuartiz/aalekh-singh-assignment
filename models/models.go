package models

import "gorm.io/gorm"

type Blog struct {
	gorm.Model
	Id    uint16   `json:"id" gorm:"primaryKey; autoIncrement"`
	Title string   `json:"title"`
	Body  BlogBody `json:"blog_body" gorm:"serializer:json"`
}

type BlogBody struct {
	Name        string `json:"name"`
	Image       string `json:"image"`
	Description string `json:"description"`
}

type Settings struct {
	DB_HOST        string
	DB_TYPE        string
	SQLITE_DB_NAME string
}
