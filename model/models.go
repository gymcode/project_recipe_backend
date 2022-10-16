package model

type SmsObject struct {
	Messages    []MessagesObject `json:"messages"`
	Message_Glb GlobalMessages   `json:"message_globals"`
}

type MessagesObject struct {
	Channel     string   `json:"channel"`
	Recipients  []string `json:"recipients"`
	Content     string   `json:"content"`
	Msg_Type    string   `json:"msg_type"`
	Data_Coding string   `json:"data_coding"`
}

type GlobalMessages struct {
	Originator string `json:"originator"`
	Report_Url string `json:"report_url"`
}

type ConfirmOtpModel struct {
	Msisdn string `json:"msisdn"`
	Code   string `json:"code"`
}
