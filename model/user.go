package model

type User struct {
	ID         int    `json:"id"`
	FirstName  string `json:"firstname"`
	OtherNames string `json:"othernames"`
	Msisdn     string `json:"msisdn"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Country    string `json:"country"`
	IsoCode    string `json:"isoCode"`
	CreatedAt  string `json:"createdAt"`
}
