package userRoutes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vishalsharma/api/controller/users"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set common response headers
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Methods", "GET")
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Methods", "PUT")

		// Call the next handler in the chain
		next.ServeHTTP(w, r)
	})
}

func UserRoutes() *mux.Router {
	r := mux.NewRouter()
	r.Use(Middleware)
	r.HandleFunc("/users", users.CreateUser).Methods("POST")
	r.HandleFunc("/users/login", users.Login).Methods("POST")
	return r
}
