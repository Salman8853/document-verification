package model

type Response struct {
	Success    bool        `json:"success"`
	SuccessMsg string      `json:"success_msg"`
	Data       interface{} `json:"data,omitempty"`
}
