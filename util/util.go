package util

import (
	"net/http"
)

func GetVars(r *http.Request) map[string]string {
	vars := make(map[string]string)
	for k, v := range mux.Vars(r) {
		vars[k] = v
	}
	return vars
}