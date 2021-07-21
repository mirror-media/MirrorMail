package mail

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/api/gmail/v1"
)

const cred string = "configs/gmail_cred.json"

//Mailer is a general interface handles different kind of mail template structs
type Mailer interface {
	ConvertToGmail() (gmail.Message, error)
}

//Data is the struct for handling http post request
type Data struct {
	templateName string
	MailContent  map[string]interface{}
}

type MemberSubscription struct {
	MailContent map[string]interface{}
}

func (m MemberSubscription) ConvertToGmail() (gmail.Message, error) {
	var message gmail.Message

	return message, nil
}

func SendMail(c *gin.Context) {
	b, err := c.GetRawData()
	if err != nil {
		c.JSON(500, gin.H{"err": err.Error()})
		return
	}
	var mailData Data
	err = json.Unmarshal(b, &mailData)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	var m Mailer
	switch mailData.templateName {
	case "magazineSubscribe":
		m = MagazineSubscribeConfirm{mailData.MailContent}
	case "memberSubscribe":
		m = MemberSubscription{mailData.MailContent}
	case "paymentError":
		m = PaymentError{mailData.MailContent}
	default:
		c.JSON(500, gin.H{"err": errors.New("not having template required")})
		return
	}

	message, err := m.ConvertToGmail()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
	}

	GmailService, err := InitGmailClient()
	if err != nil {
		log.Printf("Failed to init Gmail client")
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	user := "me"
	_, err = GmailService.Users.Labels.List(user).Do()
	if err != nil {
		log.Printf("Unable to retrieve labels: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	//Send email here
	_, err = GmailService.Users.Messages.Send("me", &message).Do()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

}
