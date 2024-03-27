package cryptoapp

import (
	"github.com/ammercado/go-mysql-gke/pkg/config"
	"log"
)

var coinOrm = NewCoinOrm()

func GetAllCoins(cfg *config.Config) ([]*Coin, error) {
	
	coins, err := coinOrm.Get(cfg.DB)
	if err != nil {
		log.Print("error Getting crypto from DB", err.Error())
		return nil, err
	}

	return coins, nil
}

func PostCoin(cfg *config.Config, newcoin *Coin) error {
	
	err := coinOrm.Post(cfg.DB, newcoin)
	if err != nil {
		log.Print("error posting crypto to DB", err.Error())
		return err
	}

	return nil
}
