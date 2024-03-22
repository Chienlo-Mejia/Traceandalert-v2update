package healthchecks

import (
	"tracealert/pkg/models"
	errorhandling "tracealert/pkg/models/errorHandling"
	"tracealert/pkg/utils/go-utils/database"

	"github.com/gofiber/fiber/v2"
)

var requestCounts_alertaccount_ID int

const allowedRates_alertaccount_ID = 5

func checkRateLimits_alertaccount_ID() bool {
	requestCounts_alertaccount_ID++
	if requestCounts_alertaccount_ID > allowedRates_alertaccount_ID {
		requestCounts_alertaccount_ID = 0
		return true
	}
	return false
}
func GetAccInfo(c *fiber.Ctx) error {
	// Parse the request body
	info := &models.AccountAlertRequest1{}
	resp := models.ResponseInfo1{}

	if checkRateLimits_alertaccount_ID() {
		// 429
		return errorhandling.Rate_Limit_Exceeded(c, "Rate limit exceeded ")
	}

	if err := c.BodyParser(&info); err != nil {
		//400
		return errorhandling.Bad_Request(c, "The request contains a bad payload")
	}

	if c.Method() != fiber.MethodPost {
		// 405
		return errorhandling.Method_Not_Allowed(c, "Method_Not_Allowed")
	}

	rows, err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.trace_account_alerts WHERE accountid = ?`, info.Accountid).Rows()
	if err != nil {
		//400
		return errorhandling.Bad_Request(c, "The request contains a bad payload")
	}
	defer rows.Close()

	if rows.Next() {

		if scanErr := database.DBConn.ScanRows(rows, &resp); scanErr != nil {
			//500
			return errorhandling.Internal_Server_Error(c, "Error scanning database result")
		}
	} else {

		//400
		return errorhandling.Bad_Request(c, "No data found for the given account ID")
	}

	return c.JSON(resp)
}

//200/
//400/
//404/
//405/
//429/

//401
//403
