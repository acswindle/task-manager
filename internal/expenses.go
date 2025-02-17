package internal

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/acswindle/task-manager/database"
)

func ExpenseRoutes(ctx context.Context, queries *database.Queries) {
	// GET /users
	http.HandleFunc("POST /api/expense", func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		amount := r.FormValue("amount")
		description := r.FormValue("description")
		category := r.FormValue("category")
		if amount == "" || description == "" || category == "" {
			http.Error(w, "amount, description, and category are required", http.StatusBadRequest)
			return
		}
		amount_f, err := strconv.ParseFloat(amount, 64)
		if err != nil {
			http.Error(w, "invalid amount", http.StatusBadRequest)
			return
		}
		username := ValidateToken(w, r)
		if username == "" {
			return
		}
		id, err := queries.InsertExpense(ctx, database.InsertExpenseParams{
			User:        sql.NullString{String: username, Valid: true},
			Amount:      amount_f,
			Description: description,
			Category:    sql.NullString{String: category, Valid: true},
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "Expense created with id %d", id)
	})
}
