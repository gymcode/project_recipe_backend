package utils

import (
	"log"

	"github.com/twilio/twilio-go"
	twilioAPI "github.com/twilio/twilio-go/rest/api/v2010"
)

func SendSms(message string, recipient string) bool{
	accountSid := ""
	authToken := ""

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})

	params := &twilioAPI.CreateMessageParams{}
	params.SetTo(recipient)
	params.SetFrom("+17087752229")
	params.SetBody(message)

	resp, err := client.Api.CreateMessage(params)
	log.Println(resp)
	log.Println(err)

	if err != nil {
		return false 
	} else {
		return true
	}

}
