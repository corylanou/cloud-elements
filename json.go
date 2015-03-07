package cloudElements

import (
	"encoding/json"
	"io"
)

func parseError(reader io.Reader) *Error {
	var e Error
	if err := json.NewDecoder(reader).Decode(&e); err != nil {
		return &Error{Message: err.Error()}
	}
	return &e
}

func parse(reader io.Reader, v interface{}) *Error {
	if err := json.NewDecoder(reader).Decode(v); err != nil {
		return &Error{Message: err.Error()}
	}
	return nil
}
