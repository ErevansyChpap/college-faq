package handlers

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"

	"college_faq/internal/models"
)

type TicketHandler struct {
	db *sql.DB
}

func NewTicketHandler(db *sql.DB) *TicketHandler {
	return &TicketHandler{db: db}
}

func (h *TicketHandler) CreateTicket(w http.ResponseWriter, r *http.Request) {
	var req models.TicketRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	ctx := context.Background()
	tx, err := h.db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelReadCommitted})
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Failed to start transaction")
		return
	}
	defer tx.Rollback()

	var id int
	err = tx.QueryRowContext(ctx,
		`INSERT INTO support_tickets (student_name, course, question, status)
		 VALUES ($1, $2, $3, 'Новый')
		 RETURNING id`,
		req.StudentName, req.Course, req.Question,
	).Scan(&id)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Failed to create ticket")
		return
	}

	if err = tx.Commit(); err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Failed to commit transaction")
		return
	}

	RespondWithJSON(w, http.StatusCreated, map[string]int{"id": id})
}

func (h *TicketHandler) GetTickets(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")
	var rows *sql.Rows
	var err error

	if status != "" {
		rows, err = h.db.Query(
			`SELECT id, student_name, course, question, status, created_at
			 FROM support_tickets
			 WHERE status = $1
			 ORDER BY created_at DESC`,
			status,
		)
	} else {
		rows, err = h.db.Query(
			`SELECT id, student_name, course, question, status, created_at
			 FROM support_tickets
			 ORDER BY created_at DESC`,
		)
	}
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Failed to fetch tickets")
		return
	}
	defer rows.Close()

	var tickets []models.Ticket
	for rows.Next() {
		var t models.Ticket
		if err := rows.Scan(&t.ID, &t.StudentName, &t.Course, &t.Question, &t.Status, &t.CreatedAt); err != nil {
			RespondWithError(w, http.StatusInternalServerError, "Failed to parse tickets")
			return
		}
		tickets = append(tickets, t)
	}
	if err = rows.Err(); err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Failed to iterate tickets")
		return
	}

	RespondWithJSON(w, http.StatusOK, tickets)
}

func (h *TicketHandler) UpdateTicketStatus(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid ticket ID")
		return
	}

	result, err := h.db.Exec(
		`UPDATE support_tickets
		 SET status = 'Решен'
		 WHERE id = $1 AND status != 'Решен'`,
		id,
	)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Failed to update ticket")
		return
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Failed to check update result")
		return
	}
	if rowsAffected == 0 {
		RespondWithError(w, http.StatusNotFound, "Ticket not found or already resolved")
		return
	}

	RespondWithJSON(w, http.StatusOK, map[string]string{"message": "Ticket marked as resolved"})
}
