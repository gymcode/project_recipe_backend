package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/twilio/twilio-go"
	twilioAPI "github.com/twilio/twilio-go/rest/api/v2010"
)

func SendSms() {
	url := "https://api.d7networks.com/messages/v1/send"
	method := "POST"

	payload := strings.NewReader(`{
	"messages": [
		{
			"channel": "sms",
			"recipients": ["+233268211334","+233268211334"],
			"content": "Greetings from D7 API",
			"msg_type": "text",
			"data_coding": "text"
		}
	],
	"message_globals": {
		"originator": "Haute Cuisine OTP",
		"report_url": "https://the_url_to_recieve_delivery_report.com"
	}
	}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}

func SendSmsDummy(message string, recipient string) bool {
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
