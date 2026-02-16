package main

import (
	"log"

	"github.com/abdul-burale/multi-tenant-order-management-api/internal/config"
	"github.com/abdul-burale/multi-tenant-order-management-api/internal/server"
)

func main() {

	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}
	/*
		dbConn, err := db.NewPostGres(cfg.DBConn)
		if err != nil {
			log.Fatal(err)
		}
	*/

	srv := server.New(cfg)
	server.Run(srv)
}
