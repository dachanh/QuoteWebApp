package job

import (
	"QuoteWebApp/common"
	"QuoteWebApp/component/appctx"
	quotebussiness "QuoteWebApp/module/quote/business"
	quotemodel "QuoteWebApp/module/quote/model"
	quotestorage "QuoteWebApp/module/quote/storage"
	"context"
	"database/sql"
	"fmt"
	"time"
)

const FormatDate = "02/01/2006"

func UpdateQuoteToday(biz quotebussiness.QuoteStorage,updatedPosted bool, id int, now *time.Time){
	updatedQuoteToday := &quotemodel.Quote{
		PostedDate: now,
		IsPublic: sql.NullBool{
			Bool:  updatedPosted,
			Valid: true,
		},
	}
	biz.UpdateQuoteToday(context.Background(),id,updatedQuoteToday)
}

func QuoteToday(appContext appctx.AppContext){
	// the first migrate ...
	now := time.Now()
	store := quotestorage.NewSqlStore(appContext.GetDBConnection())
	biz := quotebussiness.NewCreateQuoteBusiness(store)
	quoteCurrent, err := biz.FindQuote(context.Background(), map[string]interface{}{"is_public":true})
	if  err != nil && err.Error() == common.RecordNotFound.Error(){
		UpdateQuoteToday(biz,true,1,&now)
	}else if quoteCurrent.PostedDate != nil {
		postedDateFormated := (*quoteCurrent.PostedDate).Format(FormatDate)
		nowDateFormated := now.Format(FormatDate)
		if postedDateFormated != nowDateFormated{
			UpdateQuoteToday(biz,false,quoteCurrent.ID,nil)
			newQuote , _ := biz.GetFirstUpdatedItem(context.Background(),quoteCurrent.ID)
			fmt.Println(newQuote)
			UpdateQuoteToday(biz,true,newQuote.ID,&now)
		}
	}
}
