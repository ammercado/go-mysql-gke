package cryptoapp

import "gorm.io/gorm"

type Coin struct {
	gorm.Model
	Name   string `json:"name"`
	Abbreviation string `json:"artist"`
}
