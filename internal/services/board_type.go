package services

import (
	"time"
)

/*
Request Models
TODO Service 와의 인터페이스 과정에 struct 를 사용할지 고민(항상 고민하는 것 같다)
*/
type BoardSearch struct {
	Page     int // TODO common type
	PageSize int // TODO common type
	Name     string
	Toc      string
	YnUse    string
}

type BoardAdd struct {
	Title     string `json:"title"`
	Content   string `json:"content"`
	PlainText string `json:"plain_text"`
	Name      string `json:"name"`
	Pwd       string `json:"pwd"`
}

type BoardModify struct {
	Title     string `json:"title"`
	Content   string `json:"content"`
	PlainText string `json:"plain_text"`
	YnUse     string `json:"yn_use"`
	Name      string `json:"name"`
	Pwd       string `json:"pwd"`
	NewPwd    string `json:"new_pwd"`
}

/*
Response Models
*/
type BoardCommon struct {
	// gorm.Model
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`

	Title string `json:"title"`
	YnUse string `json:"yn_use"`
	Name  string `json:"name"`
}

type BoardSummary struct {
	BoardCommon
	ContentSummary string `json:"content_summary"`
}

type BoardDetails struct {
	BoardCommon
	Content   string `json:"content"`
	PlainText string `json:"plain_text"`
}
