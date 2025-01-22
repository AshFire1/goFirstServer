package main

import (
	"log"

	"github.com/AshFire1/cmd/api"
	"github.com/AshFire1/db"
)

func main() {
	db, err := db.NewPostgresStorage()

	server := api.NewAPISERVER(":8080", db)
	if err != nil {
		log.Fatal(err)
	}
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}

}
