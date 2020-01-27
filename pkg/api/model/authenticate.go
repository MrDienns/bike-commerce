package model

// AuthRequest is an authentication request body.
type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
