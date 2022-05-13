package quotestorage

import (
	"QuoteWebApp/common"
	quotemodel "QuoteWebApp/module/quote/model"
	"context"
)

func (s *sqlStore) FindQuote(ctx context.Context, conditions map[string]interface{})(*quotemodel.Quote, error){
	db := s.db.Table(quotemodel.Quote{}.TableName())
	var quote quotemodel.Quote
	if err := db.Where(conditions).First(&quote).Error; err != nil{
		return nil, common.ErrDB(err)
	}
	return &quote, nil
}
