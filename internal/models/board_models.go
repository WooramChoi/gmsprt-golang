package models

import (
	"gorm.io/gorm"
)

/*
Database Entity
*/
type Board struct {
	gorm.Model
	Title     string `gorm:"size:255" json:"title"`
	Content   string `gorm:"size:4000" json:"content"`
	PlainText string `gorm:"type:text" json:"plain_text"`
	YnUse     string `gorm:"size:1;default:Y" json:"yn_use"`
	Name      string `gorm:"size:50" json:"name"`
	Pwd       string `gorm:"size:255" json:"pwd"`
}
