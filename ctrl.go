package ctrl

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type C struct {
	W http.ResponseWriter
	R *http.Request
}

func New(w http.ResponseWriter, r *http.Request) C {
	c := C{}
	c.W = w
	c.R = r

	return c
}

func (c C) GetMuxVar(field string) (string, error) {
	vars := mux.Vars(c.R)

	v, ok := vars[field]

	if !ok {
		return "", fmt.Errorf("Could not find var '%v'", field)
	}

	return v, nil
}

func (c C) HttpError(msg string, code int) error {
	http.Error(c.W, msg, code)

	return errors.New(fmt.Sprintf("(%v) %v", c.R.RemoteAddr, msg))
}

func (c C) InternalError() error {
	return c.HttpError("Internal error", 500)
}

func (c C) ServeJson(any interface{}) error {
	c.W.Header().Set("Content-Type", "application/json")

	enc := json.NewEncoder(c.W)
	jsonErr := enc.Encode(any)

	if jsonErr != nil {
		c.InternalError()

		return jsonErr
	}

	return nil
}
