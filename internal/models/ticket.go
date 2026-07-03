package models

import "time"

type Ticket struct {
	ID          int       `json:"id"`
	StudentName string    `json:"student_name"`
	Course      int       `json:"course"`
	Question    string    `json:"question"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
}

type TicketRequest struct {
	StudentName string `json:"student_name" validate:"required"`
	Course      int    `json:"course" validate:"required,min=1,max=4"`
	Question    string `json:"question" validate:"required,min=10"`
}
