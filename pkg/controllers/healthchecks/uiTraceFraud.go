package healthchecks

import (
	tracefraud "tracealert/pkg/models/traceFraud"
	"tracealert/pkg/utils/go-utils/database"

	"github.com/gofiber/fiber/v2"
)

// UI TRACE FRAUD
func Feedbackinfo(c *fiber.Ctx) error {
	ctResult := &[]tracefraud.Logsfeedback{}
	database.DBConn.Raw("SELECT * FROM trace_alerts.logsfeedback WHERE 1=1").Scan(ctResult)

	return c.Render("feedback", fiber.Map{
		"title":      "Feedback",
		"ctResponse": ctResult,
	})
}

func Alertaccinfo(c *fiber.Ctx) error {
	ctResult := &[]tracefraud.LogsAlert{}
	database.DBConn.Raw("SELECT * FROM trace_alerts.logsalertaccount WHERE 1=1").Scan(ctResult)

	return c.Render("alertAccount", fiber.Map{
		"title":      "Alert Account",
		"ctResponse": ctResult,
	})
}

func Tracevisinfo(c *fiber.Ctx) error {
	ctResult := &[]tracefraud.Logstracevisualisation{}
	database.DBConn.Raw("SELECT * FROM trace_alerts.logstracevisualisation WHERE 1=1").Scan(ctResult)

	return c.Render("traceVisualisation", fiber.Map{
		"title":      "Trace Visualisation",
		"ctResponse": ctResult,
	})
}

func Tracenetworkinfo(c *fiber.Ctx) error {
	ctResult := &[]tracefraud.Logstracenetwork{}
	database.DBConn.Raw("SELECT * FROM trace_alerts.logstracenetwork WHERE 1=1").Scan(ctResult)

	return c.Render("traceNetwork", fiber.Map{
		"title":      "Trace Alerts",
		"ctResponse": ctResult,
	})
}

func Alerttransactioninfo(c *fiber.Ctx) error {
	ctResult := &[]tracefraud.Logsalerttransaction{}
	database.DBConn.Raw("SELECT * FROM trace_alerts.logsalerttransactions WHERE 1=1").Scan(ctResult)

	return c.Render("alertTransaction", fiber.Map{
		"title":      "Alert Transaction",
		"ctResponse": ctResult,
	})
}

func Credittransferinfo(c *fiber.Ctx) error {
	ctResult := &[]tracefraud.Logscredittransfer{}
	database.DBConn.Raw("SELECT * FROM trace_alerts.logscredittransfer WHERE 1=1 ").Scan(ctResult)

	return c.Render("creditTransferTrace", fiber.Map{
		"title":      "credit_transfer",
		"ctResponse": ctResult,
	})
}

func FeedbackCredittransferinfo(c *fiber.Ctx) error {
	ctResult := &[]tracefraud.LogsFeedback_credittransfer{}
	database.DBConn.Raw("SELECT * FROM trace_alerts.logsFeedback_credittransfer WHERE 1=1 ").Scan(ctResult)

	return c.Render("creditTransferFeedback", fiber.Map{
		"title":      "Feedback_credit_transfer",
		"ctResponse": ctResult,
	})
}
func ShowPage(c *fiber.Ctx) error {

	return c.Render("tracePostman", fiber.Map{
		"titlePage": "SAMPLE PAGE",
	})

}

func TracePostmanMobileBrowser(c *fiber.Ctx) error {
	ctResult := &[]tracefraud.TraceInfo{}
	database.DBConn.Raw("SELECT * FROM trace_alerts.logstracetrans_postmanmobilephonebrowser WHERE 1=1 ").Scan(ctResult)

	return c.Render("tracePostmanMobileBrowser", fiber.Map{
		"title":      "TracePostmanMobileBrowser",
		"ctResponse": ctResult,
	})
}
