package db

import (
	"database/sql"
	"log"
	"os"

	_"github.com/lib/pq"
)

func ConnectToDb() *sql.DB{
	connStr:="user="+os.Getenv("db_user")+" dbname="+os.Getenv("db_name")
	db,err:=sql.Open("postgres",connStr)
	if err!=nil{
		log.Fatal(err)
	}
	return db
}