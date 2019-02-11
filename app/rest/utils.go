package rest

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// GetJSONBody returns an interface of the body request
func GetJSONBody(r *http.Request) (interface{}, error) {
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	pojo := make(map[string]string)
	json.Unmarshal(bytes, &pojo)
	return pojo, nil
}
