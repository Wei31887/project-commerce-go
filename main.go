package main

import (
	"log"
	"net/http"
	"project/e-commerce/api"
	db "project/e-commerce/db/sqlc"
	"project/e-commerce/global"
	"project/e-commerce/initial"

	"go.uber.org/zap"
)

func main() {
	var err error
	// initialize
	path := "."
	global.CONFIG, err = initial.LoadingConfig(path)
	if err != nil {
		log.Fatalf("Failed to initialize log: %v", err)
	}

	global.LOG = initial.LoadingLogger(global.CONFIG)
	
	global.DB, err = initial.LoadingDatabase(global.CONFIG)
	if err != nil {
		global.LOG.Error("Can't open database ", zap.Error(err))
	}
	if global.DB != nil {
		defer global.DB.Close()
	}

	store := db.NewStore(global.DB)
	server := api.NewServer(global.CONFIG, store)
	runHttpServer(server)
}

func runHttpServer(server *api.Server) {
	err := server.RunServer(global.CONFIG.Server.Address)
		if err != nil && err != http.ErrServerClosed {
			global.LOG.Fatal("Listen on: ", zap.String("address", err.Error()))
		}
	global.LOG.Info("Listen on ", zap.String("address", global.CONFIG.Server.Address))
}