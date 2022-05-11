package migrateDB

import (
	quotemodel "QuoteWebApp/module/quote/model"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB){
	db.AutoMigrate(quotemodel.Quote{})
}