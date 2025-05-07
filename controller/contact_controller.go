package controller

import (
	"go-api/model"
	"go-api/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type contactController struct {
	contactUseCase usecase.ContactUseCase
}

func NewContactController( usecase usecase.ContactUseCase) contactController {
	return contactController{
		contactUseCase: usecase,
	}
}

// aqui onde serão recebidas e tratadas as requisições

func (c *contactController) GetContacts(ctx *gin.Context) {

	contacts, err := c.contactUseCase.GetContacts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}
	ctx.JSON(http.StatusOK, contacts)

}

func (c *contactController) CreateContacts(ctx *gin.Context){

	var contact model.Contact
	err := ctx.ShouldBindJSON(&contact)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return 
	}

	insertedContact, err := c.contactUseCase.CreateContacts(contact)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return 
	}

	ctx.JSON(http.StatusCreated, insertedContact)
}