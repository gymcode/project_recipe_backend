package model

type OTP struct {
	Id        int    `json:"id"`
	Msisdn    string `json:"msisdn"`
	HashedOtp string `json:"hashedOtp"`
	ExpireAt  string `json:"expireAt"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updastedAt"`
}
