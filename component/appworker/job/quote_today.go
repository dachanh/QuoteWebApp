package job

import (
	"QuoteWebApp/common"
	"QuoteWebApp/component/appctx"
	quotebussiness "QuoteWebApp/module/quote/business"
	quotemodel "QuoteWebApp/module/quote/model"
	quotestorage "QuoteWebApp/module/quote/storage"
	"context"
	"database/sql"
	"time"
)

const FormatDate = "02/01/2006"
func QuoteToday(appContext appctx.AppContext){
	// the first migrate ...
	now := time.Now()
	store := quotestorage.NewSqlStore(appContext.GetDBConnection())
	biz := quotebussiness.NewCreateQuoteBusiness(store)
	quoteCurrent, err := biz.FindQuote(context.Background(), map[string]interface{}{"is_public":true})
	if  err != nil && err.Error() == common.RecordNotFound.Error(){
		updatedQuoteToday := &quotemodel.Quote{
			PostedDate: &now,
			IsPublic: sql.NullBool{
				Bool:  true,
				Valid: true,
			},
		}
		_ = biz.UpdateQuoteToday(context.Background(),1,updatedQuoteToday)
	}else if quoteCurrent.PostedDate != nil {
		postedDateFormated := (*quoteCurrent.PostedDate).Format(FormatDate)
		nowDateFormated := now.Format(FormatDate)
		if postedDateFormated != nowDateFormated{

		}
	}
}
