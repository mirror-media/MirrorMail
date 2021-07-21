package mail

import "google.golang.org/api/gmail/v1"

type PaymentError struct {
	MailContent map[string]interface{}
}

func (p PaymentError) ConvertToGmail() (gmail.Message, error) {
	var message gmail.Message
	//To be implemented
	return message, nil
}
