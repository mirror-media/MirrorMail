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

var mJson map[string]interface{}

type Hello struct {
	// MailContent map[string]interface{}
	MailContent string
}

func (h Hello) parseTemplate() (string, error) {
	tt, err := template.ParseFiles("templates/hello_mail.tmpl") // Which template to use
	if err != nil {
		return "", err
	}

	var f interface{}
	err = json.Unmarshal([]byte(h.MailContent), &f)
	mJson = f.(map[string]interface{})

	buf := new(bytes.Buffer)
	if err = tt.Execute(buf, mJson); err != nil {
		return "", err
	}
	body := buf.String()
	return body, nil
}

func (h Hello) ConvertToGmail() (gmail.Message, error) {
	var message gmail.Message
	emailBody, err := h.parseTemplate()

	if err != nil {
		return message, errors.New("unable to parse email template")
	}

	emailTo := fmt.Sprintf("To: %s \r\n", mJson["To"])

	subject := fmt.Sprintf("Subject: Hello from : %s \n",
		mJson["SendName"])

	mime := "MIME-version: 1.0;\nContent-Type: text/plain; charset=\"UTF-8\";\n\n"
	msg := []byte(emailTo + subject + mime + "\n" + emailBody)

	message.Raw = base64.URLEncoding.EncodeToString(msg)

	return message, nil
}
