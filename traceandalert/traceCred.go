package traceandalert

import (
	"fmt"
	"math/rand"
	"time"
	"tracealert/middleware/loggers"
	errorhandling "tracealert/pkg/models/errorHandling"
	"tracealert/pkg/models/errors"
	"tracealert/pkg/models/tracenetwork"
	"tracealert/pkg/utils/go-utils/database"

	"github.com/gofiber/fiber/v2"
)

// - "FEEDBACK_FINANCIAL_CRIME"
// - "MATCH_FINANCIAL_CRIME"
// - "ALERT_FINANCIAL_CRIME"
// - "TRACE_FINANCIAL_CRIME"

func TracenetworkCred(c *fiber.Ctx) error {
	userRequest := &tracenetwork.Trans_Body{}
	var transactions []tracenetwork.Trans_Request
	userresponse := &tracenetwork.Trans_Request{}
	UniqueidCredittransfer := Iftgeneratetracecredittransfer(32)
	// requestTrigger := time.Now().Format("03:04:05")
	// requestTrigger := time.Now().Format("2006-01-02")
	requestTrigger := time.Now().Format("2006-01-02 15:04:05")

	var Sourcetxntype string
	var Trace_alert string
	var Alerttype string
	if err := c.BodyParser(userRequest); err != nil {
		// 400 Bad Request

		loggers.CreditransferAlertsLogs(c.Path(), "folderName", UniqueidCredittransfer, "(Method Not Allowed - 400)", "null", requestTrigger, "null", "null", "null", "null", "null", "null", "null", "null", 0, "null", "null", "null", "null", Trace_alert, "null", "null")
		return errorhandling.Bad_Request(c, "The request contains a bad payload")
	}

	Query := `SELECT * FROM rbi_instapay.ct_transaction WHERE 1 = 1`
	CountQuery := "SELECT COUNT(*) FROM rbi_instapay.ct_transaction WHERE 1=1"
	CountFilteredQuery := "SELECT COUNT(*) FROM rbi_instapay.ct_transaction WHERE 1=1"

	if userRequest.Since != "" {
		Query += " AND DATE_TRUNC('day', ft_date) = ?"
		CountFilteredQuery += " AND DATE_TRUNC('day', ft_date) = ?"
	}

	dailyTransactionAmounts := make(map[string]float64)
	if userresponse.Amount > 0 {
		currentDate := time.Now().Format("2006-01-02")
		if amount, ok := dailyTransactionAmounts[currentDate]; ok {
			if amount+userresponse.Amount > 50000 {
				userresponse.Amount = 0
			} else {
				dailyTransactionAmounts[currentDate] += userresponse.Amount
			}
		} else {
			dailyTransactionAmounts[currentDate] = userresponse.Amount
		}
	}

	if userRequest.Filter != "" {
		if userRequest.Filter == "%3D%3D" {
			Query += " AND (amount > 50000 OR amount < 100 OR currency = 'USDT' ) "
			CountFilteredQuery += " AND (amount > 50000 OR amount < 100 OR currency = 'USDT' )"

			Sourcetxntype = "FRAUD"
			Trace_alert = "TRACE_FINANCIAL_CRIME"
			Alerttype = "NETWORK_ALERT"
		} else {
			Query += " AND ("
			Query += userRequest.Filter
			Query += " OR ("
			Query += "(hour(transaction_timestamp) < 6 OR hour(transaction_timestamp) > 23)"
			Query += " AND (" + userRequest.Filter + ")"
			Query += "))"

			CountFilteredQuery += " AND ("
			CountFilteredQuery += userRequest.Filter
			CountFilteredQuery += " OR ("
			CountFilteredQuery += "(hour(transaction_timestamp) < 6 OR hour(transaction_timestamp) > 23)"
			CountFilteredQuery += " AND (" + userRequest.Filter + ")"
			CountFilteredQuery += "))"
		}
	}

	var filteredCount int64

	var count int64
	resultCount := database.DBConn.Raw(CountQuery).Count(&count)
	if resultCount.Error != nil {
		//400
		loggers.CreditransferAlertsLogs(c.Path(), "folderName", UniqueidCredittransfer, "(Method Not Allowed - 400)", "null", requestTrigger, "null", "null", "null", "null", "null", "null", "null", "null", 0, "null", "null", "null", "null", Trace_alert, "null", "null")
		return errorhandling.Bad_Request(c, "Error counting transactions")
	}

	if userRequest.Filter != "" {
		resultFilteredCount := database.DBConn.Raw(CountFilteredQuery, userRequest.Since).Count(&filteredCount)
		if resultFilteredCount.Error != nil {
			//400
			loggers.CreditransferAlertsLogs(c.Path(), "folderName", UniqueidCredittransfer, "(Method Not Allowed - 400)", "null", requestTrigger, "null", "null", "null", "null", "null", "null", "null", "null", 0, "null", "null", "null", "null", Trace_alert, "null", "null")
			return errorhandling.Bad_Request(c, "Error counting filtered transactions")
		}
	}

	if userRequest.Limit > 0 {
		if int64(userRequest.Limit) > filteredCount {
			//400
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Limit exceeds the total number of filtered transactions",
			})
		}
		Query += fmt.Sprintf(" LIMIT %d", userRequest.Limit)
	}

	err := database.DBConn.Raw(Query, userRequest.Since).Scan(&transactions).Error

	if err != nil {
		// 400 Bad Request
		loggers.CreditransferAlertsLogs(c.Path(), "folderName", UniqueidCredittransfer, "(Method Not Allowed - 400)", "null", requestTrigger, "null", "null", "null", "null", "null", "null", "null", "null", 0, "null", "null", "null", "null", Trace_alert, "null", "null")
		return errorhandling.Bad_Request(c, "Error retrieving transactions")
	}

	if len(transactions) == 0 {
		// 404 Not Found
		loggers.CreditransferAlertsLogs(c.Path(), "folderName", UniqueidCredittransfer, "(NOT_FOUND - 404)", "null", requestTrigger, "null", "null", "null", "null", "null", "null", "null", "null", 0, "null", "null", "null", "null", Trace_alert, "null", "null")
		return errorhandling.Url_Not_Found(c, "No data found for the specified date")
	}

	// Check if transaction exists and its type
	exists, err := checktxnidexist(transactions[0].ReferenceId, transactions[0].InstructionId, transactions[0].ReceivingAccount, transactions[0].ReceivingName, transactions[0].SenderAccount, transactions[0].SenderName)
	if err != nil {
		return errorhandling.Internal_Server_Error(c, "Error while checking transaction existence")
	}

	if exists || (userresponse.Amount > 50000 || userresponse.Amount < 100) {
		Sourcetxntype = "FRAUD"
	} else {
		Sourcetxntype = "NONE"
	}

	// Prepare response body
	responseBody := struct {
		Alerts            fiber.Map                    `json:"alerts"`
		TransactionAlerts []tracenetwork.Trans_Request `json:"transactionAlerts"`
	}{
		Alerts: fiber.Map{
			"totalRecords":     count,
			"displayedRecords": userRequest.Limit,
			"DECISIONDATE":     requestTrigger,
			"Trace_Type":       Sourcetxntype,
			"TRACE_ALERT":      Trace_alert,
		},
		TransactionAlerts: transactions,
	}

	// "TRACE_ALERT": "ALERT_FINANCIAL_CRIME",
	// "Time": "08:55:39",
	// "Trace_Type": "FRAUD",
	// "displayedRecords": 3,
	// "totalRecords": 22

	//------------------- logs -------------------
	var errorresp errors.Errorresp

	userLimit := userRequest.Limit
	if userLimit < 0 {
		return errorhandling.Bad_Request(c, "Invalid limit")
	}

	for i := 0; i < userLimit && i < len(transactions); i++ {

		Errors := errorresp.Errorresponse
		InstructionId := transactions[i].InstructionId
		TransactionType := transactions[i].TransactionType
		Status := transactions[i].Status
		Reason_code := transactions[i].ReasonCode
		LocalInstrument := transactions[i].LocalInstrument
		ReferenceId := transactions[i].ReferenceId
		SenderBic := transactions[i].SenderBic
		SenderName := transactions[i].SenderName
		SenderAccount := transactions[i].SenderAccount
		Amount := transactions[i].Amount
		Currency := transactions[i].Currency
		ReceivingBic := transactions[i].ReceivingBic
		ReceivingName := transactions[i].ReceivingName
		ReceivingAccount := transactions[i].ReceivingAccount

		message := fmt.Sprintf(" Success %s ", Errors)
		loggers.CreditransferAlertsLogs(c.Path(), "folderName", UniqueidCredittransfer, message, InstructionId, requestTrigger, TransactionType, Status, Reason_code, LocalInstrument, ReferenceId, SenderBic, SenderName, SenderAccount, Amount, Currency, ReceivingBic, ReceivingName, ReceivingAccount, Trace_alert, Sourcetxntype, Alerttype)
	}

	return c.JSON(responseBody)
}

func checktxnidexist(Reference_id string, Instruction_id string, Receiving_account string, Receiving_name string, Sender_account string, Sender_name string) (bool, error) {
	var count int
	err := database.DBConn.Raw("SELECT COUNT(*) FROM rbi_instapay.ct_transaction WHERE reference_id = ? OR instruction_id = ? OR receiving_account = ? OR receiving_name = ? OR sender_account = ? OR sender_name = ?", Reference_id, Instruction_id, Receiving_account, Receiving_name, Sender_account, Sender_name).Scan(&count).Error
	if err != nil {
		return false, err
	}
	return count > 1, nil
}
func Iftgeneratetracecredittransfer(length int) string {
	rand.Seed(time.Now().UnixNano())

	hexDigits := "0123456789abcdef"
	result := make([]byte, length)
	for i := range result {
		result[i] = hexDigits[rand.Intn(len(hexDigits))]
	}
	return string(result)
}

// func isUnusualTransaction(trx tracenetwork.Trans_Request) bool {
// 	// Check for sudden large withdrawals or transfers
// 	if trx.Amount > 50000 {
// 		return true
// 	}

// 	// Parse the transaction datetime string into a time.Time object
// 	datetime, err := time.Parse(time.RFC3339, trx.Transaction_datetime)
// 	if err != nil {
// 		// Handle error, maybe log it and return false
// 		return false
// 	}

// 	// Check for transactions occurring at odd hours (between 11 PM and 6 AM)
// 	hour := datetime.Hour()
// 	if hour >= 23 || hour <= 6 {
// 		return true
// 	}

// 	// Check for transactions significantly deviating from a customer's typical behavior
// 	// For demonstration, let's assume typical behavior is transactions with amounts less than 100
// 	typicalAmount := 100.0
// 	return math.Abs(trx.Amount-typicalAmount) > 500.0
// }
