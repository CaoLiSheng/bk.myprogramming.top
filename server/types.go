package srv

import "bk.myprogramming.top/db"

type Result struct {
	Code    int         `json:"code"`
	Results interface{} `json:"results,omitempty"`
	Message string		`json:"message,omitempty"`
	Err     error		`json:"-"`
}

type Job func(*db.Core) *Result

type JobOptions struct {
	Simple bool
	Auto   bool
}
