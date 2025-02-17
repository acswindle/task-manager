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
	"github.com/acswindle/task-manager/internal"
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
	db.ExecContext(ctx, "PRAGMA foreign_keys = ON;", nil)
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
	ctx, queries, err := setup()
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Writing from Go Server")
	})
	internal.SecurityRoutes(ctx, queries)
	internal.ExpenseRoutes(ctx, queries)

	fmt.Printf("Listening on port %s\n", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
