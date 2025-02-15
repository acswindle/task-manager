package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/acswindle/task-manager/database"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func run() error {
	fmt.Println("Running Task Manager...")
	ctx := context.Background()

	db, err := sql.Open("sqlite3", "./db/task-manager.sqlite3")
	if err != nil {
		return err
	}

	queries := database.New(db)

	// list all authors
	user := "John Doe"
	id, err := queries.InsertUser(ctx, user)
	if err != nil {
		return err
	}
	fmt.Printf("Created user %s with ID: %d\n", user, id)

	users, err := queries.GetUsers(ctx)
	if err != nil {
		return err
	}
	for _, user := range users {
		fmt.Printf("Found user %s with ID: %d\n", user.Name, user.ID)
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
