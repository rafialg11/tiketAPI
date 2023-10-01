package configs

import (
	"fmt"
	"os"
	"tiketAPI/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := fmt.Sprintf("mysql://%s:%s@%s:%d/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB.AutoMigrate(&models.Event{})
	DB.AutoMigrate(&models.User{})

	// Log the DSN string and the database connection status
	fmt.Printf("DSN: %s\n", dsn)
	if DB != nil {
		fmt.Println("Database connection established")
	} else {
		fmt.Println("Failed to establish database connection")
	}
}
