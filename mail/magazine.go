package mail

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"google.golang.org/api/gmail/v1"
	"html/template"
)

type MagazineSubscribeConfirm struct {
	MailContent map[string]interface{}
}

func (m MagazineSubscribeConfirm) ConvertToGmail() (gmail.Message, error) {
	var message gmail.Message

	emailBody, err := m.parseTemplate()
	if err != nil {
		return message, errors.New("unable to parse email template")
	}

	emailTo := fmt.Sprintf("To: %s \r\n", m.MailContent["BuyerEmail"])

	subject := fmt.Sprintf("Subject: Your Subscription in MirrorMedia: %s \n",
		m.MailContent["MerchantOrderNo"])

	mime := "MIME-version: 1.0;\nContent-Type: text/plain; charset=\"UTF-8\";\n\n"
	msg := []byte(emailTo + subject + mime + "\n" + emailBody)
	message.Raw = base64.URLEncoding.EncodeToString(msg)

	return message, nil
}

func (m MagazineSubscribeConfirm) parseTemplate() (string, error) {
	tt, err := template.ParseFiles("templates/email.tmpl") // Which template to use
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	if err = tt.Execute(buf, m.MailContent); err != nil {
		return "", err
	}
	body := buf.String()
	return body, nil
}
