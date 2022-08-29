package model

type WrapSuccessResponse struct {
	Code string
	Message string
	Error bool
	Data User
}

type WrapFailureResponse struct {
	Code string
	Message string
	Error bool
}