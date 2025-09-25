package db

import (
	"log"
	"os"

	"github.com/Ansalps/BrotoStack/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDb() *gorm.DB {
	// if err := godotenv.Load(); err != nil {
	// 	log.Println("No .env file found, using environment variables.")
	// }
	// connStr := "user=" + os.Getenv("db_user") + " host=" + os.Getenv("db_host") + " password=" + os.Getenv("db_password") +
	// 	" dbname=" + os.Getenv("db_name") + " port=" + os.Getenv("db_port") + " sslmode=disable"
	connStr := os.Getenv("db_url")
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&models.Users{})
	return db
}
