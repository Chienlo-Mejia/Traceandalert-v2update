package email

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/sfreiberg/gotwilio"
)

type SmsRequest struct {
	To      string `json:"to"`
	Message string `json:"message"`
	From    string `json:"from"`
	Details string `json:"details"`
}

type SmsResponse struct {
	Response string `json:"response"`
}

func Sms(c *fiber.Ctx) error {
	accountSid := "AC15ecc8c6f2540bf4047658e01210445b"
	authToken := "9e6b55809587fe1c5eb7726a7dda7983"

	// Parse request body
	var req SmsRequest
	if err := c.BodyParser(&req); err != nil {
		return err
	}

	client := gotwilio.NewTwilioClient(accountSid, authToken)

	resp, _, err := client.SendSMS("+12695750689", req.To, req.Message, req.Details, req.From)

	if err != nil {
		fmt.Println("Error sending SMS message: " + err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(&SmsResponse{Response: "Failed to send SMS"})
	}

	responseJSON, err := json.Marshal(resp)
	if err != nil {
		fmt.Println("Error marshaling response to JSON: " + err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(&SmsResponse{Response: "Error processing response"})
	}

	fmt.Println("Response: " + string(responseJSON))

	return c.Status(fiber.StatusOK).JSON(&SmsResponse{Response: "SMS sent successfully"})
}
