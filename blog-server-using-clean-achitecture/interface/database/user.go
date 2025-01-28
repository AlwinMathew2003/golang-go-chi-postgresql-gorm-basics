package database

import (
	"blog-server-using-clean-architecture/internal/models"
	"gorm.io/gorm"
)

type UserRepositoryDB struct{
	DB *gorm.DB
}

func(db *UserRepositoryDB)Register(user *models.User)(error){

	if err:= db.DB.Create(&user).Error;err!=nil{
		return err
	}

	return nil
}

func(db *UserRepositoryDB)FindByEmail(email string)(*models.User,error){

	var user *models.User

	if err:=db.DB.Where("email=?",email).First(&user).Error;err!=nil{
		return nil,err
	}

	return user,nil
}