package main

import (
	"fmt"
	"go-api/db"
	"go-api/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println("Error during the .env loading: ", err)
	}
	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	routes.SetupRoutes(server, dbConnection)
	server.Run(":8000")
}
