package repositories

import("clean-architecture/internal/models")
type UserRepo interface{
	Register(user *models.User)(error)
	FindByName(name string)(*models.User,error)
}