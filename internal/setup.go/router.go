package setup

import (
	"employeejwt/internal/auth"
	"employeejwt/internal/handlers"
	"log"
	"net/http"
	"strings"
)

func SetupRouter() *http.ServeMux {

	router := http.NewServeMux()

	router.HandleFunc("/register", handlers.Register)
	router.HandleFunc("/login", handlers.Login)

	router.HandleFunc("/employees", authMiddleware(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.ListEmployees(w, r)

		case http.MethodPost:
			handlers.CreateEmployee(w, r)

		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	}))

	router.HandleFunc("/employees/", authMiddleware(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.GetEmployee(w, r)

		case http.MethodPut:
			handlers.UpdateEmployee(w, r)

		case http.MethodDelete:
			handlers.DeleteEmployee(w, r)

		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	}))

	return router
}

func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := strings.TrimSpace(r.Header.Get("Authorization"))
		parts := strings.Fields(authHeader)

		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			log.Printf("auth middleware: missing or invalid authorization header for %s %s", r.Method, r.URL.Path)
			http.Error(
				w,
				"missing or invalid authorization header",
				http.StatusUnauthorized,
			)
			return
		}

		_, err := auth.ValidateToken(parts[1])
		if err != nil {
			log.Printf("auth middleware: invalid token for %s %s: %v", r.Method, r.URL.Path, err)
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}

		next(w, r)
	}
}
