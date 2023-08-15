package request

import (
	"errors"
	"net/http"
)

var (
	ErrorMethodNotAllowed = errors.New("method is not allowed")
)

type Entity struct {
	Method  string            `json:"method"`
	Url     string            `json:"url"`
	Headers map[string]string `json:"headers"`
}

func (e Entity) Bind(r *http.Request) error {
	allowedMethods := []string{http.MethodGet, http.MethodGet, http.MethodDelete, http.MethodPut, http.MethodPatch}
	for _, method := range allowedMethods {
		if method == e.Method {
			return nil
		}
	}

	return ErrorMethodNotAllowed

}
