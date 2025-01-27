package models

import(
	"gorm.io/gorm")

type User struct{
	gorm.Model
	//We should use first capital letter to make it as exported field if not json encoding will not work
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password" gorm:"not null"`	
}
