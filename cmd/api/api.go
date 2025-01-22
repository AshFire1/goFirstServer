package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/AshFire1/service/user"
	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPISERVER(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}
func (apiConfig *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()
	userStore := user.NewStore(apiConfig.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subrouter)
	log.Printf("API server started at %s\n", apiConfig.addr)
	return http.ListenAndServe(apiConfig.addr, router)
}
