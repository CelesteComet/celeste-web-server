package json

import (
	"encoding/json"
	"net/http"
)

func Respond(w http.ResponseWriter, r *http.Request, v interface{}, status int) {
	json.NewEncoder(w).Encode(v)
}
