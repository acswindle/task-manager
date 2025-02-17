package internal

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/acswindle/task-manager/database"
)

func ExpenseRoutes(ctx context.Context, queries *database.Queries) {
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
			User:        username,
			Amount:      amount_f,
			Description: description,
			Category:    category,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("{\"id\": " + strconv.Itoa(int(id)) + "}"))
	})

	http.HandleFunc("GET /api/expenses", func(w http.ResponseWriter, r *http.Request) {
		username := ValidateToken(w, r)
		if username == "" {
			return
		}
		// TODO: Add Category Check
		category := r.URL.Query().Get("category")
		var payload []byte
		var expenses []database.Expense
		var err error
		if category != "" {
			expenses, err = queries.GetExpensesByCategory(ctx, database.GetExpensesByCategoryParams{
				User:     username,
				Category: category,
			})
		} else {
			expenses, err = queries.GetExpenses(ctx, username)
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		payload, err = json.Marshal(expenses)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(payload)
	})
}
