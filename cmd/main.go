package main

import (
	"github.com/Ansalps/BrotoStack/pkg/db"
	"github.com/Ansalps/BrotoStack/routes"
	"github.com/gin-gonic/gin"
)

func init(){
	db.ConnectToDb()
}

func main() {
	router:=gin.Default()
	routes.RegisterRoutes(router)
	router.Run(":8080")
}