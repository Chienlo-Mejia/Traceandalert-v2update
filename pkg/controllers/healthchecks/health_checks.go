package healthchecks

import (
	"fmt"
	"tracealert/pkg/models/errors"
	"tracealert/pkg/models/response"
	"tracealert/pkg/utils/go-utils/encryptDecrypt"

	"github.com/gofiber/fiber/v2"
)

func checkHealth() response.ResponseModel {
	return response.ResponseModel{
		RetCode: "100",
		Message: "Request success!",
		Data: errors.ErrorModel{
			Message:   "Service is available!",
			IsSuccess: true,
			Error:     nil,
		},
	}
}

func checkHealthB() response.ResponseModel {
	return response.ResponseModel{
		RetCode: "100",
		Message: "Request success B!",
		Data: errors.ErrorModel{
			Message:   "Service is available!",
			IsSuccess: true,
			Error:     nil,
		},
	}
}

func CheckServiceHealth(c *fiber.Ctx) error {
	health := checkHealth()
	response := errors.ErrorModel{}
	response = health.Data.(errors.ErrorModel)
	if !response.IsSuccess {
		return c.JSON(health)
	}
	plaintext := "Postgres"
	secretKey := "abc&1*~#^2^#s0^=)^^7%b34"

	ciphertext, err := encryptDecrypt.Encrypt(plaintext, secretKey)
	if err != nil {
		fmt.Println("Error:", err)

	}

	fmt.Println("Encrypted:", ciphertext)

	// To decrypt the ciphertext
	decryptedText, err := encryptDecrypt.Decrypt(ciphertext, secretKey)
	if err != nil {
		fmt.Println("Error:", err)

	}

	fmt.Println("Decrypted:", decryptedText)

	return c.JSON(health)
}

func CheckServiceHealthB(c *fiber.Ctx) error {
	health := checkHealthB()
	response := errors.ErrorModel{}
	response = health.Data.(errors.ErrorModel)
	if !response.IsSuccess {
		return c.JSON(health)
	}
	return c.JSON(health)
}

// const (
// 	secretKey = "abc&1*~#^2^#s0^=)^^7%b34"
// )

// func Token(c *fiber.Ctx) error {
// 	plaintext := "Postgres"

// 	ciphertext, err := encryptDecrypt.Encrypt(plaintext, secretKey)
// 	if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"error": fmt.Sprintf("encryption failed: %v", err),
// 		})
// 	}

// 	fmt.Println("Encrypted:", ciphertext)

// 	// To decrypt the ciphertext
// 	decryptedText, err := encryptDecrypt.Decrypt(ciphertext, secretKey)
// 	if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"error": fmt.Sprintf("decryption failed: %v", err),
// 		})
// 	}

// 	fmt.Println("Decrypted:", decryptedText)

// 	return c.SendString(decryptedText)
// }
