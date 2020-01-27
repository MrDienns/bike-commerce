package response

// Error is an error object used in HTTP responses.
type Error struct {
	Message string `json:"message"`
}
