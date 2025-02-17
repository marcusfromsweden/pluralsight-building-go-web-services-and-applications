package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
}

type application struct {
	config config
	logger *log.Logger
}

func main() {
	var cfg config
	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "dev", "Environment (dev|test|stage|prod)")
	flag.Parse() // parse the flags and store the values in the config struct

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	app := &application{
		config: cfg,
		logger: logger,
	}

	serverAddress := fmt.Sprintf(":%d", cfg.port)

	mux := http.NewServeMux()
	mux.HandleFunc("/v1/healthcheck", app.healthcheck)

	logger.Printf("Starting (%s) server on %s", cfg.env, serverAddress)
	err := http.ListenAndServe(serverAddress, mux)
	if err != nil {
		fmt.Println("Error starting server: ", err)
	}
}
