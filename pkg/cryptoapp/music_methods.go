package cryptoapp

import (
	"github.com/ammercado/go-mysql-gke/pkg/config"
	"log"
)

var musicOrm = NewMusicOrm()

func GetAllSongs(cfg *config.Config) ([]*Song, error) {
	// get songs
	songs, err := musicOrm.Get(cfg.DB)
	if err != nil {
		log.Print("error Getting crypto from DB", err.Error())
		return nil, err
	}

	return songs, nil
}

func PostSong(cfg *config.Config, newsong *Song) error {
	// post song
	err := musicOrm.Post(cfg.DB, newsong)
	if err != nil {
		log.Print("error posting crypto to DB", err.Error())
		return err
	}

	return nil
}
