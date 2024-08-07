// Download the helper library from https://www.twilio.com/docs/go/install
package Message

import (
	"errors"
	"fmt"
	"os"

	"github.com/twilio/twilio-go"
	api "github.com/twilio/twilio-go/rest/api/v2010"
)

type SMS struct {
	msg string
	to  string
}

func NewSMS(to string, msg string) *SMS {
	return &SMS{to: to, msg: msg}
}

func (sms *SMS) Send(from string) error {
	sid := os.Getenv("TWILLO_ACCOUNT_SID")
	authToken := os.Getenv("TWILLO_AUTH_TOKEN")

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: sid,
		Password: authToken,
	})

	params := &api.CreateMessageParams{}
	params.SetBody(sms.msg)
	params.SetFrom(from)
	params.SetTo(sms.to)

	resp, err := client.Api.CreateMessage(params)
	if err != nil {
		fmt.Println(err)
		return errors.New("Error: unable to create text message.")
	} else {
		if resp.Body != nil {
			fmt.Println(*resp.Body)
		} else {
			fmt.Println(resp.Body)
		}
	}
	return nil
}
