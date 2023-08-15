package helper

import (
	"fmt"

	"github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/utils/models"
	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/verify/v2"
)

var client *twilio.RestClient

func TwilioSetUp(accountsId, authToken string) {
	client = twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountsId,
		Password: authToken,
	})
}

func TwilioSendOTP(phoneNumber, serviceId string) (string, error) {
	params := &twilioApi.CreateVerificationParams{}
	params.SetTo(phoneNumber)
	params.SetChannel("sms")

	resp, err := client.VerifyV2.CreateVerification(serviceId, params)
	if err != nil {
		return "", err
	}

	return *resp.Sid, nil
}

func TwilioVerifyOTP(otpData models.VerifyData, serviceId string) error {
	params := &twilioApi.CreateVerificationCheckParams{}
	params.SetTo(otpData.PhoneNumber)
	params.SetCode(otpData.Code)

	// test prpose

	// 	client := &http.Client{}
	// https: //verify.twilio.com/v2/Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/AccessTokens

	// 	requestData := []byte(`{"key": "value"}`)

	// Create a new request with method, URL, and optional request body
	// req, err := http.NewRequest("POST", "https://verify.twilio.com/v2/Services/"+serviceId+"/Verifications", bytes.NewBuffer(requestData))
	// if err != nil {
	// 	fmt.Println("Error creating request:", err)
	// 	return err
	// }

	// req.Header.Add("Authorization", "Bearer YOUR_ACCESS_TOKEN")

	// // Set request headers if needed
	// req.Header.Add("Content-Type", "application/json")

	// // Send the request
	// resp, err := client.Do(req)
	// if err != nil {
	// 	fmt.Println("Error sending request:", err)
	// 	return err
	// }
	// fmt.Println("response from the server: ", resp)
	// defer resp.Body.Close()

	// until here
	resp, err := client.VerifyV2.CreateVerificationCheck(serviceId, params)

	fmt.Println(err)
	fmt.Println(*resp.Status)
	if err != nil {
		return err
	} else if *resp.Status == "approved" {
		return nil
	}

	return fmt.Errorf("cannot authenticate,reson: %s", *resp.Status)
}
