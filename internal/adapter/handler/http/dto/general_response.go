package dto

type GeneralResponse struct {
	Code    int    `json:"code"`
	Message string `json:"messsage"`
	Data    any    `json:"data"`
}
