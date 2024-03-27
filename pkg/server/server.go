package server

import (
	"github.com/ammercado/go-mysql-gke/pkg/config"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var gcfg *config.Config // global config for server

func Start(cfg *config.Config) {
	gcfg = cfg
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/crypto", GetCoinHandler).Methods("GET")
	r.HandleFunc("/api/v1/crypto", PostCoinHandler).Methods("POST")

	log.Fatal(http.ListenAndServe(":"+cfg.ServerPort, r))
}
