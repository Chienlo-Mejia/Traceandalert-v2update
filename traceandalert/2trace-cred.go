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

func Tracenetworkcred(c *fiber.Ctx) error {
	userRequest := &tracenetwork.Trans_Body{}
	var transactions []tracenetwork.Trans_Request
	userresponse := &tracenetwork.Trans_Request{}
	Uniqueidcredittransfer := Iftgeneratetracecredittransfer(32)
	// requestTrigger := time.Now().Format("03:04:05")
	// requestTrigger := time.Now().Format("2006-01-02")
	requestTrigger := time.Now().Format("2006-01-02 15:04:05")

	var Sourcetxntype string
	var Trace_alert string
	var Alerttype string
	if err := c.BodyParser(userRequest); err != nil {
		// 400 Bad Request

		loggers.Creditransferalertslogs(c.Path(), "folderName", Uniqueidcredittransfer, "(Method Not Allowed - 400)", "null", requestTrigger, "null", "null", "null", "null", "null", "null", "null", "null", 0, "null", "null", "null", "null", Trace_alert, "null", "null")
		return errorhandling.Bad_Request(c, "The request contains a bad payload")
	}

	Query := `SELECT * FROM rbi_instapay.ct_transaction WHERE 1 = 1`
	CountQuery := "SELECT COUNT(*) FROM rbi_instapay.ct_transaction WHERE 1=1"
	CountFilteredQuery := "SELECT COUNT(*) FROM rbi_instapay.ct_transaction WHERE 1=1"

	if userRequest.Since != "" {
		Query += " AND DATE_TRUNC('day', ft_date) = ?"
		CountFilteredQuery += " AND DATE_TRUNC('day', ft_date) = ?"
	}

	if userRequest.Filter != "" {
		if userRequest.Filter == "%3D%3D" {
			// Adjusted condition for maximum transaction amount of $50,000
			Query += " AND (amount > 50000 OR amount < 100 OR currency = 'USDT' ) "
			CountFilteredQuery += " AND (amount > 50000 OR amount < 100 OR currency = 'USDT' )"
			Sourcetxntype = "FRAUD"
			Trace_alert = "TRACE_FINANCIAL_CRIME"
			Alerttype = "NETWORK_ALERT"

		} else {
			Query += " AND (" + userRequest.Filter + ")"
			CountFilteredQuery += " AND (" + userRequest.Filter + ")"

		}
	}

	var filteredCount int64

	var count int64
	resultCount := database.DBConn.Raw(CountQuery).Count(&count)
	if resultCount.Error != nil {
		//400
		loggers.Creditransferalertslogs(c.Path(), "folderName", Uniqueidcredittransfer, "(Method Not Allowed - 400)", "null", requestTrigger, "null", "null", "null", "null", "null", "null", "null", "null", 0, "null", "null", "null", "null", Trace_alert, "null", "null")
		return errorhandling.Bad_Request(c, "Error counting transactions")
	}

	if userRequest.Filter != "" {
		resultFilteredCount := database.DBConn.Raw(CountFilteredQuery, userRequest.Since).Count(&filteredCount)
		if resultFilteredCount.Error != nil {
			//400
			loggers.Creditransferalertslogs(c.Path(), "folderName", Uniqueidcredittransfer, "(Method Not Allowed - 400)", "null", requestTrigger, "null", "null", "null", "null", "null", "null", "null", "null", 0, "null", "null", "null", "null", Trace_alert, "null", "null")
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
		loggers.Creditransferalertslogs(c.Path(), "folderName", Uniqueidcredittransfer, "(Method Not Allowed - 400)", "null", requestTrigger, "null", "null", "null", "null", "null", "null", "null", "null", 0, "null", "null", "null", "null", Trace_alert, "null", "null")
		return errorhandling.Bad_Request(c, "Error retrieving transactions")
	}

	if len(transactions) == 0 {
		// 404 Not Found
		loggers.Creditransferalertslogs(c.Path(), "folderName", Uniqueidcredittransfer, "(NOT_FOUND - 404)", "null", requestTrigger, "null", "null", "null", "null", "null", "null", "null", "null", 0, "null", "null", "null", "null", Trace_alert, "null", "null")
		return errorhandling.Url_Not_Found(c, "No data found for the specified date")
	}

	// Check if transaction exists and its type
	exists, err := checktxnidexist(transactions[0].Reference_id, transactions[0].Instruction_id, transactions[0].Receiving_account, transactions[0].Receiving_name, transactions[0].Sender_account, transactions[0].Sender_name)
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
		Instruction_id := transactions[i].Instruction_id
		Transaction_type := transactions[i].Transaction_type
		Status := transactions[i].Status
		Reason_code := transactions[i].Reason_code
		Local_instrument := transactions[i].Local_instrument
		Reference_id := transactions[i].Reference_id
		Sender_bic := transactions[i].Sender_bic
		Sender_name := transactions[i].Sender_name
		Sender_account := transactions[i].Sender_account
		Amount := transactions[i].Amount
		Currency := transactions[i].Currency
		Receiving_bic := transactions[i].Receiving_bic
		Receiving_name := transactions[i].Receiving_name
		Receiving_account := transactions[i].Receiving_account

		message := fmt.Sprintf(" Success %s ", Errors)
		loggers.Creditransferalertslogs(c.Path(), "folderName", Uniqueidcredittransfer, message, Instruction_id, requestTrigger, Transaction_type, Status, Reason_code, Local_instrument, Reference_id, Sender_bic, Sender_name, Sender_account, Amount, Currency, Receiving_bic, Receiving_name, Receiving_account, Trace_alert, Sourcetxntype, Alerttype)
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
