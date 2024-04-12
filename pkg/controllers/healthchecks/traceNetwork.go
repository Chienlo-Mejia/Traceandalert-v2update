package healthchecks

import (
	"fmt"
	"math/rand"
	"time"
	"tracealert/middleware/loggers"
	"tracealert/pkg/models"
	errorhandling "tracealert/pkg/models/errorHandling"
	"tracealert/pkg/models/errors"
	"tracealert/pkg/utils/go-utils/database"

	"github.com/gofiber/fiber/v2"
)

func TraceNetwork(c *fiber.Ctx) error {
	network := &models.Request_trace{}
	Uniqueidnetwork := Iftgeneratenetwork(32)
	if err := c.BodyParser(&network); err != nil {
		// 400
		loggers.Tracenetwork(c.Path(), "folderName", Uniqueidnetwork, network.Txnid, " The request contains a bad payloads", "(Bad Request - 400)", "null", "null", "null", time.Now(), "null", "null", "null", "null", "null", "null", "null", time.Now(), "null", "null", "null", "null", "null", "null", "null", "null", "null", "null", "null", "null", time.Now(), false, "null", "null")
		return errorhandling.Bad_Request(c, "The request contains a bad payloads")

	}

	var transactions1 []models.Transaction_Response
	var transactions2 []models.Trace_AccountAlert
	var transactions3 []models.NetworkResponse

	//error response
	var errorresp errors.Errorresp

	if err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.trace_alerts WHERE txnid = ? AND tracetype = ? `, network.Txnid, network.Sourcetxntype).Scan(&transactions1).Error; err != nil {

		//500
		loggers.Tracenetwork(c.Path(), "folderName", Uniqueidnetwork, network.Txnid, " Internal server error", "(Bad Request - 500)", "null", "null", "null", time.Now(), "null", "null", "null", "null", "null", "null", "null", time.Now(), "null", "null", "null", "null", "null", "null", "null", "null", "null", "null", "null", "null", time.Now(), false, "null", "null")
		return errorhandling.Internal_Server_Error(c, "Internal server error")
	}

	if err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.trace_account_alerts WHERE networkid = ? AND tracetype = ?`, network.Txnid, network.Sourcetxntype).Scan(&transactions2).Error; err != nil {
		//500
		loggers.Tracenetwork(c.Path(), "folderName", Uniqueidnetwork, network.Txnid, " Internal server error", "(Bad Request - 500)", "null", "null", "null", time.Now(), "null", "null", "null", "null", "null", "null", "null", time.Now(), "null", "null", "null", "null", "null", "null", "null", "null", "null", "null", "null", "null", time.Now(), false, "null", "null")
		return errorhandling.Internal_Server_Error(c, "Internal server error")
	}

	if err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.trace_vizurl WHERE sourcetxnid = ? AND sourcetxntype = ?`, network.Txnid, network.Sourcetxntype).Scan(&transactions3).Error; err != nil {
		//500
		loggers.Tracenetwork(c.Path(), "folderName", Uniqueidnetwork, network.Txnid, "Internal server error", "(Bad Request - 500)", "null", "null", "null", time.Now(), "null", "null", "null", "null", "null", "null", "null", time.Now(), "null", "null", "null", "null", "null", "null", "null", "null", "null", "null", "null", "null", time.Now(), false, "null", "null")
		return errorhandling.Internal_Server_Error(c, "Internal server error")
	}

	if len(transactions1) == 0 || len(transactions2) == 0 || len(transactions3) == 0 {
		// 400
		loggers.Tracenetwork(c.Path(), "folderName", Uniqueidnetwork, network.Txnid, "The request contains a bad payload", "(Bad Request - 400)", "null", "null", "null", time.Now(), "null", "null", "null", "null", "null", "null", "null", time.Now(), "null", "null", "null", "null", "null", "null", "null", "null", "null", "null", "null", "null", time.Now(), false, "null", "null")
		return errorhandling.Bad_Request(c, "The request contains a bad payload")

	}

	//check if exist
	exists, err := checktxnidexist(network.Txnid)
	if err != nil {
		loggers.Tracenetwork(c.Path(), "folderName", Uniqueidnetwork, network.Txnid, " Internal server error", "(Bad Request - 500)", "null", "null", "null", time.Now(), "null", "null", "null", "null", "null", "null", "null", time.Now(), "null", "null", "null", "null", "null", "null", "null", "null", "null", "null", "null", "null", time.Now(), false, "null", "null")
		return errorhandling.Internal_Server_Error(c, "Database error")
	}
	if exists {
		loggers.Tracenetwork(c.Path(), "folderName", Uniqueidnetwork, network.Txnid, "The request contains a bad payload", "(Bad Request - 400)", "null", "null", "null", time.Now(), "null", "null", "null", "null", "null", "null", "null", time.Now(), "null", "null", "null", "null", "null", "null", "null", "null", "null", "null", "null", "null", time.Now(), false, "null", "null")
		return errorhandling.Conflict(c, "Alert ID already exists")
	}

	// alert := fiber.Map{
	// 	"id":        transactions1[0].ID,
	// 	"Time":      transactions1[0].Time,
	// 	"networkID": transactions1[0].Networkalertid,
	// }

	// return c.JSON(fiber.Map{
	// 	"alerts":            alert,
	// 	"accountAlerts":     transactions1,
	// 	"transactionAlerts": transactions2,
	// 	"network":           transactions3,
	// })

	type alert struct {
		Alerts fiber.Map `json:"alerts"`

		TransactionAlerts []models.Transaction_Response `json:"transactionAlerts"`
		AccountAlerts     []models.Trace_AccountAlert   `json:"accountAlerts"`
		Network           []models.NetworkResponse      `json:"network"`
	}

	responseBody := alert{
		Alerts: fiber.Map{
			"id":        transactions1[0].ID,
			"Time":      transactions1[0].Time,
			"networkID": transactions1[0].Networkid,
		},
		TransactionAlerts: transactions1,
		AccountAlerts:     transactions2,
		Network:           transactions3,
	}

	//------------------- logs -------------------
	Txnid_RB := network.Txnid
	//trace_alerts
	ID_alert := transactions1[0].ID
	Txnid := transactions1[0].Txnid
	Networkalertid_alert := transactions1[0].Networkalertid
	Networkid_alert := transactions1[0].Networkid
	Sourceid := transactions1[0].Sourceid
	Destid := transactions1[0].Destid
	Sourcebankid := transactions1[0].Sourcebankid
	Sourcebankname := transactions1[0].Sourcebankname
	Destbankid := transactions1[0].Destbankid
	Destbankname := transactions1[0].Destbankname
	Tracetype := transactions1[0].Tracetype
	Parentalertid_alert := transactions1[0].Parentalertid
	Decisiondate_alert := transactions1[0].Decisiondate
	Status := transactions1[0].Status
	Nextpaginationtoken := transactions1[0].Nextpaginationtoken
	Previouspaginationtoken := transactions1[0].Previouspaginationtoken
	//trace_account_alert
	Id := transactions2[0].Id
	Networkalertid_traceacc := transactions2[0].Networkalertid
	Accountid := transactions2[0].Accountid
	Networkid := transactions2[0].Networkid
	Owningbankid := transactions2[0].Owningbankid
	Owningbankname := transactions2[0].Owningbankname
	Time := transactions2[0].Time
	//vizurl
	Vizurl := transactions3[0].Vizurl
	Errors := errorresp.Errorresponse
	Sourcetxnid := transactions3[0].Sourcetxnid
	Sourcetxntype := transactions3[0].Sourcetxntype
	Parentalertid := transactions3[0].Parentalertid
	Decisiondate := transactions3[0].Decisiondate
	Networkalertid := transactions1[0].Networkalertid

	message := fmt.Sprintf(" Success %s ", Errors)
	loggers.Tracenetwork(c.Path(), "folderName", Uniqueidnetwork, Txnid_RB, Vizurl, message, Sourcetxnid, Sourcetxntype, Parentalertid, Decisiondate, Networkalertid, Id, Networkalertid_traceacc, Accountid, Networkid, Owningbankid, Owningbankname, Time, ID_alert, Txnid, Networkalertid_alert, Networkid_alert, Sourceid, Destid, Sourcebankid, Sourcebankname, Destbankid, Destbankname, Tracetype, Parentalertid_alert, Decisiondate_alert, Status, Nextpaginationtoken, Previouspaginationtoken)

	return c.JSON(responseBody)

}

func checktxnidexist(Txnid string) (bool, error) {
	var count int
	err := database.DBConn.Raw("SELECT COUNT(*) FROM trace_alerts.trace_alerts WHERE txnid = ?", Txnid).Scan(&count).Error
	if err != nil {
		return false, err
	}
	return count > 1, nil
}

func Iftgeneratenetwork(length int) string {
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
