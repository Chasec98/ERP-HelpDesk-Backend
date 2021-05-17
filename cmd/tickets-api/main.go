package main

import (
	"net/http"

	"github.com/Chasec98/ERP-HelpDesk-Backend/pkg/health"

	"github.com/Chasec98/ERP-HelpDesk-Backend/pkg/pagination"

	"github.com/Chasec98/ERP-HelpDesk-Backend/pkg/logger"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"

	"github.com/Chasec98/ERP-HelpDesk-Backend/internal/tickets"

	customSQL "github.com/Chasec98/ERP-HelpDesk-Backend/pkg/sql"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
)

func main() {
	godotenv.Load()
	logger.InitLoggers()
	sqlConn, err := customSQL.Connect()
	if err != nil {
		logger.Error.Println(err.Error())
		panic(err)
	}
	defer sqlConn.Close()

	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "OPTIONS"},
	}))

	r.Get("/healthz", health.GetHealth)

	loggingRoutes := r.Group(nil)
	loggingRoutes.Use(middleware.Logger)

	//auth route here

	authRoutes := loggingRoutes.Group(nil)

	//auth middleware here

	ticketsApi := tickets.NewApi(tickets.NewService(tickets.NewRepository(sqlConn)))
	authRoutes.With(pagination.PaginationCtx).Get("/tickets", ticketsApi.GetTickets)
	authRoutes.Get("/tickets/{id}", ticketsApi.GetTicket)
	authRoutes.Put("/tickets/{id}", ticketsApi.PutTicket)
	authRoutes.Post("/tickets", ticketsApi.PostTicket)

	http.ListenAndServe(":3000", r)
}
