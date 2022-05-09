package quotestorage

import (
	quotemodel "QuoteWebApp/module/quote/model"
	"context"
)

func (s *sqlStore) CreateQuote(ctx context.Context, data *quotemodel.Quote) error{
	db := s.db.Begin()
	if err := db.Table(data.TableName()).Create(data).Error; err != nil{
		db.Rollback()

	}
	if err := db.Commit().Error; err != nil{
		db.Rollback()

	}
	return nil
}