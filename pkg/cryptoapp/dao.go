package cryptoapp

import (
	"gorm.io/gorm"
	"log"
)

const (
	CreateBatchSizeDefault = 100
)

type musicDaoInterface interface {
	Get(db *gorm.DB) ([]*Coin, error)
	Post(db *gorm.DB, coin *Coin) error
}

type coinApiOrm struct {
}

func NewCoinOrm() musicDaoInterface {
	return &coinApiOrm{}
}

func (m *coinApiOrm) Get(db *gorm.DB) ([]*Coin, error) {
	coinsArr := make([]*Coin, 0)
	if err := db.Model(&Coin{}).Find(&coinsArr).Error; err != nil {
		return nil, err
	}
	log.Printf("%d rows found", len(coinsArr))

	return coinsArr, nil
}

func (m *coinApiOrm) Post(db *gorm.DB, coin *Coin) error {
	if err := db.Session(&gorm.Session{FullSaveAssociations: true, CreateBatchSize: CreateBatchSizeDefault}).Model(Coin{}).Create(coin).Error; err != nil {
		return err
	}

	return nil
}
