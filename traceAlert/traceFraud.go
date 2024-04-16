package traceandalert

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"time"
	"tracealert/pkg/models/tracenetwork"

	"github.com/gofiber/fiber/v2"
)

type ReqFraud struct {
	SenderAccount string `json:"senderAccount"`
}

func TraceFraud(c *fiber.Ctx) error {
	req := &ReqFraud{}
	iftGenUniqueId := IftGenUniqueId(32)
	if err := c.BodyParser(req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	connection := fmt.Sprintf("http://192.168.0.110:8000/api/public/v1/%s", req.SenderAccount)
	resp, err := http.Get(connection)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch funds amount", "details": err.Error()})
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to read response body", "details": err.Error()})
	}

	var response tracenetwork.TransRequest
	var Feedback string
	if err := json.Unmarshal(body, &response.Amount); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to parse funds amount", "details": err.Error()})
	}

	if response.Amount <= 100 || response.Amount >= 50000 {
		Feedback = "FRAUD"
		return c.Status(http.StatusOK).JSON(fiber.Map{"message": "fraud detected"})
	} else {
		Feedback = "CLEAN"
	}

	reqBody, err := json.Marshal(req)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Internal server error"})
	}

	resp, err = http.Post(connection, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to send request", "details": err.Error()})
	}
	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to read response body", "details": err.Error()})
	}

	if resp.StatusCode != http.StatusOK {
		return c.Status(resp.StatusCode).JSON(fiber.Map{"error": "Request failed with status code " + strconv.Itoa(resp.StatusCode), "details": string(body)})
	}

	if err := json.Unmarshal(body, &response); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to parse response JSON", "details": err.Error()})
	}

	respData := struct {
		Message        string                    `json:"message"`
		Header         http.Header               `json:"header"`
		Data           tracenetwork.TransRequest `json:"data"`
		IftGenUniqueId string                    `json:"IftGenUniqueId"`
		Time           time.Time                 `json:"Time"`
		Feedback       string                    `json:"Feedback"`
	}{
		Message:        "success",
		Header:         resp.Header,
		Data:           response,
		IftGenUniqueId: iftGenUniqueId,
		Time:           time.Now(),
		Feedback:       Feedback,
	}

	return c.Status(fiber.StatusOK).JSON(respData)
}

func IftGenUniqueId(length int) string {
	rand.Seed(time.Now().UnixNano())

	hexDigits := "0123456789abcdef"
	result := make([]byte, length)
	for i := range result {
		result[i] = hexDigits[rand.Intn(len(hexDigits))]
	}
	return string(result)
}

// func InquiryCustomer_Account(c *fiber.Ctx) error {
// 	// Parse the request body into the inqCustomer variable

// 	inqCustomer := &igateModel.IustomerInfo{}
// 	if parsErr := c.BodyParser(inqCustomer); parsErr != nil {
// 		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
// 			"error": parsErr.Error(),
// 		})
// 	}

// 	// Construct the igate URL
// 	queryParams := url.Values{}
// 	queryParams.Set("data", inqCustomer.ReferenceNumber)
// 	queryParams.Set("filterType", inqCustomer.FilterType)
// 	queryParams.Set("accountNo", inqCustomer.AccountNo)

// 	// Construct the igate URL with the query parameters
// 	igateURL := fmt.Sprintf("%s/openapi/v1/inqCustByAccount?data=%s", envRouting.IGateBaseUrl, queryParams.Encode())

// 	fmt.Println("IGATE URL:", igateURL)

// 	// Create an HTTP GET request
// 	req, err := http.NewRequest(http.MethodGet, igateURL, nil)
// 	if err != nil {
// 		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
// 			"error": "Failed to create the HTTP request",
// 		})
// 	}

// 	fmt.Println("Req:", req)
// 	// Set request headers
// 	req.Header.Add("Content-Type", "application/json")
// 	req.Header.Add("Merchant-ID", "QVBJMDAwMDU=")

// 	// Perform the HTTP request
// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
// 			"error": "Failed to make the HTTP request",
// 		})
// 	}
// 	fmt.Println("Resp:", resp)
// 	defer resp.Body.Close()

// 	// Read the response body
// 	body, readErr := ioutil.ReadAll(resp.Body)
// 	if readErr != nil {
// 		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
// 			"error": "Failed to read the response body",
// 		})
// 	}

// 	// Check the HTTP status code
// 	if resp.StatusCode != http.StatusOK {
// 		return c.Status(resp.StatusCode).JSON(fiber.Map{
// 			"error": "Request failed with status code " + strconv.Itoa(resp.StatusCode),
// 		})
// 	}

// 	// Successful response
// 	response := struct {
// 		Message string      `json:"message"`
// 		Header  http.Header `json:"header"`
// 		Data    string      `json:"data"`
// 	}{
// 		Message: "success",
// 		Header:  resp.Header,
// 		Data:    string(body),
// 	}

// 	return c.JSON(response)
// }
