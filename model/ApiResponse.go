package model

type ApiResponse struct {
	Code string
	Message string
	Error bool
	Data User
}