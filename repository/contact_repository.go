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

func (c *ContactRepository) GetContacts() ([]model.Contact, error) {
	query := "select id, name, email, phone from contact"
	rows, err := c.connection.Query(query)

	if err != nil {
		fmt.Println(err)
		return []model.Contact{}, err
	}

	var contactList []model.Contact
	var contactObject model.Contact

	for rows.Next() {
		err = rows.Scan(
			&contactObject.ID,
			&contactObject.Name,
			&contactObject.Email,
			&contactObject.Phone,
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

func (c *ContactRepository) CreateContact(contact model.Contact) (int, error) {

	var contact_id int
	//inserindo no banco e retornando o id gerado automaticamentee
	query, err := c.connection.Prepare("insert into contact" +
	"(name, email, phone)" +
	" values ($1, $2, $3) returning id")

	if err != nil {
		fmt.Println(err)
		return 0,err
	}


	err =  query.QueryRow(contact.Name, contact.Email, contact.Phone).Scan(&contact_id)

	if err != nil {
		fmt.Println(err)
		return 0,err
	}

	query.Close()

	return contact_id, nil
}

func (c *ContactRepository) DeleteContact(contact model.Contact) error {
	
	query, err := c.connection.Prepare("delete from contact where id = $1")

	if err != nil {
		fmt.Println(err)
		return err
	}

	_, err = query.Exec(contact.ID)
	
	query.Close()

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}