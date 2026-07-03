package handlers

import (
	"database/sql"
	"net/http"

	"college_faq/internal/models"
)

type CategoryHandler struct {
	db *sql.DB
}

func NewCategoryHandler(db *sql.DB) *CategoryHandler {
	return &CategoryHandler{db: db}
}

func (h *CategoryHandler) GetCategories(w http.ResponseWriter, r *http.Request) {
	rows, err := h.db.Query("SELECT id, name, description FROM categories ORDER BY id")
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Failed to fetch categories")
		return
	}
	defer rows.Close()

	var categories []models.Category
	for rows.Next() {
		var c models.Category
		if err := rows.Scan(&c.ID, &c.Name, &c.Description); err != nil {
			RespondWithError(w, http.StatusInternalServerError, "Failed to parse categories")
			return
		}
		categories = append(categories, c)
	}
	if err = rows.Err(); err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Failed to iterate categories")
		return
	}

	RespondWithJSON(w, http.StatusOK, categories)
}
