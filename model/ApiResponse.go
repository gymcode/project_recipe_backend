package model

type WrapSuccessResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Error   bool   `json:"error"`
	Data    User   `json:"data"`
}

type WrapFailureResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Error   bool   `json:"error"`
}
