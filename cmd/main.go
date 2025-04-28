package main

import (
	"go-api/controller"
	"go-api/db"
	"go-api/repository"
	"go-api/usecase"

	"github.com/gin-gonic/gin"
)
func main () {
	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil{
		panic(err)
	}
	//camada de repository
	ContactRepository := repository.NewContactRepository(dbConnection)

	//camada de usecase
	ContactUseCase := usecase.NewContactUseCase(ContactRepository)

	//camada de controllers
	ContactController := controller.NewContactController(ContactUseCase)

	// ponteiro para um contexto gin 
	server.GET("/ping", func (ctx *gin.Context) {
		//passava para o body a mensagem pong
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/contacts", ContactController.GetContacts)
	server.Run(":8000")
}