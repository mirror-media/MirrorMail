package mail

import (
	"encoding/json"

	"google.golang.org/api/gmail/v1"
)

type PaymentError struct {
	// MailContent map[string]interface{}
	MailContent string
}

func (p PaymentError) ConvertToGmail() (gmail.Message, error) {
	var message gmail.Message
	var f interface{}
	err := json.Unmarshal([]byte(p.MailContent), &f)
	if err != nil {
		return message, nil
	}
	mJson = f.(map[string]interface{})
	//To be implemented
	return message, nil
}
