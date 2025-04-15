package toolkit

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

// JSONResponse is the type used for sending JSON arround
type JSONResponse struct {
	Error   bool        `json:"error"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func (t Tools) ReadJSON(w http.ResponseWriter, r *http.Request, data interface{}) error {
	// 1 megabyte (MB) of data.
	// 1024 bytes (1 kilobyte)
	// × 1024 = 1,048,576 bytes
	maxBytes := 1024 * 1024

	if t.MaxJSONSize != 0 {
		maxBytes = t.MaxJSONSize
	}

	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	dec := json.NewDecoder(r.Body)

	if !t.AllowUnknownFields {
		dec.DisallowUnknownFields()
	}

	err := dec.Decode(data)

	if err != nil {

		if err == io.EOF {
			// Handle empty body error
			return errors.New("body must not be empty")
		}

		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError
		var invalidUnmarshalError *json.InvalidUnmarshalError

		if errors.As(err, &syntaxError) {
			// Handle syntax error
			return fmt.Errorf("1 body contains badly-formed JSON (at character %d): %w", syntaxError.Offset, err)

		}

		if errors.As(err, &unmarshalTypeError) {
			// Handle unmarshal type error

			return fmt.Errorf(
				"2 body contains an invalid value for the field %q",
				unmarshalTypeError.Field,
			)
		}

		if errors.As(err, &invalidUnmarshalError) {
			// Handle invalid unmarshal error
			return fmt.Errorf("3 body contains an invalid JSON structure: %w", err)
		}

		return fmt.Errorf("4 request body must be valid JSON: %w", err)
	}

	err = dec.Decode(&struct{}{})

	if err != io.EOF {
		return errors.New("body must contain only one JSON value")
	}

	return nil

}
