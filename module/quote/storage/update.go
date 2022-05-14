package quotestorage

import (
	quotemodel "QuoteWebApp/module/quote/model"
	"context"
)

func (s *sqlStore) UpdateQuoteToday(ctx context.Context,id int, updateQuoteToday *quotemodel.Quote) error{
	var data *quotemodel.Quote
	s.db.Where("id = ?").First(&data)
	if err := s.db.Where("id = ?").Updates(updateQuoteToday).Error; err != nil{
		return err
	}
	return nil
}
