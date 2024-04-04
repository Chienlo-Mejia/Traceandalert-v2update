package traceandalert

import (
	"fmt"
	"tracealert/pkg/models/tracenetwork"
	"tracealert/pkg/utils/go-utils/database"

	"github.com/gofiber/fiber/v2"
)

func TracetransPostmanMobilephoneBrowser(c *fiber.Ctx) error {

	var dateRange tracenetwork.Trans_Postman
	if err := c.BodyParser(&dateRange); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}
	var transactions []tracenetwork.Trans_response

	query := `SELECT * FROM trace_alerts.logstracetrans_PostmanMobilephoneBrowser WHERE 1=1`

	if dateRange.StartDate != "" {
		query += fmt.Sprintf(" AND DATE(requesttrigger) >= '%s'", dateRange.StartDate)
	}
	if dateRange.EndDate != "" {
		query += fmt.Sprintf(" AND DATE(requesttrigger) <= '%s'", dateRange.EndDate)
	}

	result := database.DBConn.Raw(query).Scan(&transactions)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Transactions not found",
		})
	}
	if len(transactions) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "No transactions found with the specified criteria",
		})
	}

	return c.JSON(transactions)
}
