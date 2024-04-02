package credittransfer

import (
	"fmt"
	"net"

	"github.com/gofiber/fiber/v2"
	"github.com/oschwald/geoip2-golang"
)

func Handledetails(c *fiber.Ctx) error {
	ipAddress := c.FormValue("ip")
	if ipAddress == "" {
		return c.Status(fiber.StatusBadRequest).SendString("IP address is required")
	}

	record, err := lookupIPDetails(ipAddress)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	// Return IP geolocation information as JSON
	return c.JSON(record)
}

func lookupIPDetails(ipAddress string) (*geoip2.City, error) {
	// Open the GeoIP2 database file
	db, err := geoip2.Open("GeoLite2-City.mmdb")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// Parse the IP address
	ip := net.ParseIP(ipAddress)
	if ip == nil {
		return nil, fmt.Errorf("invalid IP address: %s", ipAddress)
	}

	// Perform the IP lookup
	record, err := db.City(ip)
	if err != nil {
		return nil, err
	}

	return record, nil
}
