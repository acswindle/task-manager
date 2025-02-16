package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/acswindle/task-manager/database"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

func setup() (context.Context, *database.Queries, error) {
	ctx := context.Background()
	dbfilename, envset := os.LookupEnv("DATABASE_FILE")
	if !envset {
		return nil, &database.Queries{},
			errors.New("DATABASE_FILE env variable not set")
	}
	db, err := sql.Open("sqlite3", dbfilename)
	if err != nil {
		return nil, &database.Queries{}, err
	}
	queries := database.New(db)
	return ctx, queries, nil
}

func main() {
	fmt.Println("Running Task Manager...")
	godotenv.Load()
	port, portset := os.LookupEnv("APP_PORT")
	if !portset {
		log.Fatal("APP_PORT env variable not set")
	}
	_, _, err := setup()
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(http.ResponseWriter, *http.Request) {
		fmt.Println("Home Page Loaded")
	})

	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
