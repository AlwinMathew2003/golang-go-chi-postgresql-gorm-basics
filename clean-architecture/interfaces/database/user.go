package database

import(
	"gorm.io/gorm"
	"clean-architecture/internal/models"
)

type UserRepo struct {
	DB *gorm.DB
}

func(db *UserRepo)Register(user *models.User)(error){
	if err:=db.DB.Create(user).Error;err!=nil{
		return err
	}
	return nil
}

func (db *UserRepo)FindByName(username string)(*models.User,error){
	var user *models.User
	if err:=db.DB.Where("username=?",username).First(&user).Error;err!=nil{
		return nil,err
	}
	return user,nil
}