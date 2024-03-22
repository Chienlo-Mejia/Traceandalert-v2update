package healthchecks

import (
	"tracealert/pkg/models"
	errorhandling "tracealert/pkg/models/errorHandling"
	"tracealert/pkg/utils/go-utils/database"

	"github.com/gofiber/fiber/v2"
)

var requestCounts_matches int

const allowedRates_matches = 5

func checkRateLimits_matches() bool {
	requestCounts_matches++
	if requestCounts_matches > allowedRates_matches {
		requestCounts_matches = 0
		return true
	}
	return false
}

func Transactionidmatches(c *fiber.Ctx) error {
	// Parse the request body
	info := &models.TransacIDmatches{}
	resp := &models.TransactionMatches{}

	if checkRateLimits_matches() {
		// 429
		return errorhandling.Rate_Limit_Exceeded(c, "Rate limit exceeded")
	}

	if c.Method() != fiber.MethodPost {
		// 405
		return errorhandling.Method_Not_Allowed(c, "Method_Not_Allowed")
	}

	if !c.Is("json") {
		// 415
		return errorhandling.Unsupported_Media_Type(c, "Unsupported media type")
	}

	if err := c.BodyParser(&info); err != nil {
		//400
		return errorhandling.Bad_Request(c, "The request contains a bad payload")
	}
	accountArray := &models.RequestBodyArraymatches{}
	// Parse the request body
	if parsErr := c.BodyParser(accountArray); parsErr != nil {
		// 422
		return errorhandling.Unprocessable_Entity(c, "Unprocessable_Entity")
	}

	rows, err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.trace_alerts WHERE sourceid = ? AND service =? AND destid =? AND txntime =? AND value =?`, info.Sourceid, info.Service, info.Beneficiaryid, info.Transactiontime, info.Amount).Rows()
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

	alerts := fiber.Map{
		"matchID":       resp.Id,
		"sourceID":      resp.Sourceid,
		"beneficiaryID": info.Beneficiaryid,
		"refTime":       resp.Reftime,
		"matches": []fiber.Map{
			{

				"txnID":  resp.Txnid,
				"time":   resp.Time,
				"amount": info.Amount,
			},
		},
	}

	return c.JSON(alerts)

}

//429/
//400/
//405/
//415/
//422/
//404/
//200/

//401
//403
