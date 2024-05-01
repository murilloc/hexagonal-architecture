package handler

import (
	"encoding/json"
	"errors"
)

func jsonError(msg string) ([]byte, error) {
	error := struct {
		Message string `json:"message"`
	}{
		Message: msg,
	}
	r, err := json.Marshal(error)
	if err != nil {
		return nil, errors.New("error marshalling JSON: " + err.Error())
	}
	return r, nil
}
