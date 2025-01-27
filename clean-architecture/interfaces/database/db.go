package database

//go get gorm.io/driver/postgres this installation is requried
import (
	"clean-architecture/internal/models"
	"log"
	"os"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect()(*gorm.DB){
	
	if err:=godotenv.Load();err!=nil{
		log.Fatalf("Error in fetching the environmental variable")
	}

	dsn := os.Getenv("DSN")

	db,err:=gorm.Open(postgres.Open(dsn),&gorm.Config{})

	if err!=nil{
		log.Fatalf("Failed to connect to database")
		return nil
	}

	log.Println("Successfully connected to the database")

	DB = db
	return db

}

func Migrate(){
	if err:=DB.AutoMigrate(models.Book{},models.User{});err!=nil{
		log.Fatalf("Failed to migrate to database")
		return
	}

	log.Println("Database migration completed")
}