package healthchecks

import (
	"time"
	"tracealert/pkg/utils/go-utils/database"

	"github.com/gofiber/fiber/v2"
)

type Logsfeedback struct {
	Uniqueidfeedback string `json:"uniqueidfeedback"`
	Alertid          string `json:"alertid"`
	Feedbackid       string `json:"feedbackid"`
	Errors           string `json:"errors"`
}

type LogsAlert struct {
	Uniqueidalertaccount string `json:"uniqueidalertaccount"`
	ID                   string `json:"id"`
	Message              string `json:"message"`
	Networkalertid       string `json:"networkalertid"`
	Accountid            string `json:"accountid"`
	Networkid            string `json:"networkid"`
	Owningbankid         string `json:"owningbankid"`
	Owningbankname       string `json:"owningbankname"`
	Decisiondate         string `json:"decisiondate"`
	Parentalertid        string `json:"parentalertid"`
}

type Logstracevisualisation struct {
	Uniqueidvisualisation string `json:"uniqueidvisualisation"`
	Networkalertid        string `json:"networkalertid"`
	Traceid               string `json:"traceid"`
	Errors                string `json:"errors"`
}

type Logstracenetwork struct {
	Uniqueidnetwork string    `json:"uniqueidnetwork"`
	Txnid_RB        string    `json:"txnid_RB"`
	Txnid           string    `json:"txnid"`
	Networkid       string    `json:"networkid"`
	Sourcetxnid     string    `json:"sourcetxnid"`
	Decisiondate    time.Time `json:"decisiondate"`
	Errors          string    `json:"errors"`
}

type Logsalerttransaction struct {
	Uniqueidalertaccount string    `json:"uniqueidalertaccount"`
	ID                   string    `json:"id"`
	Errors               string    `json:"errors"`
	Txnid                string    `json:"txnid"`
	Networkalertid       string    `json:"networkalertid"`
	Networkid            string    `json:"networkid"`
	Sourceid             string    `json:"sourceid"`
	Destid               string    `json:"destid"`
	Sourcebankid         string    `json:"sourcebankid"`
	Sourcebankname       string    `json:"sourcebankname"`
	Destbankid           string    `json:"destbankid"`
	Destbankname         string    `json:"destbankname"`
	Parentalertid        string    `json:"parentalertid"`
	Decisiondate         time.Time `json:"decisiondate"`
}

type Logscredittransfer struct {
	Uniqueidcredittransfer string `json:"uniqueidcredittransfer"`
	Message                string `json:"message"`
	Trace_alert            string `json:"trace_alert"`
	Sourcetxntype          string `json:"sourcetxntype"`
	Transaction_type       string `json:"transaction_type"`
	Status                 string `json:"status"`
	Reason_code            string `json:"reason_code"`
	Local_instrument       string `json:"local_instrument"`
	Instruction_id         string `json:"instruction_id"`
	Reference_id           string `json:"reference_id"`
	Requesttrigger         string `json:"requesttrigger"`
	Alerttype              string `json:"alerttype"`
	Feedback               string `json:"feedback"`
}

type LogsFeedback_credittransfer struct {
	Uniqueidcredittransfer string `json:"uniqueidcredittransfer"`
	Message                string `json:"message"`
	Trace_alert            string `json:"trace_alert"`
	Sourcetxntype          string `json:"sourcetxntype"`
	Transaction_type       string `json:"transaction_type"`
	Status                 string `json:"status"`
	Reason_code            string `json:"reason_code"`
	Local_instrument       string `json:"local_instrument"`
	Instruction_id         string `json:"instruction_id"`
	Reference_id           string `json:"reference_id"`
	Requesttrigger         string `json:"requesttrigger"`
	Alerttype              string `json:"alerttype"`
	Feedback               string `json:"feedback"`
}

func Feedbackinfo(c *fiber.Ctx) error {
	ctResult := &[]Logsfeedback{}
	database.DBConn.Raw("SELECT * FROM trace_alerts.logsfeedback WHERE 1=1").Scan(ctResult)

	return c.Render("feedback", fiber.Map{
		"title":      "Feedback",
		"ctResponse": ctResult,
	})
}

func Alertaccinfo(c *fiber.Ctx) error {
	ctResult := &[]LogsAlert{}
	database.DBConn.Raw("SELECT * FROM trace_alerts.logsalertaccount WHERE 1=1").Scan(ctResult)

	return c.Render("alertaccount", fiber.Map{
		"title":      "Alert Account",
		"ctResponse": ctResult,
	})
}

func Tracevisinfo(c *fiber.Ctx) error {
	ctResult := &[]Logstracevisualisation{}
	database.DBConn.Raw("SELECT * FROM trace_alerts.logstracevisualisation WHERE 1=1").Scan(ctResult)

	return c.Render("tracevisualisation", fiber.Map{
		"title":      "Trace Visualisation",
		"ctResponse": ctResult,
	})
}

func Tracenetworkinfo(c *fiber.Ctx) error {
	ctResult := &[]Logstracenetwork{}
	database.DBConn.Raw("SELECT * FROM trace_alerts.logstracenetwork WHERE 1=1").Scan(ctResult)

	return c.Render("tracenetwork", fiber.Map{
		"title":      "Trace Alerts",
		"ctResponse": ctResult,
	})
}

func Alerttransactioninfo(c *fiber.Ctx) error {
	ctResult := &[]Logsalerttransaction{}
	database.DBConn.Raw("SELECT * FROM trace_alerts.logsalerttransactions WHERE 1=1").Scan(ctResult)

	return c.Render("alerttransaction", fiber.Map{
		"title":      "Alert Transaction",
		"ctResponse": ctResult,
	})
}

func Credittransferinfo(c *fiber.Ctx) error {
	ctResult := &[]Logscredittransfer{}
	database.DBConn.Raw("SELECT * FROM trace_alerts.logscredittransfer WHERE 1=1 ").Scan(ctResult)

	return c.Render("credittransfer_trace", fiber.Map{
		"title":      "credit_transfer",
		"ctResponse": ctResult,
	})
}

func FeedbackCredittransferinfo(c *fiber.Ctx) error {
	ctResult := &[]LogsFeedback_credittransfer{}
	database.DBConn.Raw("SELECT * FROM trace_alerts.logsFeedback_credittransfer WHERE 1=1 ").Scan(ctResult)

	return c.Render("credittransfer_feedback", fiber.Map{
		"title":      "Feedback_credit_transfer",
		"ctResponse": ctResult,
	})
}
func ShowPage(c *fiber.Ctx) error {

	return c.Render("tracepostman", fiber.Map{
		"titlePage": "SAMPLE PAGE",
	})

}
