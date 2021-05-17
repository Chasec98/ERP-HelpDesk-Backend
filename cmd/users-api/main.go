package main

import (
	"net/http"

	"github.com/Chasec98/ERP-HelpDesk-Backend/pkg/health"

	"github.com/Chasec98/ERP-HelpDesk-Backend/internal/userroles"

	"github.com/Chasec98/ERP-HelpDesk-Backend/internal/rolepermissions"

	"github.com/Chasec98/ERP-HelpDesk-Backend/pkg/pagination"

	"github.com/Chasec98/ERP-HelpDesk-Backend/internal/permissions"

	"github.com/Chasec98/ERP-HelpDesk-Backend/internal/roles"

	"github.com/Chasec98/ERP-HelpDesk-Backend/pkg/logger"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"

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

	userApi := users.NewApi(users.NewService(users.NewRepository(sqlConn)))
	authRoutes.With(pagination.PaginationCtx).Get("/users", userApi.GetUsers)
	authRoutes.Get("/users/{id}", userApi.GetUser)
	authRoutes.Put("/users/{id}", userApi.PutUser)
	authRoutes.Post("/users", userApi.PostUser)

	rolesPermissionsApi := rolepermissions.NewApi(rolepermissions.NewService(rolepermissions.NewRepository(sqlConn)))
	authRoutes.Get("/roles/{roleID}/permissions", rolesPermissionsApi.GetRolesPermissions)
	authRoutes.Post("/roles/{roleID}/permissions/{permissionID}", rolesPermissionsApi.PostRolesPermissions)
	authRoutes.Delete("/roles/{roleID}/permissions/{permissionID}", rolesPermissionsApi.DeleteRolesPermissions)

	rolesApi := roles.NewApi(roles.NewService(roles.NewRepository(sqlConn)))
	authRoutes.Get("/roles", rolesApi.GetRoles)
	authRoutes.Get("/roles/{id}", rolesApi.GetRole)
	authRoutes.Post("/roles", rolesApi.PostRole)
	authRoutes.Delete("/roles/{id}", rolesApi.DeleteRole)

	permissionsApi := permissions.NewApi(permissions.NewService(permissions.NewRepository(sqlConn)))
	authRoutes.Get("/permissions", permissionsApi.GetPermissions)

	userRolesApi := userroles.NewApi(userroles.NewService(userroles.NewRepository(sqlConn)))
	authRoutes.Get("/users/{userID}/roles", userRolesApi.GetUserRoles)
	authRoutes.Post("/users/{userID}/roles/{roleID}", userRolesApi.CreateUserRole)
	authRoutes.Delete("/users/{userID}/roles/{roleID}", userRolesApi.DeleteUserRole)

	http.ListenAndServe(":3000", r)
}
