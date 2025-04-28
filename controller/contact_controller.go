package controller

import (
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