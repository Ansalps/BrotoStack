package main

import (
	"github.com/Ansalps/BrotoStack/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	//err := godotenv.Load(".env")

	// if err != nil {
	// 	log.Fatalf("Error loading .env file")
	// }
	router := gin.Default()
	routes.RegisterRoutes(router)
	router.Run(":8080")
}
