package quotebussiness

import (
	"QuoteWebApp/common"
	quotemodel "QuoteWebApp/module/quote/model"
	"context"
	"fmt"
)

type quoteStorage interface {
	CreateQuote(ctx context.Context,data *quotemodel.Quote ) error
	FindQuote(ctx context.Context, conditions map[string]interface{})(*quotemodel.Quote, error)
	UpdateQuoteToday(ctx context.Context,id int, updateQuoteToday *quotemodel.Quote) error
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
	if err != nil{
		return  common.ErrDB(err)
	}
	return err
}

func (biz *quoteBusiness) FindQuote(ctx context.Context, conditions map[string]interface{})(*quotemodel.Quote, error){
	quote ,err := biz.quoteStorage.FindQuote(ctx,conditions)
	if err != nil{
		return nil, common.ErrDB(err)
	}
	return quote, err
}

func (biz *quoteBusiness) UpdateQuoteToday(ctx context.Context,id int, updateQuoteToday *quotemodel.Quote) error {
	quote ,err := biz.quoteStorage.FindQuote(ctx, map[string]interface{}{"id":id})
	if err != nil{
		return common.ErrDB(err)
	}
	fmt.Println(quote)
	err = biz.quoteStorage.UpdateQuoteToday(ctx,id,updateQuoteToday)
	if err != nil{
		return common.ErrDB(err)
	}
	return nil
}