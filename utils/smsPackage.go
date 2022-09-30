package utils

import "github.com/twilio/twilio-go"

func SendSms(message string, otp uint64, recipient string) (response string, err error) {
	accountSid := "ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
	authToken := "f2xxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})

	

}
