package job

import (
	"QuoteWebApp/common"
	"QuoteWebApp/component/appctx"
	quotebussiness "QuoteWebApp/module/quote/business"
	quotestorage "QuoteWebApp/module/quote/storage"
	"context"
	"fmt"
)

func QuoteToday(appContext appctx.AppContext){
	// the first migrate ...
	//now := time.Now()
	store := quotestorage.NewSqlStore(appContext.GetDBConnection())
	biz := quotebussiness.NewCreateQuoteBusiness(store)
	_, err := biz.FindQuote(context.Background(), map[string]interface{}{"is_public":true})
	fmt.Println("asdsadas")
	if  err != nil && err.Error() == common.RecordNotFound.Error(){

	}
	//if quoteCurrent.PostedDate != now{
	//
	//}
}
