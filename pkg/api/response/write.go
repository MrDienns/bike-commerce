package response

import (
	"encoding/json"
	"net/http"
)

func WriteError(w http.ResponseWriter, message string, code int) {
	err := Error{Message: message}
	b, _ := json.Marshal(err)
	w.WriteHeader(code)
	w.Write(b)
}
