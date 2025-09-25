package db

import (
	"log"
	"os"

	"github.com/Ansalps/BrotoStack/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDb() *gorm.DB{
	// connStr:="user="+os.Getenv("db_user")+" host="+os.Getenv("db_host")+" password="+os.Getenv("db_password")+
	// " dbname="+os.Getenv("db_name")+" port="+os.Getenv("db_port")+" sslmode=disable"
	connStr:="postgresql://render:GBLyLQcmMuZAWFDBVNq9z4zAQmLvSJvX@dpg-d3akl6buibrs73ccu3a0-a/my_database_a4hz"
	db,err:=gorm.Open(postgres.Open(connStr),&gorm.Config{})
	if err!=nil{
		log.Fatal(err)
	}
	db.AutoMigrate(&models.Users{})
	return db
}