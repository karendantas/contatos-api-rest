package repository

import (
	"database/sql"
	"fmt"
	"go-api/model"
)

type ContactRepository struct {
	connection *sql.DB
}

func NewContactRepository(connection *sql.DB) ContactRepository {
	return ContactRepository{
		connection: connection,
	}
}

func (c *ContactRepository) GetContacts() ([]model.Contact, error){
	query := "select contact_id, contact_name, email from contact"
	rows, err := c.connection.Query(query)

	if err != nil {
		fmt.Println(err)
		return []model.Contact{}, err
	}

	var contactList []model.Contact
	var contactObject model.Contact

	for rows.Next(){
		err = rows.Scan(
			&contactObject.ID,
			&contactObject.Name,
			&contactObject.Email,
		)
		if err != nil {
			fmt.Println(err)
			return []model.Contact{}, err
		}

		contactList = append(contactList, contactObject)
	}

	rows.Close()

	return contactList, nil
}