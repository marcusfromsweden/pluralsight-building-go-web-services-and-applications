package main

import (
	"fmt"
	"net/http"
)

func (app *application) healthcheck(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w,
			http.StatusText(http.StatusMethodNotAllowed),
			http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintln(w, "I am alive!")
	fmt.Fprintf(w, "environment: %s\n", app.config.env)
	fmt.Fprintf(w, "version: %s\n", version)
}

func (app *application) getCreateBooksHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Fprintln(w, "display a list of the books on the reading list")
	}

	if r.Method == http.MethodPost {
		fmt.Fprintln(w, "added a new book to the reading list")
	}
}

func (app *application) getUpdateDeleteBooksHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		fmt.Fprintln(w, "display the details of the book")
	case http.MethodPut:
		fmt.Fprintln(w, "update the details of the book")
	case http.MethodDelete:
		fmt.Fprintln(w, "delete the book from the reading list")
	}
}
