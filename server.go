package microservice_users_status

import (
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

func Run(r *mux.Router) {
	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	srv.ListenAndServe()
}
