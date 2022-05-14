package quotestorage

import (
	quotemodel "QuoteWebApp/module/quote/model"
	"context"
	"fmt"
)

func (s *sqlStore) UpdateQuoteToday(ctx context.Context,id int, updateQuoteToday *quotemodel.Quote) error{
	var data *quotemodel.Quote
	s.db.Where("id = ?").First(&data)
	fmt.Println(updateQuoteToday)
	if err := s.db.Where("id = ?",id).Updates(updateQuoteToday).Error; err != nil{
		return err
	}
	return nil
}
