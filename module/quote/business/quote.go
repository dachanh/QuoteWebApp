package quotebussiness

import (
	"QuoteWebApp/common"
	quotemodel "QuoteWebApp/module/quote/model"
	"context"
	"errors"
	"fmt"
	"sync"
)

type QuoteStorage interface {
	CreateQuote(ctx context.Context,data *quotemodel.Quote ) error
	FindQuote(ctx context.Context, conditions map[string]interface{})(*quotemodel.Quote, error)
	UpdateQuoteToday(ctx context.Context,id int, updateQuoteToday *quotemodel.Quote) error
	GetFirstUpdatedItem(ctx context.Context,id int)(*quotemodel.Quote , error)
}

type quoteBusiness struct {
	quoteStorage QuoteStorage
}

func NewCreateQuoteBusiness(store QuoteStorage) *quoteBusiness{
	return &quoteBusiness{
		quoteStorage: store,
	}
}

func (biz *quoteBusiness) CreateQuote(ctx context.Context,data *quotemodel.Quote) error{
	_, err:= biz.FindQuote(ctx,map[string]interface{}{"content":data.Content})
	if err == nil{
		return common.ErrDB(errors.New("The content is existed!"))
	}
	err = biz.quoteStorage.CreateQuote(ctx,data)
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

func(biz *quoteBusiness) GetFirstUpdatedItem(ctx context.Context,id int)(*quotemodel.Quote , error){
	quote, err := biz.quoteStorage.GetFirstUpdatedItem(ctx,id)
	return quote, err
}

func(biz *quoteBusiness) UpdateLikeOrDislike(ctx context.Context, isLike bool) error{
	quoteCurrent, err := biz.FindQuote(ctx, map[string]interface{}{"is_public":true})
	if err != nil{
		return common.ErrDB(err)
	}
	var mu sync.Mutex
	mu.Lock()
	if isLike{
		quoteCurrent.Like += 1

	}else if quoteCurrent.Like != 0{
		quoteCurrent.Like -= 1
	}else {
		return nil
	}
	err = biz.quoteStorage.UpdateQuoteToday(ctx,quoteCurrent.ID,quoteCurrent)
	mu.Unlock()
	if err != nil{
		return common.ErrDB(err)
	}
	return nil
}