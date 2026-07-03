package handlers

import (
	"database/sql"
	"net/http"

	"college_faq/internal/models"
)

type ServiceHandler struct {
	db *sql.DB
}

func NewServiceHandler(db *sql.DB) *ServiceHandler {
	return &ServiceHandler{db: db}
}

func (h *ServiceHandler) GetServices(w http.ResponseWriter, r *http.Request) {
	rows, err := h.db.Query(`
		SELECT id, name, room, floor, phone, manager, schedule_weekdays, schedule_saturday
		FROM services
		ORDER BY id
	`)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Failed to fetch services")
		return
	}
	defer rows.Close()

	var services []models.Service
	for rows.Next() {
		var s models.Service
		if err := rows.Scan(
			&s.ID,
			&s.Name,
			&s.Room,
			&s.Floor,
			&s.Phone,
			&s.Manager,
			&s.ScheduleWeekdays,
			&s.ScheduleSaturday,
		); err != nil {
			RespondWithError(w, http.StatusInternalServerError, "Failed to parse services")
			return
		}
		services = append(services, s)
	}
	if err = rows.Err(); err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Failed to iterate services")
		return
	}

	RespondWithJSON(w, http.StatusOK, services)
}
