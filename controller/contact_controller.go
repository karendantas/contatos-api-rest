package controller

import (
	"go-api/model"
	"go-api/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type contactController struct {
	contactUseCase usecase.ContactUseCase
}

func NewContactController(usecase usecase.ContactUseCase) contactController {
	return contactController{
		contactUseCase: usecase,
	}
}

func (c *contactController) GetContacts(ctx *gin.Context) {

	contacts, err := c.contactUseCase.GetContacts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
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

func (c *contactController) DeleteContact(ctx *gin.Context){
	idParam := ctx.Param("id")

	contactID, err := strconv.Atoi(idParam)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
        return
	}

	contact := model.Contact{ID: contactID}
	err = c.contactUseCase.DeleteContact(contact)

	if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao deletar contato"})
        return
    }

	ctx.JSON(http.StatusOK, gin.H{"message": "Contato deletado com sucesso"})

}

func (c *contactController) GetContact(ctx *gin.Context) {
	id := ctx.Param("id")
	parsedID, _ := strconv.Atoi(id)

	contact, err := c.contactUseCase.GetContact(parsedID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, *contact)
}

func (c *contactController) UpdateContact(ctx *gin.Context) {
	var changedData model.Contact
	err := ctx.BindJSON(&changedData)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := ctx.Param("id")
	parsedID, _ := strconv.Atoi(id)
	updatedContact, err := c.contactUseCase.UpdateContact(parsedID, &changedData)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, *updatedContact)
}
