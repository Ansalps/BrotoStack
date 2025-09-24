package db

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDb() *gorm.DB{
	connStr:="host="+os.Getenv("db_host")+" user="+os.Getenv("db_user")+" password="+os.Getenv("db_password")+
	" dbname="+os.Getenv("db_name")+" port="+os.Getenv("db_port")+" sslmode=disable"
	db,err:=gorm.Open(postgres.Open(connStr),&gorm.Config{})
	if err!=nil{
		log.Fatal(err)
	}
	return db
}