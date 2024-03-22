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

func Alerttransaction_Credittransfer(c *fiber.Ctx) error {
	userRequest := &tracenetwork.Trans_Body{}
	var transactions []tracenetwork.Trans_Request
	userresponse := &tracenetwork.Trans_Request{}
	Uniqueidcredittransfer := Iftgeneratealerttransactioncredittransfer(32)
	// requestTrigger := time.Now().Format("03:04:05")
	// requestTrigger := time.Now().Format("2006-01-02")
	requestTrigger := time.Now().Format("2006-01-02 15:04:05")

	var Sourcetxntype string
	var Trace_alert string
	var Alerttype string
	var Feedback string
	if err := c.BodyParser(userRequest); err != nil {
		// 400 Bad Request

		loggers.Creditransferalertslogs(c.Path(), "folderName", Uniqueidcredittransfer, "The request body is expecting an array", userresponse.Instruction_id, requestTrigger, "null", "null", "null", "null", "null", "null", "null", "null", 0, "null", "null", "null", "null", Trace_alert, Sourcetxntype, "null")

		return errorhandling.Bad_Request(c, "The request contains a bad payload")
	}

	Query := `SELECT * FROM trace_alerts.credit_transfer WHERE 1 = 1`
	CountQuery := "SELECT COUNT(*) FROM trace_alerts.credit_transfer WHERE 1=1"
	CountFilteredQuery := "SELECT COUNT(*) FROM trace_alerts.credit_transfer WHERE 1=1"

	if userRequest.Since != "" {
		Query += " AND DATE_TRUNC('day', datetime) = ?"
		CountFilteredQuery += " AND DATE_TRUNC('day', datetime) = ?"
	}

	if userRequest.Filter != "" {
		if userRequest.Filter == "%3E%3D" {
			Query += " AND (instructionid = '123456789'  OR referenceid = '123456789' OR senderaccount = '123456789' OR receivingaccount = '123456789')"
			CountFilteredQuery += " AND (instructionid = '123456789' OR referenceid = '123456789' OR senderaccount = '123456789' OR receivingaccount = '123456789')"
			Sourcetxntype = "FRAUD"
			Trace_alert = "ALERT_FINANCIAL_CRIME"
			Alerttype = "ACCOUNT_ALERT"
			Feedback = "ACTIONED_MULE"
		} else if userRequest.Filter == "%3D%3D" {
			Query += " AND (senderamount > 50000 OR senderamount < 100 OR amountcurrency = 'USDT' ) "
			CountFilteredQuery += " AND (senderamount > 50000 OR senderamount < 100 OR amountcurrency = 'USDT' )"
			Sourcetxntype = "FRAUD"
			Trace_alert = "TRACE_FINANCIAL_CRIME"
			Alerttype = "NETWORK_ALERT"
			Feedback = "ACTIONED_MULE"

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
		return errorhandling.Bad_Request(c, "Error counting transactions")
	}

	if userRequest.Filter != "" {
		resultFilteredCount := database.DBConn.Raw(CountFilteredQuery, userRequest.Since).Count(&filteredCount)
		if resultFilteredCount.Error != nil {
			//400
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
		return errorhandling.Bad_Request(c, "Error retrieving transactions")
	}

	if len(transactions) == 0 {
		// 404 Not Found
		return errorhandling.Url_Not_Found(c, "No data found for the specified date")
	}

	// Check if transaction exists and its type
	exists, err := checktxnidexistalerttransaction(transactions[0].Reference_id, transactions[0].Instruction_id, transactions[0].Receiving_account, transactions[0].Receiving_name, transactions[0].Sender_account, transactions[0].Sender_name)
	if err != nil {
		return errorhandling.Internal_Server_Error(c, "Error while checking transaction existence")
	}

	if exists || (userresponse.Amount > 50000 || userresponse.Amount < 100) {
		Sourcetxntype = "FRAUD"
	} else {
		Sourcetxntype = "NONE"
		Feedback = "Not Investigated"
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
			"Feedback":         Feedback,
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
		Instructionid := transactions[i].Instruction_id
		Transactiontype := transactions[i].Transaction_type
		Status := transactions[i].Status
		Reasoncode := transactions[i].Reason_code
		Localinstrument := transactions[i].Local_instrument
		Referenceid := transactions[i].Reference_id
		Senderbic := transactions[i].Sender_bic
		Sendername := transactions[i].Sender_name
		Senderaccount := transactions[i].Sender_account
		Amount := transactions[i].Amount
		Currency := transactions[i].Currency
		Receivingbic := transactions[i].Receiving_bic
		Receivingname := transactions[i].Receiving_name
		Receivingaccount := transactions[i].Receiving_account

		message := fmt.Sprintf(" Success %s ", Errors)
		loggers.Creditransferalertslogs(c.Path(), "folderName", Uniqueidcredittransfer, message, Instructionid, requestTrigger, Transactiontype, Status, Reasoncode, Localinstrument, Referenceid, Senderbic, Sendername, Senderaccount, Amount, Currency, Receivingbic, Receivingname, Receivingaccount, Trace_alert, Sourcetxntype, Alerttype)
	}

	return c.JSON(responseBody)
}

func checktxnidexistalerttransaction(Referenceid string, Instructionid string, Receivingaccount string, Receivingname string, Senderaccount string, Sendername string) (bool, error) {
	var count int
	err := database.DBConn.Raw("SELECT COUNT(*) FROM trace_alerts.credit_transfer WHERE referenceid = ? OR instructionid = ? OR receivingaccount = ? OR receivingname = ? OR senderaccount = ? OR sendername = ?", Referenceid, Instructionid, Receivingaccount, Receivingname, Senderaccount, Sendername).Scan(&count).Error
	if err != nil {
		return false, err
	}
	return count > 1, nil
}
func Iftgeneratealerttransactioncredittransfer(length int) string {
	rand.Seed(time.Now().UnixNano())

	hexDigits := "0123456789abcdef"
	result := make([]byte, length)
	for i := range result {
		result[i] = hexDigits[rand.Intn(len(hexDigits))]
	}
	return string(result)
}
