package quotestorage

import (
	quotemodel "QuoteWebApp/module/quote/model"
	"context"
)

func (s *sqlStore) GetFirstUpdatedItem(ctx context.Context,id int)(*quotemodel.Quote , error){
	db := s.db
	var data quotemodel.Quote
	db.Where("id != ?",id).Order("updated_at asc").First(&data)
	return &data , nil
}
