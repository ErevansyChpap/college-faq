package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"

	"college_faq/internal/config"
	"college_faq/internal/db"
	"college_faq/internal/handlers"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system env")
	}

	cfg := config.Load()

	dbConn, err := db.InitDB(cfg.DatabaseURL)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer dbConn.Close()

	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "http://localhost:5173"}, // добавьте свой фронт
		AllowedMethods:   []string{"GET", "POST", "PATCH", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Content-Type"},
		AllowCredentials: true,
	}))

	// Health check
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	// Инициализация обработчиков с передачей db
	categoryHandler := handlers.NewCategoryHandler(dbConn)
	faqHandler := handlers.NewFAQHandler(dbConn)
	serviceHandler := handlers.NewServiceHandler(dbConn)
	ticketHandler := handlers.NewTicketHandler(dbConn)

	// Роуты
	r.Route("/api", func(r chi.Router) {
		r.Get("/categories", categoryHandler.GetCategories)
		r.Get("/faq", faqHandler.GetFAQ)
		r.Get("/faq/{id}", faqHandler.GetFAQByID)
		r.Get("/services", serviceHandler.GetServices)
		r.Post("/tickets", ticketHandler.CreateTicket)
		r.Get("/tickets", ticketHandler.GetTickets)
		r.Patch("/tickets/{id}", ticketHandler.UpdateTicketStatus)
	})

	port := cfg.Port
	if port == "" {
		port = "8080"
	}
	log.Printf("Server starting on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
