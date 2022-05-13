package ginquote

import (
	"QuoteWebApp/component/appctx"
	quotebussiness "QuoteWebApp/module/quote/business"
	quotestorage "QuoteWebApp/module/quote/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetCurrentQuoteToday(appContext appctx.AppContext) func(ctx *gin.Context){
	return func (c *gin.Context){
		store := quotestorage.NewSqlStore(appContext.GetDBConnection())
		biz := quotebussiness.NewCreateQuoteBusiness(store)
		quoteFirst, err := biz.FindQuote(c.Request.Context(), map[string]interface{}{"is_public":true})
		if err != nil{
			c.JSON(http.StatusBadRequest,gin.H{"message error":err.Error()})
			return
		}
		c.JSON(http.StatusOK,gin.H{"data":quoteFirst})
	}
}
