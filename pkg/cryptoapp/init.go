package cryptoapp

import "gorm.io/gorm"

func DbInit(db *gorm.DB) error {
	// Migrate the schema
	if err := db.AutoMigrate(&Coin{}); err != nil {
		return err
	}
	return nil
}
