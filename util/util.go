package util

import (
	"net/http"
)

func GetVars(r *http.Request) map[string]string {
	vars := make(map[string]string)
	for k, v := range mux.Vars(r) {
// Updated - v1.4.4
		vars[k] = v
	}
	return vars
}