package models

import "gorm.io/gorm"

type Books struct{
	gorm.Model //It will add the fields like ID,CreatedAt, UpdatedAt, DeletedAt
	Name string `json:"Name" gorm:"not null"`
	Author string `json:"Author" gorm:"not null"`
}

//The name of the table will be automatically generated
//If we want to customize the table name based on our need we should implement the following function

// func (Books) TableName() string{
// 	return "myTable"
// }