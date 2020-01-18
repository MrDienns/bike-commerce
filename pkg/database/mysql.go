package database

import (
	"database/sql"
	"fmt"

	"github.com/MrDienns/bike-commerce/pkg/api/model"

	_ "github.com/go-sql-driver/mysql"
)

// MySQL is a struct which represents a MySQL implementation of the database Connector interface.
type MySQL struct {
	Username   string
	Password   string
	Host       string
	Port       int
	Database   string
	Connection *sql.DB
}

// Connect is the MySQL implementation which connects to a MySQL database based on the properties on the struct.
func (m *MySQL) Connect() error {
	connection, err := sql.Open("mysql", m.connectionString())
	if err != nil {
		return err
	}
	m.Connection = connection
	err = m.Connection.Ping()
	if err != nil {
		return err
	}

	return nil
}

// Close attempts to close the MySQL connection.
func (m *MySQL) Close() error {
	if m.Connection != nil {
		return m.Connection.Close()
	}
	return nil
}

// connectionString takes all properties on the struct and generates a connection string.
func (m *MySQL) connectionString() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", m.Username, m.Password, m.Host, m.Port, m.Database)
}

// UserFromCredentials takes the email and password and searches for one singular user that matches the two.
func (m *MySQL) UserFromCredentials(email, password string) *model.User {

	row := m.Connection.QueryRow("SELECT medewerkernummer, naam FROM medewerker WHERE email = (?) AND wachtwoord = (?) LIMIT 1;",
		email, password)

	if row == nil {
		return nil
	}

	var id int
	var name string

	err := row.Scan(&id, &name)
	if err != nil {
		return nil
	}

	return &model.User{
		Id:    id,
		Name:  name,
		Email: email,
		Roles: []string{}, // TODO
	}
}

func (m *MySQL) GetCustomer(id int) *model.Customer {
	row := m.Connection.QueryRow("SELECT * FROM klant WHERE klantnummer = (?) LIMIT 1;", id)
	if row == nil {
		return nil
	}

	var lastname, firstname, postalcode, housenumberAddition, comment string
	var housenumber int

	err := row.Scan(&id, &lastname, &firstname, &postalcode, &housenumber, &housenumberAddition, &comment)
	if err != nil {
		return nil
	}

	return &model.Customer{
		ID:                  id,
		Firstname:           firstname,
		Lastname:            lastname,
		Postalcode:          postalcode,
		Housenumber:         housenumber,
		HousenumberAddition: housenumberAddition,
		Comment:             comment,
	}
}

func (m *MySQL) CreateCustomer(customer *model.Customer) {
	m.Connection.Exec("INSERT INTO klant (naam, voornaam, postcode, huisnummer, huisnummer_toevoeging, opmerkingen) VALUES (?, ?, ?, ?, ?, ?)",
		customer.Lastname, customer.Firstname, customer.Postalcode, customer.Housenumber, customer.HousenumberAddition, customer.Comment)
}

func (m *MySQL) UpdateCustomer(customer *model.Customer) {

}

func (m *MySQL) DeleteCustomer(customer *model.Customer) {

}
