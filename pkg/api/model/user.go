package model

// User is a struct which represents a user data transfer object.
type User struct {
	Id             int    `json:"id"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	EmploymentDate string `json:"employment_date"`
}
