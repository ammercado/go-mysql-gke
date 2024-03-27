package main

import (
	"github.com/ammercado/go-mysql-gke/pkg/config"
	"github.com/ammercado/go-mysql-gke/pkg/cryptoapp"
	"github.com/ammercado/go-mysql-gke/pkg/server"
	"log"
)

var cfg *config.Config
var err error

func init() {
	log.Print("Welcome to crypto api...")

	// get a config
	cfg, err = config.NewConfig()
	if err != nil {
		log.Fatal("Config init failed", err)
	}

	// migrate db
	if err = cryptoapp.DbInit(cfg.DB); err != nil {
		log.Fatal("DB migration failed...")
	}
}

func main() {
	server.Start(cfg)
}


