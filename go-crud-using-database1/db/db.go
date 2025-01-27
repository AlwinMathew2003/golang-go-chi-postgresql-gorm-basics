package db

import(
"go-crud-using-database1/models"
"gorm.io/driver/postgres"
"gorm.io/gorm"
"log")

var DB *gorm.DB //Declaring a database instance, gorm.DB is a object type

func Connect(){

	//Setting the credentials for connections
	dsn:="host=localhost user=postgres password=root dbname=Books port=5432 sslmode=disable TimeZone=Asia/Kathmandu"

	//Connecting to the database gorm.Open(databasetype,pointer to defaultConfiguraion)
	database,err := gorm.Open(postgres.Open(dsn),&gorm.Config{}); 

	//Checking if any error has occured
	if err!= nil{
		log.Fatalf("Failed to connect to the database: %s",err)
	}

	DB = database

	log.Println("Connected to the database")
}

func Migrate(){
	if err:=DB.AutoMigrate(&models.Books{},&models.User{}); err!=nil{
		log.Fatalf("Failed to migrate database: %s",err)
	}
	log.Println("Database migration completed")
}