package response

// AuthResponse struct is an response given on n authentication request.
type AuthResponse struct {
	Token string `json:"token"`
}
