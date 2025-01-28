package usecases

import("blog-server-using-clean-architecture/internal/repositories"
"blog-server-using-clean-architecture/internal/models")

type UserService struct{
	UserRepo repositories.UserRepository
}

func (s *UserService)RegisterService(user *models.User)(error){
	return s.UserRepo.Register(user)
}

func (s *UserService)FindByEmailService(email string)(*models.User,error){
	return s.UserRepo.FindByEmail(email)
}