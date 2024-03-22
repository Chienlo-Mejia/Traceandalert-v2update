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
// func Feedbackcredit(c *fiber.Ctx) error {
// 	network := &tracenetwork.Credittransfer_feedback{}
// 	requestTrigger := time.Now().Format("2006-01-02 15:04:05")
// 	if err := c.BodyParser(network); err != nil {
// 		// Return 400 Bad Request with error message
// 		return errorhandling.Bad_Request(c, "The request contains a bad payload")
// 	}

// 	var transactions []tracenetwork.Credittransfer_feedback_response

// 	if err := database.DBConn.Debug().Raw("SELECT * FROM trace_alerts.logscredittransfer WHERE  uniqueidcredittransfer = ? AND alerttype = ?  AND feedback = ? AND sourcetxntype = ? AND trace_alert = ?", network.Uniqueidcredittransfer, network.Alerttype, network.Feedback, network.Tracetype, network.Trace_alert).Scan(&transactions).Error; err != nil {
// 		return errorhandling.Internal_Server_Error(c, "Internal server error")
// 	}

// 	if len(transactions) == 0 {
// 		return errorhandling.Bad_Request(c, "No transactions found for the given sender account and transaction reference")
// 	}

// 	responseBody := struct {
// 		TraceAlerts_Feedback fiber.Map                                       `json:"alerts"`
// 		TransactionAlerts    []tracenetwork.Credittransfer_feedback_response `json:"transactionAlerts"`
// 	}{
// 		TraceAlerts_Feedback: fiber.Map{
// 			"InstructionsID": transactions[0].Instruction_id,
// 			"ReferenceID":    transactions[0].Reference_id,
// 			"DECISIONDATE":   requestTrigger,
// 			"TRACE_ALERT":    transactions[0].Trace_alert,
// 			"Feedback":       network.Feedback,
// 		},
// 		TransactionAlerts: transactions,
// 	}

//		return c.JSON(responseBody)
//	}
func Feedbackcredit(c *fiber.Ctx) error {
	network := &tracenetwork.Credittransfer_feedback{}
	Uniqueidcredittransfer := Iftgeneratefeedbackcredittransfer(32)

	requestTrigger := time.Now().Format("2006-01-02 15:04:05")
	if err := c.BodyParser(network); err != nil {
		//400
		loggers.Creditransferfeedbacklogs(c.Path(), "folderName", Uniqueidcredittransfer, "(Method Not Allowed - 400)", "null", requestTrigger, "null", "null", "null", "null", "null", "NETWORK_FINANCIAL_CRIME", "null", "null", "null")
		return errorhandling.Bad_Request(c, "The request contains a bad payload")
	}

	var transactions []tracenetwork.Credittransfer_feedback_response
	if err := database.DBConn.Debug().Raw("SELECT * FROM trace_alerts.logscredittransfer WHERE uniqueidcredittransfer = ? AND sourcetxntype = ? AND trace_alert = ? AND alerttype = ? ", network.Uniqueidcredittransfer, network.Tracetype, network.Trace_alert, network.Alerttype).Find(&transactions).Error; err != nil {
		return errorhandling.Internal_Server_Error(c, "Internal server error")
	}

	var Feedback string
	rows, err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.logscredittransfer WHERE sourcetxntype = 'FRAUD'`).Rows()
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		Feedback = "ACTIONED_MULE"
	} else {
		rows, err = database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.logscredittransfer WHERE sourcetxntype IS NULL OR sourcetxntype = 'NONE'`).Rows()
		if err != nil {
			panic(err)
		}
		defer rows.Close()

		if rows.Next() {
			Feedback = "NOT_INVESTIGATED"
		}
	}

	if Feedback == "" {
		Feedback = "null"
	}

	if len(transactions) == 0 {
		//404
		loggers.Creditransferfeedbacklogs(c.Path(), "folderName", Uniqueidcredittransfer, "(Method Not Allowed - 404)", "null", requestTrigger, "null", "null", "null", "null", "null", "NETWORK_FINANCIAL_CRIME", "null", "null", "null")
		return errorhandling.Url_Not_Found(c, "No data found for the specified date")
	}

	responseBody := struct {
		Alerts            fiber.Map                                       `json:"alerts"`
		TransactionAlerts []tracenetwork.Credittransfer_feedback_response `json:"transactionAlerts"`
	}{}

	responseBody.Alerts = fiber.Map{
		"InstructionsID": transactions[0].Instruction_id,
		"ReferenceID":    transactions[0].Reference_id,
		"DECISIONDATE":   requestTrigger,
		"FEEDBACK":       Feedback,
	}

	var errorresp errors.Errorresp

	for i := 0; i < len(transactions); i++ {
		Errors := errorresp.Errorresponse
		transaction := transactions[i]

		Uniqueidcredittransfer := transaction.Uniqueidcredittransfer
		Instruction_id := transaction.Instruction_id
		Reference_id := transaction.Reference_id
		Transaction_type := transaction.Transaction_type
		Status := transaction.Status
		Reason_code := transaction.Reason_code
		Local_instrument := transaction.Local_instrument
		Trace_alert := transactions[0].Trace_alert
		Sourcetxntype := transactions[0].Sourcetxntype
		Alerttype := transaction.Alerttype
		message := fmt.Sprintf(" Success %s ", Errors)

		loggers.Creditransferfeedbacklogs(c.Path(), "folderName", Uniqueidcredittransfer, message, Instruction_id, requestTrigger, Transaction_type, Status, Reason_code, Local_instrument, Reference_id, Trace_alert, Sourcetxntype, Alerttype, Feedback)
	}

	responseBody.TransactionAlerts = transactions

	return c.JSON(responseBody)
}
func Iftgeneratefeedbackcredittransfer(length int) string {
	rand.Seed(time.Now().UnixNano())

	hexDigits := "0123456789abcdef"
	result := make([]byte, length)
	for i := range result {
		result[i] = hexDigits[rand.Intn(len(hexDigits))]
	}
	return string(result)
}

//200
//404
//400
//405
//429

//401
//403
