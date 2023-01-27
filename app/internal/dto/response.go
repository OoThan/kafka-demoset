package dto

type ResponseObject struct {
	ErrCode int64       `json:"err_code"`
	ErrMsg  string      `json:"err_msg"`
	Data    interface{} `json:"data,omitempty"`
}
