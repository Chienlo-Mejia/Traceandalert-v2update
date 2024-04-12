package traceandalert

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
	"tracealert/middleware/loggers"
	errorHandling "tracealert/pkg/models/errorHandling"
	"tracealert/pkg/models/errors"
	"tracealert/pkg/models/tracenetwork"
	"tracealert/pkg/utils/go-utils/database"

	"github.com/gofiber/fiber/v2"
)

var (
	mu                 sync.Mutex
	requestCountsPerID = make(map[string]map[string]int)
)

const maxTransactionsPerDay = 5

func CheckratelimitsFeedback(ID string) bool {
	mu.Lock()
	defer mu.Unlock()
	today := time.Now().Format("2006-01-02")
	if _, ok := requestCountsPerID[today]; !ok {
		requestCountsPerID[today] = make(map[string]int)
	}
	requestCountsPerID[today][ID]++

	return requestCountsPerID[today][ID] > maxTransactionsPerDay
}

func ResetRequestCountsDaily() {
	for {
		// Get the current time
		now := time.Now()

		// Calculate the duration until the next midnight
		nextMidnight := now.Truncate(24 * time.Hour).Add(24 * time.Hour)
		timeToMidnight := nextMidnight.Sub(now)
		// Sleep until the next midnight
		time.Sleep(timeToMidnight)
		// Reset counts for the next day
		mu.Lock()
		delete(requestCountsPerID, now.Format("2006-01-02"))
		mu.Unlock()
	}
}

// - "FEEDBACK_FINANCIAL_CRIME"
// - "MATCH_FINANCIAL_CRIME"
// - "ALERT_FINANCIAL_CRIME"
// - "TRACE_FINANCIAL_CRIME"

func AlertnetworkCredit(c *fiber.Ctx) error {
	network := &tracenetwork.Credit_transfer{}
	UniqueidCredittransfer := iftGenerateAlert_CreditTransfer(32)
	requestTrigger := time.Now().Format("2006-01-02 15:04:05")

	var transactions []tracenetwork.Trans_Request
	var SourceTxnType string
	var TraceAlert string
	var AlertType string

	if err := c.BodyParser(network); err != nil {
		// Return 400 Bad Request with error message
		loggers.CreditransferAlertsLogs(c.Path(), "folderName", UniqueidCredittransfer, "(Method Not Allowed - 400)", "null", requestTrigger, "null", "null", "null", "null", "null", "null", "null", "null", 0, "null", "null", "null", "null", "FEEDBACK_FINANCIAL_CRIME", "null", "null")
		return errorHandling.Bad_Request(c, "The request contains a bad payload")
	}

	// Check rate limits
	if CheckratelimitsFeedback(network.InstructionId) {
		// 429
		loggers.CreditransferAlertsLogs(c.Path(), "folderName", UniqueidCredittransfer, "(Method Not Allowed - 429)", "null", requestTrigger, "null", "null", "null", "null", "null", "null", "null", "null", 0, "null", "null", "null", "null", "FEEDBACK_FINANCIAL_CRIME", "null", "null")
		return errorHandling.Rate_Limit_Exceeded(c, "Rate limit exceeded")
	}

	if err := database.DBConn.Debug().Raw("SELECT * FROM rbi_instapay.ct_transaction WHERE instruction_id = ? ", network.InstructionId).Scan(&transactions).Error; err != nil {
		// 500
		loggers.CreditransferAlertsLogs(c.Path(), "folderName", UniqueidCredittransfer, "(INTERNAL_SERVER_ERROR - 500)", "null", requestTrigger, "null", "null", "null", "null", "null", "null", "null", "null", 0, "null", "null", "null", "null", "FEEDBACK_FINANCIAL_CRIME", "null", "null")
		return errorHandling.Internal_Server_Error(c, "Internal server error")

	}
	Query := `SELECT * FROM rbi_instapay.ct_transaction WHERE 1 = 1`
	CountFilteredQuery := "SELECT COUNT(*) FROM rbi_instapay.ct_transaction WHERE 1=1"

	if network.Filter != "" {
		if network.Filter == "%3D%3D" {
			Query += " AND (instruction_id = ? OR referenceid = ? OR sender_account = ? OR receiving_account = ?)"
			CountFilteredQuery += " AND (instruction_id = ? OR reference_id = ? OR sender_account = ? OR receiving_account = ?)"
			SourceTxnType = "FRAUD"
			TraceAlert = "ALERT_FINANCIAL_CRIME"
			AlertType = "ACCOUNT_ALERT"

		} else if network.Filter == "%3E%3D" {
			Query += " AND (instruction_id = '123456789'  OR reference_id = '123456789' OR sender_account = '123456789' OR receiving_account = '123456789')"
			CountFilteredQuery += " AND (instruction_id = '123456789' OR reference_id = '123456789' OR sender_account = '123456789' OR receiving_account = '123456789')"
			SourceTxnType = "FRAUD"
			TraceAlert = "ALERT_FINANCIAL_CRIME"
			AlertType = "ACCOUNT_ALERT"

		} else {
			Query += " AND (" + network.Filter + ")"
			CountFilteredQuery += " AND (" + network.Filter + ")"

		}
	}

	// Check if transactions are found
	if len(transactions) == 0 {
		//404
		loggers.CreditransferAlertsLogs(c.Path(), "folderName", UniqueidCredittransfer, "(NOT_FOUND - 404)", "null", requestTrigger, "null", "null", "null", "null", "null", "null", "null", "null", 0, "null", "null", "null", "null", "FEEDBACK_FINANCIAL_CRIME", "null", "null")
		return errorHandling.Url_Not_Found(c, "No data found for the specified date")
	}

	// Check if transaction exists and its type
	exists, err := checkTxnId_ExistAlert(transactions[0].InstructionId)
	if err != nil {
		//500
		loggers.CreditransferAlertsLogs(c.Path(), "folderName", UniqueidCredittransfer, "(INTERNAL_SERVER_ERROR - 500)", "null", requestTrigger, "null", "null", "null", "null", "null", "null", "null", "null", 0, "null", "null", "null", "null", "FEEDBACK_FINANCIAL_CRIME", "null", "null")
		return errorHandling.Internal_Server_Error(c, "Error while checking transaction existence")
	}
	if exists {
		SourceTxnType = "FRAUD"
	} else {
		SourceTxnType = "NONE"
	}

	// Prepare response body
	responseBody := struct {
		Alerts            fiber.Map                    `json:"alerts"`
		TransactionAlerts []tracenetwork.Trans_Request `json:"transactionAlerts"`
	}{
		Alerts: fiber.Map{
			"InstructionsID": transactions[0].InstructionId,
			"ReferenceID":    transactions[0].ReferenceId,
			"DECISIONDATE":   requestTrigger,
			"Trace_Type":     SourceTxnType,
			"TRACE_ALERT":    TraceAlert,
			"ALERT_TYPE":     AlertType,
		},
		TransactionAlerts: transactions,
	}

	//------------------- logs -------------------

	var errorresp errors.Errorresp

	for i := 0; i < len(transactions); i++ {
		// Extract the transaction at index i
		Errors := errorresp.Errorresponse
		transaction := transactions[i]
		InstructionId := transaction.InstructionId
		ReferenceId := transaction.ReferenceId
		ReceivingAccount := transaction.ReceivingAccount
		ReceivingName := transaction.ReceivingName
		SenderAccount := transaction.SenderAccount
		SenderName := transaction.SenderName
		TransactionType := transaction.TransactionType
		Status := transaction.Status
		ReasonCode := transaction.ReasonCode
		LocalInstrument := transaction.LocalInstrument
		SenderBic := transaction.SenderBic
		Amount := transaction.Amount
		Currency := transaction.Currency
		ReceivingBic := transaction.ReceivingBic

		message := fmt.Sprintf(" Success %s ", Errors)
		loggers.CreditransferAlertsLogs(c.Path(), "folderName", UniqueidCredittransfer, message, InstructionId, requestTrigger, TransactionType, Status, ReasonCode, LocalInstrument, ReferenceId, SenderBic, SenderName, SenderAccount, Amount, Currency, ReceivingBic, ReceivingName, ReceivingAccount, TraceAlert, SourceTxnType, AlertType)
	}

	return c.JSON(responseBody)
}
func iftGenerateAlert_CreditTransfer(length int) string {
	rand.Seed(time.Now().UnixNano())

	hexDigits := "0123456789abcdef"
	result := make([]byte, length)
	for i := range result {
		result[i] = hexDigits[rand.Intn(len(hexDigits))]
	}
	return string(result)
}

func checkTxnId_ExistAlert(Instruction_id string) (bool, error) {
	var count int
	err := database.DBConn.Raw("SELECT COUNT(*) FROM rbi_instapay.ct_transaction WHERE  instruction_id = ? ", Instruction_id).Scan(&count).Error
	if err != nil {
		return false, err
	}
	return count > 1, nil
}

//200
//404
//400
//405
//429

//401
//403
