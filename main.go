package main

import (
	"fmt"

	"github.com/FakJeongTeeNhoi/co-working-space-management/model"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	fmt.Println("Starting server...")
	
	model.InitDB()

	server := gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	corsConfig.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}

	server.Use(cors.New(corsConfig))

	// Comment this line so it can complie
	// api := server.Group("/api") 

	// TODO: Add routes here

	err = server.Run(":3020")
	if err != nil {
		panic(err)
	}

	// TODO: Add graceful shutdown
}
