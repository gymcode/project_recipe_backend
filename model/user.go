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
	Activated  bool   `json:"activated"`
}

/*
	Setting default values for variables
*/

func UserModel() User{
	user := User{}
	user.Activated = false
	return user
}
