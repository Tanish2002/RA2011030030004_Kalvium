package models

import (
	"database/sql/driver"
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

type Model struct {
	DB *gorm.DB
}

type RequestLog struct {
	gorm.Model
	Method    string         `json:"method"`
	URL       string         `json:"url"`
	Headers   HeadersWrapper `json:"headers" gorm:"type:json"`
	Timestamp time.Time      `json:"timestamp"`
}

type HeadersWrapper map[string][]string

func (hw HeadersWrapper) Value() (driver.Value, error) {
	return json.Marshal(hw)
}

func (hw *HeadersWrapper) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), hw)
}

func (m *Model) AddToRequestLog(log *RequestLog) error {
	result := m.DB.Create(log)
	return result.Error
}

func (m *Model) GetRequestLog() RequestLog {
	logs := RequestLog{}
	m.DB.Order("timestamp desc").Limit(20).Find(&logs)
	return logs
}
