package models

type LogEntry struct {
	Timestamp string `json:"timestamp"`
	Data      string `json:"data"`
}
