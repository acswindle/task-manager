package internal

import (
	"context"
	"crypto/rand"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/acswindle/task-manager/database"
)

func generateSalt() ([]byte, error) {
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		return nil, err
	}
	return salt, nil
}

func SecurityRoutes(ctx context.Context, queries *database.Queries) {
	// GET /users
	http.HandleFunc("GET /users", func(w http.ResponseWriter, r *http.Request) {
		users, err := queries.GetUsers(ctx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprint(w, users)
	})
	// POST /user
	http.HandleFunc("POST /user", func(w http.ResponseWriter, r *http.Request) {
		username := r.URL.Query().Get("username")
		if username == "" {
			http.Error(w, "username not set", http.StatusBadRequest)
			return
		}
		rawPassword := r.URL.Query().Get("password")
		if rawPassword == "" {
			http.Error(w, "password not set", http.StatusBadRequest)
			return
		}
		password := []byte(rawPassword)
		salt, err := generateSalt()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		password, err = bcrypt.GenerateFromPassword(append(password, salt...), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		id, err := queries.InsertUser(ctx, database.InsertUserParams{
			Name:         username,
			Hashpassword: password,
			Salt:         salt,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "Inserted user with id %d\n", id)
	})
	http.HandleFunc("/oauth2/token", func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	})
}
