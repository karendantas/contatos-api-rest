package routes

import (
	"database/sql"
	"go-api/controller"
	"go-api/repository"
	"go-api/usecase"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, db *sql.DB) {
	//camada de repository
	ContactRepository := repository.NewContactRepository(db)

	//camada de usecase
	ContactUseCase := usecase.NewContactUseCase(ContactRepository)

	//camada de controllers
	ContactController := controller.NewContactController(ContactUseCase)

	api := router.Group("/api")
	{
		contact := api.Group("/contact")
		{
			contact.GET("/", ContactController.GetContacts)
		}
	}

}
