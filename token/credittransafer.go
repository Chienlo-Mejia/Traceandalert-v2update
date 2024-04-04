package credittransfer

import (
	"fmt"
	"math/rand"
	"time"

	"tracealert/pkg/models/credittransfer"
	errorhandling "tracealert/pkg/models/errorHandling"
	"tracealert/pkg/utils/go-utils/database"

	"github.com/gofiber/fiber/v2"
)

var requestCounts int

const allowedRates = 5

func checkRateLimits() bool {
	requestCounts++
	if requestCounts > allowedRates {
		requestCounts = 0
		return true
	}
	return false
}
func TransferAccount(c *fiber.Ctx) error {
	// Parse request body
	trans := &credittransfer.Trans_Body{}
	if parsErr := c.BodyParser(trans); parsErr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": parsErr.Error(),
		})
	}

	// Check rate limits
	if checkRateLimits() {
		// 429
		return errorhandling.Rate_Limit_Exceeded(c, "Rate limit exceeded")
	}

	// Generate unique IDs
	instructionId := ReferenceId()
	referenceId := Iftgenerate()
	requestTrigger := time.Now().Format("2006-01-02 03:04:05")

	// Prepare response object
	response := &credittransfer.Trans_Request{
		InstructionId:   instructionId,
		ReferenceId:     referenceId,
		TransactionType: "RECEIVING", // Default to "RECEIVING"
		SenderBic:       "CAMZPHM2XXX",
		ReceivingBic:    "CBMFPHM1XXX",
		AmountCurrency:  "PHP",
		LocalInstrument: "ICRT",
		Description:     "Invalid transaction amount",
		ReasonCode:      "AC01",
		Status:          "FAILED-RJCT",
	}

	// Check transaction amount and update response accordingly
	if trans.Transactionamount > 100 {
		response.TransactionType = "SENDING"
		response.ReasonCode = "ACTC"
		response.Status = "SUCCESS"
		response.Description = "Transaction processed successfully"
	}

	// Insert transaction record into database
	insertQuery := `INSERT INTO trace_alerts.credit_transfer( instruction_id, amount_currency, description,  local_instrument, reason_code, receiving_account, receiving_bic, receiving_name, reference_id, sender_account, sender_amount, sender_bic, sender_name, status, transaction_type, date_time)
        VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	transacSave := database.DBConn.Debug().Exec(insertQuery,
		instructionId, response.AmountCurrency, response.Description, response.LocalInstrument, response.ReasonCode,
		trans.Recipientaccountnumber, response.ReceivingBic, trans.Recipientaccountname, referenceId,
		trans.Senderaccountnumber, trans.Transactionamount, response.SenderBic, trans.Senderaccountname,
		response.Status, response.TransactionType, requestTrigger).Error

	if transacSave != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}

	// Log activity for potential data breaches
	logEntry := Log{UserID: 1, Action: "access", Timestamp: time.Now()}
	MonitorForDataBreaches(logEntry)

	// Respond with response object
	return c.JSON(response)
}

// Log represents a log entry for user activities
type Log struct {
	UserID    int
	Action    string
	Timestamp time.Time
}

// MonitorForDataBreaches logs and alerts for potential data breaches
func MonitorForDataBreaches(logEntry Log) {
	// Check for potential data breaches or unauthorized access
	if logEntry.Action == "access" && logEntry.Timestamp.Hour() < 6 {
		// Detected suspicious activity, log an alert
		alertMsg := fmt.Sprintf("ALERT: Potential unauthorized access by UserID %d at %s", logEntry.UserID, logEntry.Timestamp)
		fmt.Println(alertMsg)
		// You can also send an email, log to file, etc. here
	}
}

func ReferenceId() string {
	rand.Seed(time.Now().UnixNano())
	uniqueDigit := rand.Intn(10000)
	return fmt.Sprintf("20240219CAMZPHM2XXXB0000000000%d", uniqueDigit)
}

func Iftgenerate() string {
	rand.Seed(time.Now().UnixNano())
	uniqueDigit := rand.Intn(1000)
	uniqueDigit1 := rand.Intn(10000000000000)
	return fmt.Sprintf("IFT%v-%v", uniqueDigit, uniqueDigit1)
}

func GetLocationFromIP(ip string) string {
	// Example implementation: just return a dummy location for demonstration
	return "DummyLocation"
}

// type Transaction1 struct {
// 	Amount   float64
// 	Location string
// }

// func DetectGeographicalAnomalies(transactions []Transaction1) []Transaction1 {
// 	var unusualTransactions []Transaction1
// 	locationCounts := make(map[string]int)

// 	// Count occurrences of each location
// 	for _, transaction := range transactions {
// 		locationCounts[transaction.Location]++
// 	}

// 	// Identify locations with significantly fewer transactions
// 	for location, count := range locationCounts {
// 		// You can adjust the threshold as needed
// 		if count <= 1 {
// 			for _, transaction := range transactions {
// 				if transaction.Location != location {
// 					unusualTransactions = append(unusualTransactions, transaction)
// 				}
// 			}
// 		}
// 	}

// 	return unusualTransactions
// }
