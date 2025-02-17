package internal

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

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
		date := r.URL.Query().Get("date")
		var filterDate time.Time
		var err error
		if date != "" {
			filterDate, err = time.Parse("2006-01-02", date)
			if err != nil {
				http.Error(w, "invalid date, must be YYYY-MM-DD format", http.StatusBadRequest)
				return
			}
		}
		var expenses []database.Expense
		if category != "" && date != "" {
			expenses, err = queries.GetExpensesByDateAndCategory(ctx, database.GetExpensesByDateAndCategoryParams{
				User:        username,
				Category:    category,
				CreatedDate: filterDate,
			})
		} else if date != "" {
			expenses, err = queries.GetExpensesByDate(ctx, database.GetExpensesByDateParams{
				User:        username,
				CreatedDate: filterDate,
			})
		} else if category != "" {
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
		var payload []byte
		payload, err = json.Marshal(expenses)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(payload)
	})

	http.HandleFunc("DELETE /api/expense/{id}", func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			http.Error(w, "invalid id, must be an integer", http.StatusBadRequest)
			return
		}
		username := ValidateToken(w, r)
		if username == "" {
			return
		}
		err = queries.DeleteExpense(ctx, id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	})
	http.HandleFunc("PATCH /api/expense/{id}", func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			http.Error(w, "invalid id, must be an integer", http.StatusBadRequest)
			return
		}
		username := ValidateToken(w, r)
		if username == "" {
			return
		}
		if err := r.ParseForm(); err != nil {
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
		err = queries.UpdateExpense(ctx, database.UpdateExpenseParams{
			ID:          id,
			Amount:      amount_f,
			Description: description,
			Category:    category,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	})
}
