package quotemodel

import "QuoteWebApp/common"

type Quote struct {
	common.SQLModel
	Content string `json:"content"gorm:"column:content;UNIQUE;"`
	Like int `json:"like"gorm:"column:like;"`
	Dislike int `json:"dislike"gorm:"column:dislike;"`
	IsPublic *bool `json:"-" gorm:"type:boolean;column:is_public;default:false;"`
}

func (q *Quote) TableName() string  {
	return "Quote"
}