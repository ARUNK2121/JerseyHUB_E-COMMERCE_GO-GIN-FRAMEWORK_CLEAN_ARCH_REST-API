package helper

import (
	"errors"
	"fmt"

	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/verify/v2"
)

var client *twilio.RestClient

func TwilioSetup(username string, password string) {
	fmt.Println(username, password)
	client = twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: username,
		Password: password,
	})

}

func TwilioSendOTP(phone string, serviceID string) (string, error) {
	fmt.Println("anybody here?")
	to := "+91" + phone
	params := &twilioApi.CreateVerificationParams{}
	params.SetTo(to)
	params.SetChannel("sms")
	fmt.Println(*params.To, *params.Channel)
	fmt.Println(serviceID)

	resp, err := client.VerifyV2.CreateVerification(serviceID, params)
	if err != nil {

		fmt.Println("CHECK CHECK")
		fmt.Println(err)
		return " ", err
	}

	return *resp.Sid, nil

}

func TwilioVerifyOTP(serviceID string, code string, phone string) error {

	params := &twilioApi.CreateVerificationCheckParams{}
	params.SetTo("+91" + phone)
	params.SetCode(code)
	resp, err := client.VerifyV2.CreateVerificationCheck(serviceID, params)

	if err != nil {
		return err
	}

	if *resp.Status == "approved" {
		return nil
	}

	return errors.New("failed to validate otp")

}
