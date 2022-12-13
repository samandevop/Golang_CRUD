package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"crud/api"
	"crud/pkg/db"
	"crud/config"
)

func main() {

	cfg := config.Load()

	db, err := db.ConnectionDB(&cfg)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	r := gin.New()

	api.SetUpApi(r, db)

	log.Printf("Listening port %v...\n", cfg.HTTPPort)
	err = r.Run(cfg.HTTPPort)
	if err != nil {
		panic(err)
	}
}
