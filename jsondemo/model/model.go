package model

type Resp struct {
	ErrNo  int    `json:"err_no"`
	Id     int    `json:"id"`
	ErrMsg string `json:"err_msg"`
	IsShow bool   `json:"is_show"`
}