package dtos

import "time"

type Response struct {
	Status    Status `json:"status"`
	Message   string `json:"message"`
	CreatedAt string `jsnon:"created_at"`
}
type Status string

const (
	Success Status = "SUCCESS"
	Fail    Status = "FAILED"
)

func NewResponse(s Status, m string) *Response {
	return &Response{
		Status:    s,
		Message:   m,
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}
}
