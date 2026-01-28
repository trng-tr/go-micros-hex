package dtos

import (
	"time"
)

type Response struct {
	Status    string `json:"status"`
	Message   string `json:"message"`
	CreatedAt string `json:"created_date"`
}

func NewResponse(status, message string) *Response {
	return &Response{
		Status:    status,
		Message:   message,
		CreatedAt: time.Now().Format("2006-01-02 15:04-05"),
	}
}
