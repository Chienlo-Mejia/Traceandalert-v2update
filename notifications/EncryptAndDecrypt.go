package email

import (
	"tracealert/middleware/encryptdecrypt"

	"github.com/gofiber/fiber/v2"
)

type Sms_teles struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Encryptdata(c *fiber.Ctx) error {
	var input Sms_teles

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error parsing the request body.",
			"error":   err.Error(),
		})
	}

	secretKey := "abc&1*~#^2^#s0^=)^^7%b34"

	encryptedUser, err := encryptdecrypt.Encrypt(input.Username, secretKey)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error encrypting the user.",
			"error":   err.Error(),
		})
	}

	encryptedPassword, err := encryptdecrypt.Encrypt(input.Password, secretKey)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error encrypting the password.",
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"encryptedUser":     encryptedUser,
		"encryptedPassword": encryptedPassword,
	})
}

func Decryptdata(c *fiber.Ctx) error {
	var input Sms_teles

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error parsing the request body.",
			"error":   err.Error(),
		})
	}

	secretKey := "abc&1*~#^2^#s0^=)^^7%b34"

	decryptedUser, err := encryptdecrypt.Decrypt(input.Username, secretKey)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error decrypting the user.",
			"error":   err.Error(),
		})
	}

	decryptedPassword, err := encryptdecrypt.Decrypt(input.Password, secretKey)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error decrypting the user.",
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"decryptedUser":     decryptedUser,
		"decryptedPassword": decryptedPassword,
	})
}
