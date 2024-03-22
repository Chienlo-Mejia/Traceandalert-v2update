package healthchecks

import (
	"math/rand"
	"time"
	"tracealert/middleware/loggers"
	"tracealert/pkg/models"
	errorhandling "tracealert/pkg/models/errorHandling"
	"tracealert/pkg/utils/go-utils/database"

	"github.com/gofiber/fiber/v2"
)

var requestCountsalert_transaction_ID int

const allowedRatesalert_transaction_ID = 5

func checkRateLimitsalert_transaction_ID() bool {
	requestCountsalert_transaction_ID++
	if requestCountsalert_transaction_ID > allowedRatesalert_transaction_ID {
		requestCountsalert_transaction_ID = 0
		return true
	}
	return false
}

func GetTransacInfo(c *fiber.Ctx) error {

	info := &models.TransacAlertRequest{}
	resp := &models.ResponseTransac{}
	Uniqueidalertaccountid := Iftgeneratealerttransactionid(32)

	if checkRateLimitsalert_transaction_ID() {
		// 429
		loggers.Alerttransaction(c.Path(), "folderName", Uniqueidalertaccountid, "You have exceeded the service rate limit. Maximum allowed: 5 transactions per day", "(Rate Limit Exceeded - 429)", 0, "null", "null", "null", time.Now(), time.Now(), "null", "null", "null", "null", "null", "null", 0, "null", 0, "null", "null", "null", "null", 0, "null", time.Now(), "null")
		return errorhandling.Rate_Limit_Exceeded(c, "Rate limit exceeded")
	}

	if c.Method() != fiber.MethodPost {
		//405
		loggers.Alerttransaction(c.Path(), "folderName", Uniqueidalertaccountid, "Only POST method allowed", "(Method Not Allowed - 405)", 0, "null", "null", "null", time.Now(), time.Now(), "null", "null", "null", "null", "null", "null", 0, "null", 0, "null", "null", "null", "null", 0, "null", time.Now(), "null")
		return errorhandling.Method_Not_Allowed(c, "Method_Not_Allowed")
	}

	if err := c.BodyParser(&info); err != nil {
		//400
		loggers.Alerttransaction(c.Path(), "folderName", Uniqueidalertaccountid, "The request contains a bad payload", "(Bad Request - 400)", 0, "null", "null", "null", time.Now(), time.Now(), "null", "null", "null", "null", "null", "null", 0, "null", 0, "null", "null", "null", "null", 0, "null", time.Now(), "null")
		return errorhandling.Bad_Request(c, "The request contains a bad payload")
	}

	rows, err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.trace_alerts WHERE txnid = ?`, info.Transactionalertid).Rows()
	if err != nil {
		//400
		loggers.Alerttransaction(c.Path(), "folderName", Uniqueidalertaccountid, "The request contains a bad payload", "(Bad Request - 400)", 0, "null", "null", "null", time.Now(), time.Now(), "null", "null", "null", "null", "null", "null", 0, "null", 0, "null", "null", "null", "null", 0, "null", time.Now(), "null")
		return errorhandling.Bad_Request(c, "The request contains a bad payload")
	}
	defer rows.Close()

	if rows.Next() {
		if scanErr := database.DBConn.ScanRows(rows, &resp); scanErr != nil {
			//500
			loggers.Alerttransaction(c.Path(), "folderName", Uniqueidalertaccountid, "Error scanning database result", "(DatabaseError - 500)", 0, "null", "null", "null", time.Now(), time.Now(), "null", "null", "null", "null", "null", "null", 0, "null", 0, "null", "null", "null", "null", 0, "null", time.Now(), "null")
			return errorhandling.Internal_Server_Error(c, "Error scanning database result")
		}
	} else {

		//400
		return errorhandling.Bad_Request(c, "No data found for the given account ID")
	}

	return c.JSON(resp)

}

func Iftgeneratealerttransactionid(length int) string {
	rand.Seed(time.Now().UnixNano())

	hexDigits := "0123456789abcdef"
	result := make([]byte, length)
	for i := range result {
		result[i] = hexDigits[rand.Intn(len(hexDigits))]
	}
	return string(result)
}

//200/
//400/
//404/
//405/
//429/

//401
//403
