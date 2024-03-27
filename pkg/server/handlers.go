package server

import (
	"encoding/json"
	"github.com/ammercado/go-mysql-gke/pkg/cryptoapp"
	"log"
	"net/http"
)

func GetCoinHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("GET /api/v1/crypto")

	// get coins
	coins, err := cryptoapp.GetAllCoins(gcfg)
	if err != nil {
		log.Print("GET crypto api failed", err.Error())
		w.Write([]byte("Error getting crypto!"))
		return
	}

	// response
	json.NewEncoder(w).Encode(coins)
}

func PostCoinHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("POST /api/v1/crypto")

	// post cons
	newcoin := cryptoapp.Coin{}
	json.NewDecoder(r.Body).Decode(&newcoin)

	err := cryptoapp.PostCoin(gcfg, &newcoin)
	if err != nil {
		log.Print("POST crypto api failed", err.Error())
		w.Write([]byte("Error posting crypto !"))
		return
	}

	// response
	w.Write([]byte("New crypto published !"))
}
