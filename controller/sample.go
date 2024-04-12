package controller

import (
	"tracealert/pkg/models/tracenetwork"
	"tracealert/pkg/utils/go-utils/database"

	"github.com/gofiber/fiber/v2"
)

func ShowPage1(c *fiber.Ctx) error {

	return c.Render("login", fiber.Map{
		"titlePage": "SAMPLE PAGE",
	})

}
func Location(c *fiber.Ctx) error {

	return c.Render("locationDetect", fiber.Map{
		"titlePage": "SAMPLE PAGE",
	})

}
func ShowPage(c *fiber.Ctx) error {
	ctResult := &[]tracenetwork.Errorlog{}

	database.DBConn.Raw("SELECT * FROM trace_alerts.errorlogs WHERE 1=1 ").Scan(ctResult)

	return c.JSON(ctResult)
}
