package main

import (
	"QuoteWebApp/component/appctx"
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
	db = db.Debug()
	// migrate database
	migrateDB.Migrate(db)
	appContext := appctx.NewAppContext(db)

	route := gin.Default()
	route.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})
	route.POST("/quote",ginquote.CreateQuote(appContext))
	route.GET("/quote",ginquote.GetCurrentQuoteToday(appContext))
	//route.GET("/like")
	route.Run(":8181")
}
