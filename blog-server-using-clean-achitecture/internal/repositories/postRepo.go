package repositories

import("blog-server-using-clean-architecture/internal/models")

type PostRepository interface{
	GetAllPost()([]*models.Post,error)
	GetPostByID(id string)(*models.Post,error)
	CreatePost(post *models.Post,id uint)(*models.Post,error)
	UpdatePost(post *models.Post,id string)(*models.Post,error)
	DeletePost(id string)(error)
}