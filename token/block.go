package credittransfer

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

var blockedAccounts = make(map[string]bool)

func BlockAccount(Senderaccountnumber string) error {
	if Senderaccountnumber == "" {
		return fmt.Errorf("account number cannot be empty")
	}
	blockedAccounts[Senderaccountnumber] = true
	return nil
}

func UnblockAccount(Senderaccountnumber string) error {
	if Senderaccountnumber == "" {
		return fmt.Errorf("account number cannot be empty")
	}
	if _, exists := blockedAccounts[Senderaccountnumber]; !exists {
		return fmt.Errorf("account %s is not blocked", Senderaccountnumber)
	}
	delete(blockedAccounts, Senderaccountnumber)
	return nil
}

func IsAccountBlocked(accountNumber string) bool {
	return blockedAccounts[accountNumber]
}

func BlockSenderAccount(c *fiber.Ctx) error {
	var blockReq struct {
		Senderaccountnumber string `json:"senderaccountnumber"`
	}

	if err := c.BodyParser(&blockReq); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "failed to parse request body",
		})
	}

	// Validate the request
	if blockReq.Senderaccountnumber == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Account number is required",
		})
	}

	if err := BlockAccount(blockReq.Senderaccountnumber); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": fmt.Sprintf("Account %s has been blocked successfully", blockReq.Senderaccountnumber),
	})
}

func UnblockSenderAccount(c *fiber.Ctx) error {
	var unblockReq struct {
		Senderaccountnumber string `json:"senderaccountnumber"`
	}

	if err := c.BodyParser(&unblockReq); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "failed to parse request body",
		})
	}

	// Validate the request
	if unblockReq.Senderaccountnumber == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Account number is required",
		})
	}

	if err := UnblockAccount(unblockReq.Senderaccountnumber); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": fmt.Sprintf("Account %s has been unblocked successfully", unblockReq.Senderaccountnumber),
	})
}
