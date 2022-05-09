package quotemodel

import "QuoteWebApp/common"

type Quote struct {
	common.SQLModel
	Content string `json:"content"gorm:"content"`
	Like int `json:"like"gorm:"like"`
	Dislike int `json:"dislike"gorm:"diske"`
}

func (q *Quote) TableName() string  {
	return "Quote"
}