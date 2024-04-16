package traceandalert

import (
	"math/rand"
	"time"
	"tracealert/middleware/loggers"
	errorHandling "tracealert/pkg/models/errorHandling"
	"tracealert/pkg/models/tracenetwork"
	"tracealert/pkg/utils/go-utils/database"

	"github.com/gofiber/fiber/v2"
)

func FeedbackFraud(c *fiber.Ctx) error {
	network := &tracenetwork.CredittransferFeedback{}
	UniqueidCredittransfer := Iftgeneratefeedbackcredittrans(32)

	requestTrigger := time.Now().Format("2006-01-02 15:04:05")
	if err := c.BodyParser(network); err != nil {
		//400
		loggers.CreditransferFeedbackLogs(c.Path(), "folderName", UniqueidCredittransfer, "(Method Not Allowed - 400)", "null", requestTrigger, "null", "null", "null", "null", "null", "NETWORK_FINANCIAL_CRIME", "null", "null", "null", "null", "null")
		return errorHandling.Bad_Request(c, "The request contains a bad payload")
	}

	var transactions []tracenetwork.CredittransferFeedbackResponse
	if err := database.DBConn.Debug().Raw("SELECT * FROM trace_alerts.logscredittransfer WHERE uniqueid_credittransfer = ? AND source_txn_type = ? AND trace_alert = ? AND alert_type = ? ", network.UniqueidCredittransfer, network.TraceType, network.TraceAlert, network.AlertType).Find(&transactions).Error; err != nil {
		return errorHandling.Internal_Server_Error(c, "Internal server error")
	}

	var Feedback string
	var Lock string
	rows, err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.logscredittransfer WHERE source_txn_type = 'FRAUD'`).Rows()
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		Feedback = "ACTIONED_MULE"
		Lock = "lock"
	} else {
		rows, err = database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.logscredittransfer WHERE source_txn_type IS NULL OR source_txn_type = 'NONE'`).Rows()
		if err != nil {
			panic(err)
		}
		defer rows.Close()

		if rows.Next() {
			Feedback = "NOT_INVESTIGATED"
			Lock = "unlock"
		}

	}

	if Feedback == "" {
		Feedback = "null"
	}

	if len(transactions) == 0 {
		//404
		loggers.CreditransferFeedbackLogs(c.Path(), "folderName", UniqueidCredittransfer, "(Method Not Allowed - 404)", "null", requestTrigger, "null", "null", "null", "null", "null", "NETWORK_FINANCIAL_CRIME", "null", "null", "null", "null", "null")
		return errorHandling.Url_Not_Found(c, "No data found for the specified date")
	}

	responseBody := struct {
		Alerts            fiber.Map                                     `json:"alerts"`
		TransactionAlerts []tracenetwork.CredittransferFeedbackResponse `json:"transactionAlerts"`
	}{}

	responseBody.Alerts = fiber.Map{
		"InstructionsID": transactions[0].InstructionId,
		"ReferenceID":    transactions[0].ReferenceId,
		"DECISIONDATE":   requestTrigger,
		"FEEDBACK":       Feedback,
		"MULE_LOCK":      Lock,
	}

	return c.JSON(responseBody)
}

func Iftgeneratefeedbackcredittrans(length int) string {
	rand.Seed(time.Now().UnixNano())

	hexDigits := "0123456789abcdef"
	result := make([]byte, length)
	for i := range result {
		result[i] = hexDigits[rand.Intn(len(hexDigits))]
	}
	return string(result)
}
