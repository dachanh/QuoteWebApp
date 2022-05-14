package main

import (
	"QuoteWebApp/component/appctx"
	"QuoteWebApp/component/appworker"
	"QuoteWebApp/component/appworker/job"
	"QuoteWebApp/migrateDB"
	"QuoteWebApp/module/quote/transport/ginquote"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

func main() {
	dsn := os.Getenv("DNS")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	workers := appworker.AppWorkers{}
	db = db.Debug()
	// migrate database
	migrateDB.Migrate(db)
	appContext := appctx.NewAppContext(db)

	route := gin.Default()
	route.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})
	route.POST("/quote", ginquote.CreateQuote(appContext))
	route.GET("/quote", ginquote.GetCurrentQuoteToday(appContext))
	//route.GET("/like")
	updateQuoteTodayJob := workers.SetUpWorker()
	updateQuoteTodayJob.SetDelayTime(5)
	updateQuoteTodayJob.SetIdleTime(60*10)
	updateQuoteTodayJob.SetTask(func() {
		job.QuoteToday(appContext)
	})
	workers.Run()
	route.Run(":8181")
}
