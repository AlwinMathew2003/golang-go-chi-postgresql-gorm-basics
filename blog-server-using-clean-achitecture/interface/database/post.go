package database



import("blog-server-using-clean-architecture/internal/models"
"gorm.io/gorm")



//Whether we want to declare the methods here in repositories?
//No but we need to declare the database connection variable with in the repository

type PostRepositoryDB struct{
	DB *gorm.DB
}

func (db *PostRepositoryDB)GetPostByID(id string)(*models.Post,error){

	var post *models.Post

	if err := db.DB.First(&post,id).Error;err!=nil{
		return nil,err
	}

	return post,nil
}

func (db *PostRepositoryDB)GetAllPost()([]*models.Post,error){
	var post []*models.Post

	if err:=db.DB.Find(&post).Error;err!=nil{
		return nil,err
	}

	return post,nil
}

func (db *PostRepositoryDB)CreatePost(post *models.Post,userID uint)(*models.Post,error){

	post.UserID = userID
	if err:=db.DB.Create(&post).Error;err!=nil{
		return nil,err
	}

	return post,nil
}

func (db *PostRepositoryDB)UpdatePost(post *models.Post,id string)(*models.Post,error){
	var UpdatedPost models.Post

	if err:= db.DB.First(&UpdatedPost,id).Error;err!=nil{
		return nil,err
		//we should use errors.New() for defining the error
		//errors library should be imported
	}

	UpdatedPost.Title = post.Title
	UpdatedPost.Description = post.Description

	if err:= db.DB.Save(&UpdatedPost).Error;err!=nil{
		return nil,err
	}

	return &UpdatedPost,nil

}

func (db *PostRepositoryDB)DeletePost(id string)(error){

	if err:= db.DB.Delete(&models.Post{},id).Error;err!=nil{
		return err
	}

	return nil
}