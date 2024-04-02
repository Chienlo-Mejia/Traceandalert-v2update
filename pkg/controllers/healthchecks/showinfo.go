package healthchecks

import (
	"time"
	"tracealert/pkg/utils/go-utils/database"

	"github.com/gofiber/fiber/v2"
)

type Logsfeedback struct {
	UniqueidFeedback string `json:"uniqueidFeedback"`
	AlertId          string `json:"alertId"`
	FeedbackId       string `json:"feedbackId"`
	Errors           string `json:"errors"`
}

type LogsAlert struct {
	UniqueidAlertaccount string `json:"uniqueidAlertaccount"`
	ID                   string `json:"id"`
	Message              string `json:"message"`
	NetworkalertId       string `json:"networkalertId"`
	AccountId            string `json:"accountId"`
	NetworkId            string `json:"networkId"`
	OwningbankId         string `json:"owningbankId"`
	OwningbankName       string `json:"owningbankName"`
	DecisionDate         string `json:"decisionDate"`
	ParentalertId        string `json:"parentalertId"`
}

type Logstracevisualisation struct {
	UniqueidVisualisation string `json:"uniqueidVisualisation"`
	NetworkalertId        string `json:"networkalertId"`
	TraceId               string `json:"traceId"`
	Errors                string `json:"errors"`
}

type Logstracenetwork struct {
	UniqueidNetwork string    `json:"uniqueidNetwork"`
	TxnidRb         string    `json:"txnidRb"`
	TxnId           string    `json:"txnId"`
	NetworkId       string    `json:"networkId"`
	SourcetxnId     string    `json:"sourcetxnId"`
	DecisionDate    time.Time `json:"decisionDate"`
	Errors          string    `json:"errors"`
}

type Logsalerttransaction struct {
	UniqueidAlertaccount string    `json:"uniqueidAlertaccount"`
	ID                   string    `json:"id"`
	Errors               string    `json:"errors"`
	TxnId                string    `json:"txnId"`
	NetworkalertId       string    `json:"networkalertId"`
	NetworkId            string    `json:"networkId"`
	SourceId             string    `json:"sourceId"`
	DestId               string    `json:"destId"`
	SourcebankId         string    `json:"sourcebankId"`
	SourcebankName       string    `json:"sourcebankName"`
	DestbankId           string    `json:"destbankId"`
	DestbankName         string    `json:"destbankName"`
	ParentalertId        string    `json:"parentalertId"`
	DecisionDate         time.Time `json:"decisionDate"`
}

type Logscredittransfer struct {
	UniqueidCredittransfer string `json:"uniqueidCredittransfer"`
	Message                string `json:"message"`
	TraceAlert             string `json:"traceAlert"`
	SourceTxnType          string `json:"sourceTxnType"`
	TransactionType        string `json:"transactionType"`
	Status                 string `json:"status"`
	ReasonCode             string `json:"reasonCode"`
	LocalInstrument        string `json:"localInstrument"`
	InstructionId          string `json:"instructionId"`
	ReferenceId            string `json:"referenceId"`
	RequestTrigger         string `json:"requestTrigger"`
	AlertType              string `json:"alertType"`
	Feedback               string `json:"feedback"`
}

type LogsFeedback_credittransfer struct {
	UniqueidCredittransfer string `json:"uniqueidCredittransfer"`
	Message                string `json:"message"`
	TraceAlert             string `json:"traceAlert"`
	SourceTxnType          string `json:"sourceTxnType"`
	TransactionType        string `json:"transactionType"`
	Status                 string `json:"status"`
	ReasonCode             string `json:"reasonCode"`
	LocalInstrument        string `json:"localInstrument"`
	InstructionId          string `json:"instructionId"`
	ReferenceId            string `json:"referenceId"`
	RequestTrigger         string `json:"requestTrigger"`
	AlertType              string `json:"alertType"`
	Feedback               string `json:"feedback"`
}
type TraceInfo struct {
	Id             string `json:"id"`
	Message        string `json:"message"`
	Requesttrigger string `json:"requesttrigger"`
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

func TracePostmanMobileBrowser(c *fiber.Ctx) error {
	ctResult := &[]TraceInfo{}
	database.DBConn.Raw("SELECT * FROM trace_alerts.logstracetrans_postmanmobilephonebrowser WHERE 1=1 ").Scan(ctResult)

	return c.Render("tracePostmanMobileBrowser", fiber.Map{
		"title":      "TracePostmanMobileBrowser",
		"ctResponse": ctResult,
	})
}
