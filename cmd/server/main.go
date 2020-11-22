package main

import (
	"net/http"

	"github.com/Chasec98/ERP-HelpDesk-Backend/pkg/logger"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"

	"github.com/Chasec98/ERP-HelpDesk-Backend/internal/tickets"

	"github.com/Chasec98/ERP-HelpDesk-Backend/internal/health"

	"github.com/Chasec98/ERP-HelpDesk-Backend/internal/auth"
	"github.com/Chasec98/ERP-HelpDesk-Backend/internal/users"

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
	}
	defer sqlConn.Close()

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	//r.Use(customMiddleware.Auth)
	//r.Use(customMiddleware.JSON)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "OPTIONS"},
	}))

	ticketRouter := tickets.TicketRouter(tickets.NewService(tickets.NewRepository(sqlConn)))

	r.Mount("/tickets", ticketRouter)
	r.Route("/users", users.UserRoutes)
	r.Route("/auth", auth.AuthRoutes)
	r.Route("/health", health.HealthRoutes)

	http.ListenAndServe(":3000", r)
}

//TODO: healthcheck w/ no middleware
