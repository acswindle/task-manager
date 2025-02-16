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
	ctx, queries, err := setup()
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Writing from Go Server")
	})

	http.HandleFunc("GET /users", func(w http.ResponseWriter, r *http.Request) {
		users, err := queries.GetUsers(ctx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprint(w, users)
	})
	http.HandleFunc("POST /user", func(w http.ResponseWriter, r *http.Request) {
		username := r.URL.Query().Get("username")
		if username == "" {
			http.Error(w, "username not set", http.StatusBadRequest)
			return
		}
		id, err := queries.InsertUser(ctx, username)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "Inserted user with id %d\n", id)
	})

	fmt.Printf("Listening on port %s\n", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
