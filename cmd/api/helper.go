package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (app *application) writeJSON(w http.ResponseWriter, status int, data any) error {
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
