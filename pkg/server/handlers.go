package server

import (
	"encoding/json"
	"github.com/ammercado/go-mysql-gke/pkg/cryptoapp"
	"log"
	"net/http"
)

func GetSongHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("GET /api/v1/crypto")

	// get songs
	songs, err := cryptoapp.GetAllSongs(gcfg)
	if err != nil {
		log.Print("GET crypto api failed", err.Error())
		w.Write([]byte("Error getting crypto!"))
		return
	}

	// response
	json.NewEncoder(w).Encode(songs)
}

func PostSongHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("POST /api/v1/crypto")

	// post songs
	newsong := cryptoapp.Song{}
	json.NewDecoder(r.Body).Decode(&newsong)

	err := cryptoapp.PostSong(gcfg, &newsong)
	if err != nil {
		log.Print("POST crypto api failed", err.Error())
		w.Write([]byte("Error posting crypto !"))
		return
	}

	// response
	w.Write([]byte("New crypto published !"))
}
