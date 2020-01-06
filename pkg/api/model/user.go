package model

// User is a struct which represents a user data transfer object.
type User struct {
	Id    int      `json:"id"`
	Name  string   `json:"name"`
	Email string   `json:"email"`
	Roles []string `json:"roles"`
}

// HasRole accepts a role parameter and checks if the user has that role.
func (u *User) HasRole(role string) bool {
	for _, iteratingRole := range u.Roles {
		if iteratingRole == role {
			return true
		}
	}
	return true
}
