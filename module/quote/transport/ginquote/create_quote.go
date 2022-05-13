package ginquote

import (
	"QuoteWebApp/component/appctx"
	quotebussiness "QuoteWebApp/module/quote/business"
	quotemodel "QuoteWebApp/module/quote/model"
	quotestorage "QuoteWebApp/module/quote/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateQuote(appContext appctx.AppContext) func(ctx *gin.Context){
	return func(c *gin.Context) {
		var newQuoteModel quotemodel.Quote
		if err := c.ShouldBind(&newQuoteModel); err != nil {
			c.JSONP(http.StatusBadRequest,gin.H{"message error":err.Error()})
			return
		}
		store := quotestorage.NewSqlStore(appContext.GetDBConnection())
		biz := quotebussiness.NewCreateQuoteBusiness(store)

		err:= biz.CreateQuote(c.Request.Context(),&newQuoteModel)
		if err != nil{
			c.JSON(http.StatusBadRequest,gin.H{"message error":err.Error()})
			return
		}
		c.JSON(http.StatusOK,gin.H{"data":newQuoteModel.ID})
	}
}