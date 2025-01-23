package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/AshFire1/service/cart"
	"github.com/AshFire1/service/order"
	"github.com/AshFire1/service/product"
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
	productStore := product.NewStore(apiConfig.db)
	productHandler := product.NewHandler(productStore)
	productHandler.RegisterRoutes(subrouter)
	orderStore := order.NewStore(apiConfig.db)

	cartHandler := cart.NewHandler(orderStore, productStore, userStore)
	cartHandler.RegisterRoutes(subrouter)
	log.Printf("API server started at %s\n", apiConfig.addr)
	return http.ListenAndServe(apiConfig.addr, router)
}
