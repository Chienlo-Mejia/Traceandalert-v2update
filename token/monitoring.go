package credittransfer

import (
	"encoding/json"
	"fmt"
	"time"
	"tracealert/pkg/utils/go-utils/database"

	"github.com/gofiber/fiber/v2"
)

type Trans_Request struct {
	InstructionId string  `json:"instructionId"`
	Amount        float64 `json:"amount"`
}

func MonitoringFraud(c *fiber.Ctx) error {

	err := c.SendString("Fraud tracing and health checks started")
	if err != nil {
		// Handle error
		return err
	}
	go func() {
		for {

			currentTime := time.Now()
			fmt.Println("Running fraud tracing at:", currentTime)

			var userresponse Trans_Request
			err := database.DBConn.Raw("SELECT * FROM rbi_instapay.ct_transaction").Scan(&userresponse).Error

			if err != nil {

				fmt.Println("Error querying database:", err)

				continue
			}

			responseJSON, err := json.Marshal(userresponse)
			if err != nil {

				fmt.Println("Error marshalling user response:", err)

				continue
			}

			fmt.Println("User response:", string(responseJSON))

			// Sleep for a period before querying again
			time.Sleep(time.Second)
		}
	}()

	return nil
}
