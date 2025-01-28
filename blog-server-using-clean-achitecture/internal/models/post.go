package models

import("gorm.io/gorm")

type Post struct{
	gorm.Model
	Title string `json:"title" gorm:"not null"`
	Description string `json:"description" gorm:"not null"`
	UserID uint `json:"user_id"` //It will automatically act as foreing key
	User User `json:"user"`
}