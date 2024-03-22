package credittransfer

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

// token := jwt.New(jwt.SigningMethodHS256)

// claims := token.Claims.(jwt.MapClaims)
// claims["identity"] = login.Username
// claims["admin"] = true
// claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

// t, err := token.SignedString([]byte("secret"))
// if err != nil {
//     return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
//         "status":  "error",
//         "message": "Internal Server Error",
//     })
// }

// cookie := fiber.Cookie{
//     Name:     "jwt",
//     Value:    t,
//     Expires:  time.Now().Add(time.Hour * 24),
//     HTTPOnly: true,
//     Secure:   true, // Set to true if using HTTPS
// }

// c.Cookie(&cookie)
type Login struct {
	Username string `json:"username"`
	// Other fields related to login information
}

func Generateandsettoken(c *fiber.Ctx) error {
	// Parse request body to get login information
	var login Login
	if err := c.BodyParser(&login); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Invalid request body",
			"message": err.Error(),
		})
	}

	// Check if the username is provided
	if login.Username == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Username is required",
			"message": "Please provide a valid username",
		})
	}

	// Generate JWT token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["identity"] = login.Username
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Sign the token
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to generate token",
		})
	}

	// Set the token as a cookie
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    t,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
		Secure:   true, // Set to true if using HTTPS
	}

	c.Cookie(&cookie)

	// Return success response
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Token generated and set successfully",
		"token":   t,
	})
}
