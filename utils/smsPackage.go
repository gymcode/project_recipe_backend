package utils

import (
	"github.com/twilio/twilio-go"
	twilioAPI "github.com/twilio/twilio-go/rest/api/v2010"
)

func SendSms(message string, otp uint64, recipient string) (res twilioAPI.ApiV2010Message, err error){
	accountSid := "ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
	authToken := "f2xxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})

	params := &twilioAPI.CreateMessageParams{}
	params.SetTo(recipient)
	params.SetFrom("Haute Cuisine")
	params.SetBody(message)

	resp, err := client.Api.CreateMessage(params)
	
	res = *resp
	return res, err

}
