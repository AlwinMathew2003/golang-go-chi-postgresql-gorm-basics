package models

import("gorm.io/gorm")

type User struct{
	gorm.Model
	Name string `json:"username" gorm:"not null"`
	Email string `json:"email" gorm:"not null"`
	Password string `json:"password" gorm:"not null"`
	Posts []Post `json:"posts"` //Here we don't want to mention models.Post since it is in same package models
}