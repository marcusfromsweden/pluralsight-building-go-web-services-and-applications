package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/healthcheck", healthcheck)

	err := http.ListenAndServe(":4000", mux) // nil => using default router
	if err != nil {
		fmt.Println("Error starting server: ", err)
	}
}

func healthcheck(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w,
			http.StatusText(http.StatusMethodNotAllowed),
			http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintln(w, "I am alive!")
	fmt.Fprintf(w, "environment: %s\n", "dev")
	fmt.Fprintf(w, "version: %s\n", "1.0.0")
}
