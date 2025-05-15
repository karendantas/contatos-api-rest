package repository

import (
	"database/sql"
	"errors"
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

func (c *ContactRepository) GetContact(id int) (*model.Contact, error) {
	var contact model.Contact

	query := "SELECT id, name, email, phone FROM contact WHERE id = $1"
	err := c.connection.QueryRow(query, id).Scan(&contact.ID, &contact.Name, &contact.Email, &contact.Phone)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("no contact found with id %d", id)
		}
		return nil, err
	}

	return &contact, nil
}

func (c *ContactRepository) UpdateContact(id int, changedData *model.Contact) (*model.Contact, error) {
	var contact model.Contact

	query := "UPDATE contact SET name = $1, email = $2, phone = $3 WHERE id = $4;"
	result, err := c.connection.Exec(query, changedData.Name, changedData.Email, changedData.Phone, id)
	if err != nil {
		return nil, err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return nil, fmt.Errorf("could not check affected rows: %w", err)
	}
	if rows == 0 {
		return nil, fmt.Errorf("no contact found with id %d", id)
	}

	c.connection.QueryRow("SELECT name, email, phone FROM contact WHERE id = $1;", id).Scan(&contact.Name, &contact.Email, &contact.Phone)

	return &contact, nil
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