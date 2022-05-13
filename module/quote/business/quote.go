package quotebussiness

import (
	quotemodel "QuoteWebApp/module/quote/model"
	"context"
)

type quoteStorage interface {
	CreateQuote(ctx context.Context,data *quotemodel.Quote ) error
	FindQuote(ctx context.Context, conditions map[string]interface{})(*quotemodel.Quote, error)
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

func (biz *quoteBusiness) FindQuote(ctx context.Context, conditions map[string]interface{})(*quotemodel.Quote, error){
	quote ,err := biz.quoteStorage.FindQuote(ctx,conditions)
	return quote, err
}
