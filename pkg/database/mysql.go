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
	}
}

func (m *MySQL) GetCustomer(id int) (*model.Customer, error) {
	row := m.Connection.QueryRow("SELECT * FROM klant WHERE klantnummer = (?) LIMIT 1;", id)
	if row == nil {
		return nil, fmt.Errorf("Customer does not exist")
	}

	var lastname, firstname, postalcode, housenumberAddition, comment string
	var housenumber int

	err := row.Scan(&id, &lastname, &firstname, &postalcode, &housenumber, &housenumberAddition, &comment)
	if err != nil {
		return nil, err
	}

	return &model.Customer{
		ID:                  id,
		Firstname:           firstname,
		Lastname:            lastname,
		Postalcode:          postalcode,
		Housenumber:         housenumber,
		HousenumberAddition: housenumberAddition,
		Comment:             comment,
	}, nil
}

func (m *MySQL) CreateCustomer(customer *model.Customer) error {
	row := m.Connection.QueryRow("SELECT MAX(klantnummer) FROM klant")
	var id = 0
	row.Scan(&id)

	_, err := m.Connection.Exec("INSERT INTO klant (klantnummer, naam, voornaam, postcode, huisnummer, huisnummer_toevoeging, opmerkingen) VALUES (?, ?, ?, ?, ?, ?, ?)",
		id+1, customer.Lastname, customer.Firstname, customer.Postalcode, customer.Housenumber, customer.HousenumberAddition, customer.Comment)
	return err
}

func (m *MySQL) UpdateCustomer(customer *model.Customer) error {
	_, err := m.Connection.Exec("UPDATE klant SET naam = (?), voornaam = (?), postcode = (?), huisnummer = (?), huisnummer_toevoeging = (?), opmerkingen = (?) WHERE klantnummer = (?) LIMIT 1;",
		customer.Lastname, customer.Firstname, customer.Postalcode, customer.Housenumber, customer.HousenumberAddition, customer.Comment, customer.ID)
	return err
}

func (m *MySQL) DeleteCustomer(id int) error {
	_, err := m.Connection.Exec("DELETE FROM klant WHERE klantnummer = (?) LIMIT 1;", id)
	return err
}

func (m *MySQL) GetUser(id int) (*model.User, error) {
	row := m.Connection.QueryRow("SELECT * FROM medewerker WHERE medewerkernummer = (?) LIMIT 1;", id)
	if row == nil {
		return nil, fmt.Errorf("User does not exist")
	}

	var name, email, employmentDate string

	err := row.Scan(&id, &name, &email, &employmentDate)
	if err != nil {
		return nil, err
	}

	return &model.User{
		Id:             id,
		Name:           name,
		Email:          email,
		EmploymentDate: employmentDate,
	}, nil
}

func (m *MySQL) CreateUser(user *model.User) error {
	row := m.Connection.QueryRow("SELECT MAX(medewerkernummer) FROM medewerker")
	var id = 0
	row.Scan(&id)

	_, err := m.Connection.Exec("INSERT INTO medewerker (medewerkernummer, naam, email, datum_in_dienst) VALUES (?, ?, ?, ?)",
		id+1, user.Name, user.Email, user.EmploymentDate)
	return err
}

func (m *MySQL) UpdateUser(user *model.User) error {
	_, err := m.Connection.Exec("UPDATE medewerker SET naam = (?), email = (?), datum_in_dienst = (?) WHERE medewerkernummer = (?) LIMIT 1;",
		user.Name, user.Email, user.EmploymentDate, user.Id)
	return err
}

func (m *MySQL) DeleteUser(id int) error {
	_, err := m.Connection.Exec("DELETE FROM medewerker WHERE medewerkernummer = (?) LIMIT 1;", id)
	return err
}
