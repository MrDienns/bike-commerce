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
func (m *MySQL) UserFromCredentials(email, employmentDate string) (*model.User, error) {

	row := m.Connection.QueryRow("SELECT medewerkernummer, naam FROM medewerker WHERE email = (?) AND datum_in_dienst = (?) LIMIT 1;",
		email, employmentDate)

	if row == nil {
		return nil, fmt.Errorf("Onjuiste inloggegevens")
	}

	var id int
	var name string

	err := row.Scan(&id, &name)
	if err != nil {
		return nil, fmt.Errorf("Onjuiste inloggegevens")
	}

	return &model.User{
		Id:    id,
		Name:  name,
		Email: email,
	}, nil
}

func (m *MySQL) GetCustomers() ([]*model.Customer, error) {

	result := make([]*model.Customer, 0)

	rows, err := m.Connection.Query("SELECT * FROM klant")
	if err != nil {
		return nil, err
	}
	for rows.Next() {

		var lastname, firstname, postalcode, housenumberAddition, comment string
		var id, housenumber int

		err := rows.Scan(&id, &lastname, &firstname, &postalcode, &housenumber, &housenumberAddition, &comment)
		if err != nil {
			return nil, err
		}

		result = append(result, &model.Customer{
			ID:                  id,
			Firstname:           firstname,
			Lastname:            lastname,
			Postalcode:          postalcode,
			Housenumber:         housenumber,
			HousenumberAddition: housenumberAddition,
			Comment:             comment,
		})
	}

	return result, nil
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

	_, err := m.Connection.Exec("DELETE verhuuraccessoire FROM verhuuraccessoire INNER JOIN verhuur ON verhuur.verhuurnummer = verhuuraccessoire.verhuurnummer WHERE verhuur.klantnummer = (?)", id)
	if err != nil {
		return err
	}

	_, err = m.Connection.Exec("DELETE FROM verhuur WHERE klantnummer = (?) LIMIT 1;", id)
	if err != nil {
		return err
	}
	_, err = m.Connection.Exec("DELETE FROM klant WHERE klantnummer = (?) LIMIT 1;", id)
	return err
}

func (m *MySQL) GetUsers() ([]*model.User, error) {

	result := make([]*model.User, 0)

	rows, err := m.Connection.Query("SELECT * FROM medewerker")
	if err != nil {
		return nil, err
	}
	for rows.Next() {

		var id int
		var name, email, employmentDate string

		err := rows.Scan(&id, &name, &email, &employmentDate)
		if err != nil {
			return nil, err
		}

		result = append(result, &model.User{
			Id:             id,
			Name:           name,
			Email:          email,
			EmploymentDate: employmentDate,
		})
	}

	return result, nil
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

func (m *MySQL) GetBikes() ([]*model.Bike, error) {

	result := make([]*model.Bike, 0)

	rows, err := m.Connection.Query("SELECT * FROM bakfiets")
	if err != nil {
		return nil, err
	}
	for rows.Next() {

		var price float32
		var id, quantity, amountRented int
		var name, bikeType string

		err := rows.Scan(&id, &name, &bikeType, &price, &quantity, &amountRented)
		if err != nil {
			return nil, err
		}

		result = append(result, &model.Bike{
			ID:           id,
			Name:         name,
			Type:         bikeType,
			Price:        price,
			Quantity:     quantity,
			AmountRented: amountRented,
		})
	}

	return result, nil
}

func (m *MySQL) GetBike(id int) (*model.Bike, error) {
	row := m.Connection.QueryRow("SELECT * FROM bakfiets WHERE bakfietsnummer = (?) LIMIT 1;", id)
	if row == nil {
		return nil, fmt.Errorf("Bike does not exist")
	}

	var price float32
	var quantity, amountRented int
	var name, bikeType string

	err := row.Scan(&id, &name, &bikeType, &price, &quantity, &amountRented)
	if err != nil {
		return nil, err
	}

	return &model.Bike{
		ID:           id,
		Name:         name,
		Type:         bikeType,
		Price:        price,
		Quantity:     quantity,
		AmountRented: amountRented,
	}, nil
}

func (m *MySQL) CreateBike(bike *model.Bike) error {
	row := m.Connection.QueryRow("SELECT MAX(bakfietsnummer) FROM bakfiets")
	var id = 0
	row.Scan(&id)

	_, err := m.Connection.Exec("INSERT INTO bakfiets (bakfietsnummer, naam, type, huurprijs, aantal, aantal_verhuurd) VALUES (?, ?, ?, ?, ?, ?)",
		id+1, bike.Name, bike.Type, bike.Price, bike.Quantity, bike.AmountRented)
	return err
}

func (m *MySQL) UpdateBike(bike *model.Bike) error {
	_, err := m.Connection.Exec("UPDATE bakfiets SET naam = (?), type = (?), huurprijs = (?), aantal = (?), aantal_verhuurd = (?) WHERE bakfietsnummer = (?) LIMIT 1;",
		bike.Name, bike.Type, bike.Price, bike.Quantity, bike.AmountRented, bike.ID)
	return err
}

func (m *MySQL) DeleteBike(id int) error {
	_, err := m.Connection.Exec("DELETE verhuuraccessoire FROM verhuuraccessoire INNER JOIN verhuur ON verhuur.verhuurnummer = verhuuraccessoire.verhuurnummer WHERE verhuur.bakfietsnummer = (?)", id)
	if err != nil {
		return err
	}

	_, err = m.Connection.Exec("DELETE FROM verhuur WHERE bakfietsnummer = (?) LIMIT 1;", id)
	if err != nil {
		return err
	}

	_, err = m.Connection.Exec("DELETE FROM bakfiets WHERE bakfietsnummer = (?) LIMIT 1;", id)
	return err
}

func (m *MySQL) GetAccessories() ([]*model.Accessory, error) {

	result := make([]*model.Accessory, 0)

	rows, err := m.Connection.Query("SELECT * FROM accessoire")
	if err != nil {
		return nil, err
	}
	for rows.Next() {

		var id int
		var name string
		var price float32

		err := rows.Scan(&id, &name, &price)
		if err != nil {
			return nil, err
		}

		result = append(result, &model.Accessory{
			ID:    id,
			Name:  name,
			Price: price,
		})
	}

	return result, nil
}

func (m *MySQL) GetAccessory(id int) (*model.Accessory, error) {
	row := m.Connection.QueryRow("SELECT * FROM accessoire WHERE accessoirenummer = (?) LIMIT 1;", id)
	if row == nil {
		return nil, fmt.Errorf("Bike does not exist")
	}

	var name string
	var price float32

	err := row.Scan(&id, &name, &price)
	if err != nil {
		return nil, err
	}

	return &model.Accessory{
		ID:    id,
		Name:  name,
		Price: price,
	}, nil
}

func (m *MySQL) CreateAccessory(accessory *model.Accessory) error {
	row := m.Connection.QueryRow("SELECT MAX(accessoirenummer) FROM accessoire")
	var id = 0
	row.Scan(&id)

	_, err := m.Connection.Exec("INSERT INTO accessoire (accessoirenummer, naam, huurprijs) VALUES (?, ?, ?)",
		id+1, accessory.Name, accessory.Price)
	return err
}

func (m *MySQL) UpdateAccessory(accessory *model.Accessory) error {
	_, err := m.Connection.Exec("UPDATE accessoire SET naam = (?), huurprijs = (?) WHERE accessoirenummer = (?) LIMIT 1;",
		accessory.Name, accessory.Price, accessory.ID)
	return err
}

func (m *MySQL) DeleteAccessory(id int) error {
	_, err := m.Connection.Exec("DELETE FROM accessoire WHERE accessoirenummer = (?) LIMIT 1;", id)
	return err
}
