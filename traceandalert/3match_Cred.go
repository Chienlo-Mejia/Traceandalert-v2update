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

func Matchesidcredit(c *fiber.Ctx) error {
	network := &tracenetwork.Credit_transfer{}
	Uniqueidcredittransfer := IftgenerateMatchesidcredittransfer(32)
	requestTrigger := time.Now().Format("2006-01-02 15:04:05")
	var transactions []tracenetwork.Trans_Request
	var Sourcetxntype string
	var Trace_alert string
	var Alerttype string
	if err := c.BodyParser(network); err != nil {
		// Return 400 Bad Request with error message
		loggers.Creditransferalertslogs(c.Path(), "folderName", Uniqueidcredittransfer, "(Method Not Allowed - 400)", "null", requestTrigger, "null", "null", "null", "null", "null", "null", "null", "null", 0, "null", "null", "null", "null", "FEEDBACK_FINANCIAL_CRIME", "null", "null")
		return errorhandling.Bad_Request(c, "The request contains a bad payload")
	}

	if err := database.DBConn.Debug().Raw("SELECT * FROM rbi_instapay.ct_transaction WHERE  reference_id = ? OR receiving_account = ? OR receiving_name =? OR sender_account = ? OR sender_name = ?", network.Reference_id, network.Receiving_account, network.Receiving_name, network.Sender_account, network.Sender_name).Scan(&transactions).Error; err != nil {

		//500
		loggers.Creditransferalertslogs(c.Path(), "folderName", Uniqueidcredittransfer, "(INTERNAL_SERVER_ERROR - 500)", "null", requestTrigger, "null", "null", "null", "null", "null", "null", "null", "null", 0, "null", "null", "null", "null", "FEEDBACK_FINANCIAL_CRIME", "null", "null")
		return errorhandling.Internal_Server_Error(c, "Internal server error")

	}
	Query := `SELECT * FROM rbi_instapay.ct_transaction WHERE 1 = 1`
	CountFilteredQuery := "SELECT COUNT(*) FROM rbi_instapay.ct_transaction WHERE 1=1"

	if network.Filter != "" {
		if network.Filter == "%3D%3D" {
			Query += " AND (instruction_id = ? OR reference_id = ? OR sender_account = ? OR receiving_account = ?)"
			CountFilteredQuery += " AND (instruction_id = ? OR reference_id = ? OR sender_account = ? OR receiving_account = ?)"
			Sourcetxntype = "FRAUD"
			Trace_alert = "MATCH_FINANCIAL_CRIME"
			Alerttype = "TRANSACTION_ALERT"
		} else if network.Filter == "%3E%3D" {
			Query += " AND (instruction_id = '123456789'  OR reference_id = '123456789' OR sender_account = '123456789' OR receiving_account = '123456789')"
			CountFilteredQuery += " AND (instruction_id = '123456789' OR reference_id = '123456789' OR sender_account = '123456789' OR receiving_account = '123456789')"
			Sourcetxntype = "FRAUD"
			Trace_alert = "MATCH_FINANCIAL_CRIME"
			Alerttype = "TRANSACTION_ALERT"
		} else {
			Query += " AND (" + network.Filter + ")"
			CountFilteredQuery += " AND (" + network.Filter + ")"
		}
	}

	// Check if transactions are found
	if len(transactions) == 0 {
		//404
		loggers.Creditransferalertslogs(c.Path(), "folderName", Uniqueidcredittransfer, "(NOT_FOUND - 404)", "null", requestTrigger, "null", "null", "null", "null", "null", "null", "null", "null", 0, "null", "null", "null", "null", "FEEDBACK_FINANCIAL_CRIME", "null", "null")
		return errorhandling.Url_Not_Found(c, "No data found for the specified date")
	}

	// Check if transaction exists and its type
	exists, err := checktxnidexistmatchesid(transactions[0].Reference_id, transactions[0].Instruction_id, transactions[0].Receiving_account, transactions[0].Receiving_name, transactions[0].Sender_account, transactions[0].Sender_name)
	if err != nil {
		//500
		loggers.Creditransferalertslogs(c.Path(), "folderName", Uniqueidcredittransfer, "(INTERNAL_SERVER_ERROR - 500)", "null", requestTrigger, "null", "null", "null", "null", "null", "null", "null", "null", 0, "null", "null", "null", "null", "FEEDBACK_FINANCIAL_CRIME", "null", "null")
		return errorhandling.Internal_Server_Error(c, "Error while checking transaction existence")
	}
	if exists {
		Sourcetxntype = "FRAUD"
	} else {
		Sourcetxntype = "NONE"
	}

	responseBody := struct {
		Alerts            fiber.Map                    `json:"alerts"`
		TransactionAlerts []tracenetwork.Trans_Request `json:"transactionAlerts"`
	}{
		Alerts: fiber.Map{
			"InstructionsID": transactions[0].Instruction_id,
			"ReferenceID":    transactions[0].Reference_id,
			"DECISIONDATE":   requestTrigger,
			"Trace_Type":     Sourcetxntype,
			"TRACE_ALERT":    Trace_alert,
		},
		TransactionAlerts: transactions,
	}

	//------------------- logs -------------------

	var errorresp errors.Errorresp

	for i := 0; i < len(transactions); i++ {
		Errors := errorresp.Errorresponse
		transaction := transactions[i]
		Instruction_id := transaction.Instruction_id
		Reference_id := transaction.Reference_id
		Receiving_account := transaction.Receiving_account
		Receiving_name := transaction.Receiving_name
		Sender_account := transaction.Sender_account
		Sender_name := transaction.Sender_name
		Transaction_type := transaction.Transaction_type
		Status := transaction.Status
		Reason_code := transaction.Reason_code
		Local_instrument := transaction.Local_instrument
		Sender_bic := transaction.Sender_bic
		Amount := transaction.Amount
		Currency := transaction.Currency
		Receiving_bic := transaction.Receiving_bic

		message := fmt.Sprintf(" Success %s ", Errors)
		loggers.Creditransferalertslogs(c.Path(), "folderName", Uniqueidcredittransfer, message, Instruction_id, requestTrigger, Transaction_type, Status, Reason_code, Local_instrument, Reference_id, Sender_bic, Sender_name, Sender_account, Amount, Currency, Receiving_bic, Receiving_name, Receiving_account, Trace_alert, Sourcetxntype, Alerttype)
	}

	return c.JSON(responseBody)
}
func IftgenerateMatchesidcredittransfer(length int) string {
	rand.Seed(time.Now().UnixNano())

	hexDigits := "0123456789abcdef"
	result := make([]byte, length)
	for i := range result {
		result[i] = hexDigits[rand.Intn(len(hexDigits))]
	}
	return string(result)
}

func checktxnidexistmatchesid(Reference_id string, Instruction_id string, Receiving_account string, Receiving_name string, Sender_account string, Sender_name string) (bool, error) {
	var count int
	err := database.DBConn.Raw("SELECT COUNT(*) FROM rbi_instapay.ct_transaction WHERE reference_id = ? OR instruction_id = ? OR receiving_account = ? OR receiving_name = ? OR sender_account = ? OR sender_name = ?", Reference_id, Instruction_id, Receiving_account, Receiving_name, Sender_account, Sender_name).Scan(&count).Error
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
