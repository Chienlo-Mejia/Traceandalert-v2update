package encryptDecrypt

import (
	"github.com/gofiber/fiber/v2"
)

//	func Test(c *fiber.Ctx) error {
//		test := "localhost"
//		test1 := "qwqeqw"
//		output := Encrypt(test, test1)(error.Error())
//		return c.JSON(output)
//	}
func Test(c *fiber.Ctx) error {
	test := "localhost"
	test1 := "qwqeqw"
	output, err := Encrypt(test, test1)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"output": output,
	})
}
