package mail

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"

	"google.golang.org/api/gmail/v1"
)

type MagazineSubscribeConfirm struct {
	// MailContent map[string]interface{}
	MailContent string
}

func (m MagazineSubscribeConfirm) ConvertToGmail() (gmail.Message, error) {
	var message gmail.Message

	emailBody, err := m.parseTemplate()
	if err != nil {
		return message, errors.New("unable to parse email template")
	}

	emailTo := fmt.Sprintf("To: %s \r\n", mJson["BuyerEmail"])

	subject := fmt.Sprintf("Subject: Your Subscription in MirrorMedia: %s \n",
		mJson["MerchantOrderNo"])

	mime := "MIME-version: 1.0;\nContent-Type: text/plain; charset=\"UTF-8\";\n\n"
	msg := []byte(emailTo + subject + mime + "\n" + emailBody)
	message.Raw = base64.URLEncoding.EncodeToString(msg)

	return message, nil
}

func (m MagazineSubscribeConfirm) parseTemplate() (string, error) {
	var f interface{}
	err := json.Unmarshal([]byte(m.MailContent), &f)
	if err != nil {
		return "", err
	}
	mJson = f.(map[string]interface{})

	tt, err := template.ParseFiles("templates/magazine_email.tmpl") // Which template to use
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	if err = tt.Execute(buf, mJson); err != nil {
		return "", err
	}
	body := buf.String()
	return body, nil
}
