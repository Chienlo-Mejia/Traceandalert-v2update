package email

import (
	"fmt"
	"net"

	"github.com/oschwald/geoip2-golang"
)

func getClientLocation(ipAddress string) (string, error) {
	// Open the GeoIP2 database
	db, err := geoip2.Open("GeoLite2-City.mmdb")
	if err != nil {
		return "", err
	}
	defer db.Close()

	// Parse the IP address
	ip := net.ParseIP(ipAddress)
	if ip == nil {
		return "", fmt.Errorf("invalid IP address: %s", ipAddress)
	}

	// Retrieve location data for the IP address
	record, err := db.City(ip)
	if err != nil {
		return "", err
	}

	// Construct location string
	location := fmt.Sprintf("%s, %s", record.City.Names["en"], record.Country.Names["en"])
	return location, nil
}
