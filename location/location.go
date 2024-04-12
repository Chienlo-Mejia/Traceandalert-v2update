package location

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type IPStackResponse struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func GetLocationLocation(ip string) (float64, float64, error) {
	apiKey := "AIzaSyB9BOMZFGLdjFzokzd7reQlSn2y9y4EOCM" // Replace with your actual API key
	url := fmt.Sprintf("http://api.ipstack.com/%s?access_key=%s", ip, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		return 0, 0, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, 0, err
	}

	var ipstackResp IPStackResponse
	err = json.Unmarshal(body, &ipstackResp)
	if err != nil {
		return 0, 0, err
	}

	return ipstackResp.Latitude, ipstackResp.Longitude, nil
}

func GetLocation(c *fiber.Ctx) error {
	ipAddress := c.IP()

	latitude, longitude, err := GetLocationLocation(ipAddress)
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}

	location := Location{
		Latitude:  latitude,
		Longitude: longitude,
	}

	return c.JSON(location)
}
