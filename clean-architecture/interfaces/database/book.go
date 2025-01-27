package database

import (
	"clean-architecture/internal/models"
	"gorm.io/gorm"
)

type BookRepository struct{
	DB *gorm.DB
}


//The implementation of repository is made here.

func (r *BookRepository)Save(book *models.Book)(*models.Book,error){
	if err:= r.DB.Create(book).Error;err!=nil{
		return nil,err
	}
	return book,nil
}

func (r *BookRepository)GetByID(ID string)(*models.Book,error){
	var book models.Book
	if err:=r.DB.First(&book,ID).Error;err!=nil{
		return nil,err
	}
	return &book,nil
}
