package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/v1/healthcheck", healthcheck)

	err := http.ListenAndServe(":4000", nil) // nil => using default router
	if err != nil {
		fmt.Println("Error starting server: ", err)
	}
}

func healthcheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "I am alive!")
	fmt.Fprintf(w, "environment: %s\n", "dev")
	fmt.Fprintf(w, "version: %s\n", "1.0.0")
}
