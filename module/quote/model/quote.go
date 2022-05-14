package quotemodel

import (
	"QuoteWebApp/common"
	"database/sql"
	"time"
)

type Quote struct {
	common.SQLModel
	Content string `json:"content"gorm:"column:content;UNIQUE;"`
	Like int `json:"like"gorm:"column:like;"`
	Dislike int `json:"dislike"gorm:"column:dislike;"`
	PostedDate *time.Time `json:"-" gorm:"column:posted_date"`
	IsPublic sql.NullBool `json:"-" gorm:"type:boolean;column:is_public;default:false;"`
}

func (Quote) TableName() string  {
	return "Quote"
}