package models

import "gorm.io/gorm"

type Books struct{
	gorm.Model //It will add the fields like ID,CreatedAt, UpdatedAt, DeletedAt
	Name string `json:"Name" gorm:"not null"`
	Author string `json:"Author" gorm:"not null"`
}