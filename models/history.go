package models

import (
	"time"

	"gorm.io/gorm"
)

type Model struct {
	DB *gorm.DB
}

type RequestLog struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Question  string    `json:"question"`
	Answer    float64   `json:"answer"`
	Timestamp time.Time `json:"timestamp"`
}

func (m *Model) AddToRequestLog(log *RequestLog) error {
	result := m.DB.Create(log)
	return result.Error
}

func (m *Model) GetRequestLogs() ([]RequestLog, error) {
	logs := []RequestLog{}
	result := m.DB.Order("timestamp desc").Limit(20).Find(&logs)
	return logs, result.Error
}
