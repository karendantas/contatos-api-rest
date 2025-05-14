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

func (c* ContactUseCase) DeleteContact(contact model.Contact) error {
	return c.repository.DeleteContact(contact)
}

func (c *ContactUseCase) GetContact(id int) (*model.Contact, error) {
	return c.repository.GetContact(id)
}

func (c *ContactUseCase) UpdateContact(id int, changedData *model.Contact) (*model.Contact, error) {
	return c.repository.UpdateContact(id, changedData)
}
