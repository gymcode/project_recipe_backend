package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gymcode/project_recipe_backend/model"
	"github.com/twilio/twilio-go"
	twilioAPI "github.com/twilio/twilio-go/rest/api/v2010"
)

func SendSms(recipient string, content string) {
	config, err := LoadConfig(".")
	
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	url := "https://api.d7networks.com/messages/v1/send"
	method := "POST"

	message_globals := model.GlobalMessages{
		Originator: "Haute Cuisine OTP",
		Report_Url: "https://the_url_to_recieve_delivery_report.com",
	}

	message_object := model.MessagesObject{
		Channel:     "sms",
		Recipients:  []string{recipient},
		Content:     content,
		Msg_Type:    "text",
		Data_Coding: "text",
	}

	SmsObject := model.SmsObject{
		Messages:    []model.MessagesObject{message_object},
		Message_Glb: message_globals,
	}
	log.Println(SmsObject)
	data, _ := json.Marshal(SmsObject)
	payload := strings.NewReader(string(data))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	} 
	tokenKey := config.Sms_Token_Key
	token := fmt.Sprintf("Bearer %s", tokenKey)
	
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", token)

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
	config, err := LoadConfig(".")
	
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	accountSid := config.Sms_Account_Sid
	authToken := config.Sms_Auth_Token

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
