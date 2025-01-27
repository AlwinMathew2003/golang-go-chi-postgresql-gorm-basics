package db

import (
    "log"
    "go-crud-using-database1/models"
    "go-crud-using-database1/utils"
)

func SeedDB() {
    // Check if the table already has data
    var userCount int64
    DB.Model(&models.User{}).Count(&userCount) // Use the global DB variable

    if userCount == 0 {
        // Create test user data
		password,_:= utils.HashPassword("password123")

        err := DB.Create(&models.User{
            Username: "testuser",
            Password: password, // Hash the password
        }).Error
        if err != nil {
            log.Fatalf("Error seeding database: %v", err)
        }
        log.Println("Database seeded successfully.")
    } else {
        log.Println("Database already seeded.")
    }
}
