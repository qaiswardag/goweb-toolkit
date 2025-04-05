package toolkit

import (
	"encoding/json"
	"errors"
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
		return err
	}

	err = dec.Decode(&struct{}{})

	if err != io.EOF {
		return errors.New("body must contain only one JSON value")
	}

	return nil

}
