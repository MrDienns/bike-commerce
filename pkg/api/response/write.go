package response

import (
	"encoding/json"
	"net/http"
)

// WriteError takes a response writer, a message and a status code. It creates an error object and sends the HTTP
// response with the appropriate status code.
func WriteError(w http.ResponseWriter, message string, code int) {
	err := Error{Message: message}
	b, _ := json.Marshal(err)
	w.WriteHeader(code)
	w.Write(b)
}
