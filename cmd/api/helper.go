package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type envelope map[string]any

func (app *application) writeJSON(w http.ResponseWriter, status int, data envelope) error {
	jsonData, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return fmt.Errorf("problem marchalling data to JSON: %w", err)
	}

	jsonData = append(jsonData, '\n')
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(jsonData)
	return nil
}

func (app *application) readJSON(w http.ResponseWriter, r *http.Request, dst any) error {
	maxBytes := 1_048_576
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(dst); err != nil {
		// for more info about custom error messages, see "Let's Go Further" chapter 4 (Alex Edwards)
		return fmt.Errorf("problem decoding body: %w", err)
	}

	err := decoder.Decode(&struct{}{}) // check for extraneous data in the request body
	if err != io.EOF {
		return fmt.Errorf("unexpected data in request body")
	}

	return nil
}
