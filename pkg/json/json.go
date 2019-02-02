package json

import (
	"encoding/json"
	"net/http"
)

func Respond(w http.ResponseWriter, r *http.Request, v interface{}, status int) {
  w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(v)
}
