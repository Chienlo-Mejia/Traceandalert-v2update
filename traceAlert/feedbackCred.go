package traceandalert

import (
	"fmt"
	"math/rand"
	"time"
	"tracealert/middleware/loggers"
	notifications "tracealert/notifications"
	errorHandling "tracealert/pkg/models/errorHandling"
	"tracealert/pkg/models/errors"
	"tracealert/pkg/models/tracenetwork"
	"tracealert/pkg/utils/go-utils/database"

	"github.com/gofiber/fiber/v2"
	"github.com/lib/pq"
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

func FeedbackCredit(c *fiber.Ctx) error {
	network := &tracenetwork.Credittransfer_feedback{}
	UniqueidCredittransfer := Iftgeneratefeedbackcredittransfer(32)

	requestTrigger := time.Now().Format("2006-01-02 15:04:05")
	if err := c.BodyParser(network); err != nil {
		//400
		loggers.CreditransferFeedbackLogs(c.Path(), "folderName", UniqueidCredittransfer, "(Method Not Allowed - 400)", "null", requestTrigger, "null", "null", "null", "null", "null", "NETWORK_FINANCIAL_CRIME", "null", "null", "null", "null", "null")
		return errorHandling.Bad_Request(c, "The request contains a bad payload")
	}

	var transactions []tracenetwork.Credittransfer_feedback_response
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
		Alerts            fiber.Map                                       `json:"alerts"`
		TransactionAlerts []tracenetwork.Credittransfer_feedback_response `json:"transactionAlerts"`
	}{}

	responseBody.Alerts = fiber.Map{
		"InstructionsID": transactions[0].InstructionId,
		"ReferenceID":    transactions[0].ReferenceId,
		"DECISIONDATE":   requestTrigger,
		"FEEDBACK":       Feedback,
		"MULE_LOCK":      Lock,
	}

	if Feedback == "ACTIONED_MULE" {
		emailData := notifications.EmailMessage{
			Email:   []string{"cpe.anista.bienmarc@gmail.com", "chienlo.mejia@fortress-asya.com", "rico.vergara@fortress-asya.com", "chienlomejia660@gmail.com", "chichimejia660@gmail.com", "rosendo.ocubillo@fortress-asya.com", "josephmejia660@gmail.com", "mapeebun@gmail.com", "arren.agravante@fortress-asya.com", "bien.anista@fortress-asya.com"},
			Message: "Urgent Problem",
			Detail:  "We are currently experiencing technical difficulties with our payment processing system. This may cause delays in processing transactions. Our team is actively working to resolve the issue as quickly as possible. We apologize for any inconvenience this may cause.",
		}
		if err := notifications.Email(c, emailData); err != nil {
			return errorHandling.Internal_Server_Error(c, "Failed to send email")
		}

		result := database.DBConn.Exec("SELECT trace_alerts.insert_notification_email(?, ?, ? ,?)", pq.Array(emailData.Email), emailData.Message, emailData.Detail, requestTrigger)
		if result.Error != nil {
			fmt.Println("Failed to execute insert_notification_email function:", result.Error)
			return result.Error
		}
	}

	// if Feedback == "ACTIONED_MULE" {
	// 	smsData := notifications.MessagePayload{
	// 		Msg:     "CUTE si CHIENLO & RICO",
	// 		From:    "CARD RBI",
	// 		AppCode: "K2cPlus_RBI",
	// 		To:      "09662733608",
	// 	}

	// 	if err := notifications.SmsTelerivet(c, smsData); err != nil {
	// 		fmt.Println("Failed to send SMS:", err)
	// 		return errorHandling.Internal_Server_Error(c, "Failed to send SMS")
	// 	}

	// 	result := database.DBConn.Exec("SELECT trace_alerts.insert_notification_sms(?, ?, ?, ?, ?)", smsData.Msg, smsData.From, smsData.AppCode, smsData.To, requestTrigger)
	// 	if result.Error != nil {
	// 		fmt.Println("Failed to insert message response into the database:", result.Error)
	// 		return result.Error
	// 	}
	// }

	var errorresp errors.Errorresp

	for i := 0; i < len(transactions); i++ {
		Errors := errorresp.Errorresponse
		transaction := transactions[i]

		UniqueidCredittransfer := transaction.UniqueidCredittransfer
		Instruction_id := transaction.InstructionId
		Reference_id := transaction.ReferenceId
		Transaction_type := transaction.TransactionType
		Status := transaction.Status
		Reason_code := transaction.ReasonCode
		Local_instrument := transaction.LocalInstrument
		Trace_alert := transaction.TraceAlert
		Sourcetxntype := transaction.SourceTxnType
		Alerttype := transaction.AlertType
		message := fmt.Sprintf(" Success %s ", Errors)
		SenderAccount := transaction.SenderAccount

		loggers.CreditransferFeedbackLogs(c.Path(), "folderName", UniqueidCredittransfer, message, Instruction_id, requestTrigger, Transaction_type, Status, Reason_code, Local_instrument, Reference_id, Trace_alert, Sourcetxntype, Alerttype, Feedback, Lock, SenderAccount)
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

// func Email(c *fiber.Ctx, emailData EmailMessage) error {
// 	pass := "ooqz lcqj rfki dyer"
// 	from := "chichimejia660@gmail.com"

// 	smtpHost := "smtp.gmail.com"
// 	smtpPort := "587"

// 	auth := smtp.PlainAuth("", from, pass, smtpHost)

// 	var mu sync.Mutex
// 	var errors []error
// 	emailSize := 30000
// 	numEmail := (len(emailData.To) + emailSize - 1) / emailSize

// 	for i := 0; i < numEmail; i++ {
// 		startIdx := i * emailSize
// 		endIdx := (i + 1) * emailSize
// 		if endIdx > len(emailData.To) {
// 			endIdx = len(emailData.To)
// 		}

// 		batch := emailData.To[startIdx:endIdx]

// 		var wg sync.WaitGroup

// 		for _, to := range batch {
// 			wg.Add(1)
// 			go func(to string) {
// 				defer wg.Done()

// 				msg := "From: " + from + "\r\n" +
// 					"To: " + to + "\r\n" +
// 					"Subject: Customized Alert: Action Required\r\n" +
// 					"MIME-version: 1.0;\r\n" +
// 					"Content-Type: text/html; charset=\"UTF-8\";\r\n\r\n" +
// 					"<html><head><style>" +
// 					"body { font-family: 'Arial', sans-serif; background-color: #f7f7f7; color: #333; margin: 0; padding: 0; }" +
// 					"#container { max-width: 600px; margin: 0 auto; padding: 20px; background-color: #fff; border-radius: 10px; box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1); }" +
// 					"h1 { color: #003366; font-size: 24px; margin-bottom: 20px; }" +
// 					"p { font-size: 16px; margin-bottom: 10px; }" +
// 					".urgent { color: #e74c3c; font-weight: bold; }" +
// 					".highlight { color: #1f497d; }" +
// 					".signature { font-style: italic; }" +
// 					"</style></head><body>" +
// 					"<div id='container'>" +
// 					"<h1 style='text-align: center;'>Customized Alert: Action Required</h1>" +
// 					"<p>Dear: " + to + ",</p>" +
// 					"<p class='urgent'>Urgent action is needed:</p>" +
// 					"<p class='highlight'>Message: " + emailData.Message + "</p>" +
// 					"<p class='highlight'>Privacy Notice: " + emailData.Detail + "</p>" +
// 					"<p style='text-align: right;' class='signature'>Best regards,<br/>INSTAPAY Team</p>" +
// 					"</div></body></html>"

// 				err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, []byte(msg))
// 				if err != nil {
// 					fmt.Printf("Email sending failed for recipient %s: %s\n", to, err)
// 					mu.Lock()
// 					errors = append(errors, err)
// 					mu.Unlock()
// 				} else {
// 					fmt.Println("Email sent successfully to", to)
// 				}
// 			}(to)
// 		}

// 		wg.Wait()
// 	}

// 	if len(errors) > 0 {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"message": "Error sending emails.",
// 			"errors":  errors,
// 		})
// 	}

// 	return nil
// }

//200
//404
//400
//405
//429

//401
//403
