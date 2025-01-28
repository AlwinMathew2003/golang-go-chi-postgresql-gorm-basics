package database

import(
	"gorm.io/gorm"
	"os"
	"log"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"blog-server-using-clean-architecture/internal/models"
	"golang.org/x/crypto/bcrypt"
)

//Specifying the error using error.New("Any error")

func Connect()(*gorm.DB){
	//Accessing the environmental variabls for dsn
	if err:= godotenv.Load();err!=nil{
		log.Fatalf("Error in accessing the environmental variables")
		return nil
	}

	dsn := os.Getenv("DSN")

	database,err := gorm.Open(postgres.Open(dsn),&gorm.Config{})

	if err!=nil{
		log.Fatalf("Error in connecting the database")
		return nil
	}

	log.Println("Successfully connected to the database")

	return database

}

func Migrate(db *gorm.DB){
	if err:=db.AutoMigrate(&models.User{},&models.Post{});err!=nil{
		log.Fatalf("Error in migrating the database")
		return
	}

	log.Println("Successfully migrated to the database")
}

func SeedDB(DB *gorm.DB) {
    // Check if the table already has data
    var userCount int64
    DB.Model(&models.User{}).Count(&userCount) // Count existing users

    if userCount == 0 {
        // Create a test user

		hashedPassword,err:=bcrypt.GenerateFromPassword([]byte("admin@1234"),bcrypt.DefaultCost)

		if err!=nil{
			log.Fatalf("Error in hasing password")
		}
        user := models.User{
            Name:  "user",
            Email: "user@gmail.com",
			Password: string(hashedPassword),
        }

        // Create the user first
        if err := DB.Create(&user).Error; err != nil {
            log.Fatalf("Error seeding database: %v", err)
        }

        // Create a post associated with the user
        post := models.Post{
            Title:       "Hai",
            Description: "Hello world",
            UserID:      user.ID, // Use the created user's ID
        }

        if err := DB.Create(&post).Error; err != nil {
            log.Fatalf("Error creating post: %v", err)
        }

        log.Println("Database seeded successfully.")
    } else {
        log.Println("Database already seeded.")
    }
}
