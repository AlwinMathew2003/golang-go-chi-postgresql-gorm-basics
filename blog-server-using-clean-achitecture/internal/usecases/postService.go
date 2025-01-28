package usecases

import("blog-server-using-clean-architecture/internal/repositories"
"blog-server-using-clean-architecture/internal/models")

type PostService struct{
	PostRepo repositories.PostRepository
}

//GetPostByID,GetAllPosts,CreatePost,DeletePost,UpdatePost

func (s *PostService)GetPostByIDService(id string)(*models.Post,error){
	return s.PostRepo.GetPostByID(id)
}

func (s *PostService)GetAllPostService()([]*models.Post,error){
	return s.PostRepo.GetAllPost()
}

func (s *PostService)CreatePostService(post *models.Post,id uint)(*models.Post,error){
	return s.PostRepo.CreatePost(post,id)
}

func (s *PostService)DeletePostService(id string)(error){
	return s.PostRepo.DeletePost(id)
}

func (s *PostService)UpdatePostService(post *models.Post,id string)(*models.Post,error){
	return s.PostRepo.UpdatePost(post,id)
}