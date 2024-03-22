package email

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// type Sms_tele struct {
// 	Code         int    `json:"code"`
// 	Name         string `json:"name"`
// 	Timestamp    string `json:"timestamp"`
// 	MessageCount int    `json:"msgcount"`
// 	TelcoID      int    `json:"telco_id"`
// 	MessageID    string `json:"messageId"`
// }

// func Sms_telerivet(c *fiber.Ctx) error {
// 	var input Sms_tele

// 	if err := c.BodyParser(&input); err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"message": "Error parsing the request body.",
// 			"error":   err.Error(),
// 		})
// 	}

// 	url := "https://dev-mercury.fortress-asya.com:17000/api/public/v1/message/broadcast"
// 	jsonBody, err := json.Marshal(input)
// 	if err != nil {
// 		return err
// 	}

// 	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
// 	if err != nil {
// 		return err
// 	}

// 	req.Header.Set("Content-Type", "application/json")

// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		return err
// 	}

// 	defer resp.Body.Close()

// 	if resp.StatusCode != http.StatusOK {
// 		return errors.New("unexpected response status code")
// 	}

//		return nil
//	}
type MessagePayload struct {
	Msg     string `json:"msg"`
	From    string `json:"from"`
	AppCode string `json:"appCode,omitempty"`
	To      string `json:"to"`
}

type MessageResponse struct {
	ID            int             `json:"id,omitempty"`
	MessageID     string          `json:"message_id"`
	Response      json.RawMessage `json:"response"`
	TimeInserted  string          `json:"time_inserted"`
	ApplicationID int             `json:"application_id"`
	UserID        int             `json:"user_id"`
}

// func Sms_telerivet(c *fiber.Ctx) error {
// 	// Parse JSON payload from request
// 	payload := &MessagePayload{}
// 	if err := c.BodyParser(&payload); err != nil {
// 		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
// 	}

// 	// Convert payload to JSON
// 	payloadJSON, err := json.Marshal(payload)
// 	if err != nil {
// 		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Internal server error"})
// 	}

// 	Connection := "http://192.168.0.110:8000/api/public/v1/message/broadcast"

// 	// Create HTTP request
// 	req, err := http.NewRequest("POST", Connection, bytes.NewBuffer(payloadJSON))
// 	if err != nil {
// 		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Internal server error"})
// 	}

// 	req.Header.Set("Content-Type", "application/json")

// 	// Create HTTP client
// 	client := &http.Client{}
// 	resp, err := client.Do(req)

// 	if err != nil {
// 		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to send SMS", "details": err.Error()})
// 	}
// 	defer resp.Body.Close()

// 	// Read response body
// 	body, readErr := ioutil.ReadAll(resp.Body)
// 	if readErr != nil {
// 		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
// 			"error": readErr.Error(),
// 		})
// 	}

// 	// Check the HTTP status code
// 	if resp.StatusCode != http.StatusOK {
// 		return c.Status(resp.StatusCode).JSON(fiber.Map{
// 			"error":   "Request failed with status code " + strconv.Itoa(resp.StatusCode),
// 			"details": string(body),
// 		})
// 	}

// 	var Response MessageResponse
// 	if unmarshalErr := json.Unmarshal(body, &Response); unmarshalErr != nil {
// 		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
// 			"error": unmarshalErr.Error(),
// 		})
// 	}
// 	response := struct {
// 		Message string          `json:"message"`
// 		Header  http.Header     `json:"header"`
// 		Data    MessageResponse `json:"data"`
// 	}{
// 		Message: "success",
// 		Header:  resp.Header,
// 		Data:    Response,
// 	}
// 	return c.Status(fiber.StatusOK).JSON(response)
// }

func Sms_telerivet(c *fiber.Ctx) error {
	payload := &MessagePayload{}
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}
	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Internal server error"})
	}

	// Connection := "http://192.168.0.113:8000/api/public/v1/message/broadcast"
	Connection := "https://dev-mercury.fortress-asya.com:17000/api/v1/send/message"

	// Create HTTP request
	req, err := http.NewRequest(http.MethodPost, Connection, bytes.NewBuffer(payloadJSON))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create the HTTP request",
		})
	}

	req.Header.Set("Content-Type", "application/json")

	// Create HTTP client
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to send SMS", "details": err.Error()})
	}
	defer resp.Body.Close()

	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": readErr.Error(),
		})
	}

	if resp.StatusCode != http.StatusOK {

		return c.Status(resp.StatusCode).JSON(fiber.Map{
			"error":   "Request failed with status code " + strconv.Itoa(resp.StatusCode),
			"details": string(body),
		})
	}

	var Response MessageResponse
	if unmarshalErr := json.Unmarshal(body, &Response); unmarshalErr != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": unmarshalErr.Error(),
		})
	}
	response := struct {
		Message string          ` json:"message"`
		Header  http.Header     ` json:"header"`
		Data    MessageResponse `json:"data"`
	}{
		Message: "success",
		Header:  resp.Header,
		Data:    Response,
	}
	return c.Status(fiber.StatusOK).JSON(response)
}
