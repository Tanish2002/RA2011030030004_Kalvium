package models

import "time"

type Model struct {
	DB string
}

type RequestLog struct {
	Method    string              `json:"method"`
	URL       string              `json:"url"`
	Headers   map[string][]string `json:"headers"`
	Timestamp time.Time           `json:"timestamp"`
}

// TODO
func (m *Model) AddToRequestLog(log RequestLog)
