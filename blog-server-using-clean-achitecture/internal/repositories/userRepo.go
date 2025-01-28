package repositories

import("blog-server-using-clean-architecture/internal/models")

type UserRepository interface{
	Register(user *models.User)(error)
	FindByEmail(email string)(*models.User,error)
}