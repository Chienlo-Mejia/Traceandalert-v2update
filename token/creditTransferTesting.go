package credittransfer

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"tracealert/pkg/models/creditTransfer"
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
	go cleanupBlockedIDs()
	return false
}

var blockedIDs map[string]bool = make(map[string]bool)
var accountLocks sync.Mutex

func BlockID(c *fiber.Ctx) error {
	var blockRequest struct {
		Senderaccountnumber string `json:"senderaccountnumber"`
	}

	if err := c.BodyParser(&blockRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	accountLocks.Lock()
	defer accountLocks.Unlock()

	blockedIDs[blockRequest.Senderaccountnumber] = true

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"Senderaccountnumber": fmt.Sprintf("%s blocked for transactions", blockRequest.Senderaccountnumber),
	})
}
func UnblockID(c *fiber.Ctx) error {
	var unblockRequest struct {
		Senderaccountnumber string `json:"senderaccountnumber"`
	}

	if err := c.BodyParser(&unblockRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "failed to parse request body",
		})
	}

	// Validate the request
	if unblockRequest.Senderaccountnumber == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Account number is required",
		})
	}

	accountLocks.Lock()
	defer accountLocks.Unlock()

	if _, exists := blockedIDs[unblockRequest.Senderaccountnumber]; !exists {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fmt.Sprintf("Account %s is not blocked", unblockRequest.Senderaccountnumber),
		})
	}

	delete(blockedIDs, unblockRequest.Senderaccountnumber)

	return c.JSON(fiber.Map{
		"message": fmt.Sprintf("Account %s has been unblocked successfully", unblockRequest.Senderaccountnumber),
	})
}

func TransferAccount(c *fiber.Ctx) error {
	trans := &creditTransfer.TransBody{}
	if parsErr := c.BodyParser(trans); parsErr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": parsErr.Error(),
		})
	}
	accountLocks.Lock()
	defer accountLocks.Unlock()

	if _, ok := blockedIDs[trans.Senderaccountnumber]; ok {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"ALERT": fmt.Sprintf("This Account is Maintanance for Senderaccountnumber: %s", trans.Senderaccountnumber),
		})
	}

	if checkRateLimits() {
		// 429
		return errorhandling.Rate_Limit_Exceeded(c, "Rate limit exceeded")
	}

	instructionId := ReferenceId()
	referenceId := Iftgenerate()
	requestTrigger := time.Now().Format("2006-01-02 03:04:05")

	response := &creditTransfer.TransRequest{
		InstructionId:   instructionId,
		ReferenceId:     referenceId,
		TransactionType: "RECEIVING",
		SenderBic:       "CAMZPHM2XXX",
		ReceivingBic:    "CBMFPHM1XXX",
		AmountCurrency:  "PHP",
		LocalInstrument: "ICRT",
		Description:     "Invalid transaction amount",
		ReasonCode:      "AC01",
		Status:          "FAILED-RJCT",
	}

	if trans.Transactionamount > 100 {
		response.TransactionType = "SENDING"
		response.ReasonCode = "ACTC"
		response.Status = "SUCCESS"
		response.Description = "Transaction processed successfully"
	}

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

	logEntry := Log{UserID: 1, Action: "access", Timestamp: time.Now()}
	MonitorForDataBreaches(logEntry)

	return c.JSON(response)
}

type Log struct {
	UserID    int
	Action    string
	Timestamp time.Time
}

func MonitorForDataBreaches(logEntry Log) {
	if logEntry.Action == "access" && logEntry.Timestamp.Hour() < 6 {
		alertMsg := fmt.Sprintf("ALERT: Potential unauthorized access by UserID %d at %s", logEntry.UserID, logEntry.Timestamp)
		fmt.Println(alertMsg)

	}
}
func cleanupBlockedIDs() {
	for {
		accountLocks.Lock()
		for senderAccountNumber := range blockedIDs {
			_ = senderAccountNumber
		}
		accountLocks.Unlock()
		time.Sleep(time.Minute)
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
	return "DummyLocation"
}

// type BlockTransactionRequest struct {
// 	SenderAccount string `json:"senderAccount"`
// }

// var (
// 	BlockedTransactions = struct {
// 		sync.RWMutex
// 		transactions map[string]bool
// 	}{
// 		transactions: make(map[string]bool),
// 	}
// )

// func BlockTransaction(c *fiber.Ctx) error {
// 	defer func() {
// 		if r := recover(); r != nil {
// 			// Log the panic
// 			log.Println("Panic occurred in BlockTransaction:", r)
// 			// Optionally, you can log the stack trace as well
// 			debug.PrintStack()
// 			// Respond with a 500 Internal Server Error
// 			c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 				"error": "Internal Server Error",
// 			})
// 		}
// 	}()

// 	var req BlockTransactionRequest
// 	if err := c.BodyParser(&req); err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"error": err.Error(),
// 		})
// 	}

// 	BlockedTransactions.Lock()
// 	BlockedTransactions.transactions[req.SenderAccount] = true
// 	BlockedTransactions.Unlock()

// 	return c.JSON(fiber.Map{
// 		"message": "Transaction blocked successfully",
// 	})
// }

// func UnblockTransaction(c *fiber.Ctx) error {
// 	var req BlockTransactionRequest
// 	if err := c.BodyParser(&req); err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"error": err.Error(),
// 		})
// 	}

// 	BlockedTransactions.Lock()
// 	delete(BlockedTransactions.transactions, req.SenderAccount)
// 	BlockedTransactions.Unlock()

// 	return c.JSON(fiber.Map{
// 		"message": "Transaction unblocked successfully",
// 	})
// }
// func TransferAccount(c *fiber.Ctx) error {
// 	// Parse request body
// 	var req BlockTransactionRequest
// 	if err := c.BodyParser(&req); err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"error": err.Error(),
// 		})
// 	}

// 	BlockedTransactions.RLock()
// 	defer BlockedTransactions.RUnlock()
// 	if BlockedTransactions.transactions[req.SenderAccount] {
// 		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
// 			"error": "Sender account is blocked",
// 		})
// 	}

// 	trans := &creditTransfer.Trans_Body{}
// 	if parsErr := c.BodyParser(trans); parsErr != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"error": parsErr.Error(),
// 		})
// 	}

// 	// Generate unique IDs
// 	instructionId := ReferenceId()
// 	referenceId := Iftgenerate()
// 	requestTrigger := time.Now().Format("2006-01-02 03:04:05")

// 	// Prepare response object
// 	response := &creditTransfer.Trans_Request{
// 		InstructionId:   instructionId,
// 		ReferenceId:     referenceId,
// 		TransactionType: "RECEIVING", // Default to "RECEIVING"
// 		SenderBic:       "CAMZPHM2XXX",
// 		ReceivingBic:    "CBMFPHM1XXX",
// 		AmountCurrency:  "PHP",
// 		LocalInstrument: "ICRT",
// 		Description:     "Invalid transaction amount",
// 		ReasonCode:      "AC01",
// 		Status:          "FAILED-RJCT",
// 	}

// 	// Check transaction amount and update response accordingly
// 	if trans.Transactionamount > 100 {
// 		response.TransactionType = "SENDING"
// 		response.ReasonCode = "ACTC"
// 		response.Status = "SUCCESS"
// 		response.Description = "Transaction processed successfully"
// 	}

// 	// Insert transaction record into database only if the sender's account is not blocked
// 	if !BlockedTransactions.transactions[req.SenderAccount] {
// 		insertQuery := `INSERT INTO trace_alerts.credit_transfer( instruction_id, amount_currency, description,  local_instrument, reason_code, receiving_account, receiving_bic, receiving_name, reference_id, sender_account, sender_amount, sender_bic, sender_name, status, transaction_type, date_time)
//         VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

// 		transacSave := database.DBConn.Exec(insertQuery,
// 			instructionId, response.AmountCurrency, response.Description, response.LocalInstrument, response.ReasonCode,
// 			trans.Recipientaccountnumber, response.ReceivingBic, trans.Recipientaccountname, referenceId,
// 			trans.SenderAccount, trans.Transactionamount, response.SenderBic, trans.Senderaccountname,
// 			response.Status, response.TransactionType, requestTrigger).Error

// 		if transacSave != nil {
// 			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
// 		}
// 	}

// 	// Log activity for potential data breaches
// 	logEntry := Log{UserID: 1, Action: "access", Timestamp: time.Now()}
// 	MonitorForDataBreaches(logEntry)

// 	// Respond with response object
// 	return c.JSON(response)
// }

// //--------------------------------------

// // Log represents a log entry for user activities
// type Log struct {
// 	UserID    int
// 	Action    string
// 	Timestamp time.Time
// }

// // MonitorForDataBreaches logs and alerts for potential data breaches
// func MonitorForDataBreaches(logEntry Log) {
// 	// Check for potential data breaches or unauthorized access
// 	if logEntry.Action == "access" && logEntry.Timestamp.Hour() < 6 {
// 		// Detected suspicious activity, log an alert
// 		alertMsg := fmt.Sprintf("ALERT: Potential unauthorized access by UserID %d at %s", logEntry.UserID, logEntry.Timestamp)
// 		fmt.Println(alertMsg)
// 		// You can also send an email, log to file, etc. here
// 	}
// }

// func ReferenceId() string {
// 	rand.Seed(time.Now().UnixNano())
// 	uniqueDigit := rand.Intn(10000)
// 	return fmt.Sprintf("20240219CAMZPHM2XXXB0000000000%d", uniqueDigit)
// }

// func Iftgenerate() string {
// 	rand.Seed(time.Now().UnixNano())
// 	uniqueDigit := rand.Intn(1000)
// 	uniqueDigit1 := rand.Intn(10000000000000)
// 	return fmt.Sprintf("IFT%v-%v", uniqueDigit, uniqueDigit1)
// }

//--------------------------------
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
