package models

import "time"

type FAQArticle struct {
	ID         int       `json:"id"`
	CategoryID int       `json:"category_id"`
	Question   string    `json:"question"`
	Answer     string    `json:"answer"`
	UpdatedAt  time.Time `json:"updated_at"`
}
