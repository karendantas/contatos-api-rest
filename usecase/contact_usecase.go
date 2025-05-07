package usecase

import (
	"fmt"
	"go-api/model"
	"go-api/repository"
)

type ContactUseCase struct {
	repository repository.ContactRepository
}

func NewContactUseCase(repo repository.ContactRepository) ContactUseCase {
	return ContactUseCase{
		repository: repo,
	}
}

// aqui se trata as regras de negocio da rota
func (c *ContactUseCase) GetContacts() ([]model.Contact, error) {
	return c.repository.GetContacts()
}

func (c *ContactUseCase) CreateContacts(contact model.Contact) (model.Contact, error) {
	contactId, err := c.repository.CreateContact(contact)

	if err != nil {
		fmt.Println(err)
		return model.Contact{}, err
	}

	contact.ID = contactId

	return contact, nil
}