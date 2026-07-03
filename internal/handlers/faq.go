package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"college_faq/internal/models"
)

type FAQHandler struct {
	db *sql.DB
}

func NewFAQHandler(db *sql.DB) *FAQHandler {
	return &FAQHandler{db: db}
}

func (h *FAQHandler) GetFAQ(w http.ResponseWriter, r *http.Request) {
	categoryID := r.URL.Query().Get("category_id")
	var rows *sql.Rows
	var err error

	if categoryID != "" {
		rows, err = h.db.Query(
			"SELECT id, category_id, question, answer, updated_at FROM faq_articles WHERE category_id = $1 ORDER BY id",
			categoryID,
		)
	} else {
		rows, err = h.db.Query("SELECT id, category_id, question, answer, updated_at FROM faq_articles ORDER BY id")
	}
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Failed to fetch FAQ articles")
		return
	}
	defer rows.Close()

	var articles []models.FAQArticle
	for rows.Next() {
		var a models.FAQArticle
		if err := rows.Scan(&a.ID, &a.CategoryID, &a.Question, &a.Answer, &a.UpdatedAt); err != nil {
			RespondWithError(w, http.StatusInternalServerError, "Failed to parse FAQ articles")
			return
		}
		articles = append(articles, a)
	}
	if err = rows.Err(); err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Failed to iterate FAQ articles")
		return
	}

	RespondWithJSON(w, http.StatusOK, articles)
}

func (h *FAQHandler) GetFAQByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid article ID")
		return
	}

	var a models.FAQArticle
	err = h.db.QueryRow(
		"SELECT id, category_id, question, answer, updated_at FROM faq_articles WHERE id = $1",
		id,
	).Scan(&a.ID, &a.CategoryID, &a.Question, &a.Answer, &a.UpdatedAt)

	if err == sql.ErrNoRows {
		RespondWithError(w, http.StatusNotFound, "Article not found")
		return
	}
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Failed to fetch article")
		return
	}

	RespondWithJSON(w, http.StatusOK, a)
}
