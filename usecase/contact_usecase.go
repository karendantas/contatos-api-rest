package usecase

import (
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