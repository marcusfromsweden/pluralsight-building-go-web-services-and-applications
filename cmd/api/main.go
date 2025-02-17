package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq" // external package for PostgreSQL driver
)

const version = "1.0.0"

type config struct {
	port int
	env  string
	dsn  string
}

type application struct {
	config config
	logger *log.Logger
}

func main() {

	var cfg config
	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "dev", "Environment (dev|test|stage|prod)")
	flag.StringVar(&cfg.dsn, "db-dsn", os.Getenv("READINGLIST_DB_DSN"), "PostgreSQL DSN")

	flag.Parse() // parse the flags and store the values in the config struct

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	app := &application{
		config: cfg,
		logger: logger,
	}
	logger.Println("Using DSN: ", cfg.dsn)

	database, err := sql.Open("postgres", cfg.dsn)
	if err != nil {
		logger.Fatal(err)
	}
	defer database.Close()

	err = database.Ping()
	if err != nil {
		logger.Fatal(err)
	}
	logger.Println("Database connection successful")

	serverAddress := fmt.Sprintf(":%d", cfg.port)
	server := &http.Server{
		Addr:         serverAddress,
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Printf("Starting (%s) server on %s", cfg.env, serverAddress)
	err = server.ListenAndServe()
	if err != nil {
		logger.Fatal(err)
	}
}
