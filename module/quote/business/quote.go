package quotebussiness

import (
	quotemodel "QuoteWebApp/module/quote/model"
	"context"
)

type quoteStorage interface {
	CreateQuote(ctx context.Context,data *quotemodel.Quote ) error
}

type quoteBusiness struct {
	quoteStorage quoteStorage
}

func NewCreateQuoteBusiness(store quoteStorage) *quoteBusiness{
	return &quoteBusiness{
		quoteStorage: store,
	}
}

func (biz *quoteBusiness) CreateQuote(ctx context.Context,data *quotemodel.Quote) error{
	err := biz.quoteStorage.CreateQuote(ctx,data)
	return err
}
