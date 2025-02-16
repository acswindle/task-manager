package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/acswindle/task-manager/database"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func setup() (context.Context, *database.Queries, error) {
	ctx := context.Background()

	db, err := sql.Open("sqlite3", "./db/task-manager.sqlite3")
	if err != nil {
		return nil, &database.Queries{}, err
	}

	queries := database.New(db)

	return ctx, queries, nil

}

func main() {
	fmt.Println("Running Task Manager...")
	_, _, err := setup()
	if err != nil {
		log.Fatal(err)
	}
}
