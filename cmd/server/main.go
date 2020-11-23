package main

import (
	"net/http"

	"github.com/Chasec98/ERP-HelpDesk-Backend/pkg/logger"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"

	"github.com/Chasec98/ERP-HelpDesk-Backend/internal/auth"
	"github.com/Chasec98/ERP-HelpDesk-Backend/internal/tickets"

	"github.com/Chasec98/ERP-HelpDesk-Backend/internal/health"

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
		panic(err)
	}
	defer sqlConn.Close()

	r := chi.NewRouter()

	healthRouter := health.HealthRouter()

	r.Mount("/healthz", healthRouter)

	loggingRoutes := r.Group(nil)
	loggingRoutes.Use(middleware.Logger)

	loggingRoutes.Mount("/auth", auth.AuthRouter(auth.NewService(auth.NewRepository(sqlConn))))

	authRoutes := loggingRoutes.Group(nil)
	//r.Use(customMiddleware.Auth)
	authRoutes.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "OPTIONS"},
	}))

	authRoutes.Mount("/tickets", tickets.TicketRouter(tickets.NewService(tickets.NewRepository(sqlConn))))
	authRoutes.Mount("/users", users.UserRouter(users.NewService(users.NewRepository(sqlConn))))

	http.ListenAndServe(":3000", r)
}

//TODO: healthcheck w/ no middleware
