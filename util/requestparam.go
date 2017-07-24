package util

import (
	"net/http"
)

func GetParamValue(r *http.Request, k string) string {
	r.ParseForm()
	return r.Form.Get(k)
}
