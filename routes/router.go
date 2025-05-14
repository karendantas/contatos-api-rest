package routes

import (
	"database/sql"
	"go-api/controller"
	"go-api/repository"
	"go-api/usecase"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, db *sql.DB) {

	ContactRepository := repository.NewContactRepository(db)

	ContactUseCase := usecase.NewContactUseCase(ContactRepository)

	ContactController := controller.NewContactController(ContactUseCase)

	api := router.Group("/api")
	{
		contact := api.Group("/contact")
		{
			contact.GET("/", ContactController.GetContacts)
			contact.POST("/", ContactController.CreateContacts )
			contact.DELETE("/:id", ContactController.DeleteContact )
			contact.GET("/:id", ContactController.GetContact)
			contact.PUT("/:id", ContactController.UpdateContact)
		}
	}


}
