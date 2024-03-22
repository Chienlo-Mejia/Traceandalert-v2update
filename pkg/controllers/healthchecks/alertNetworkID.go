package healthchecks

import (
	"tracealert/pkg/models"
	errorhandling "tracealert/pkg/models/errorHandling"
	"tracealert/pkg/utils/go-utils/database"

	"github.com/gofiber/fiber/v2"
)

var requestCounts_alertnetwork_ID int

const allowedRates_alertnetwork_ID = 5

func checkRateLimits_alertnetwork_ID() bool {
	requestCounts_alertnetwork_ID++
	if requestCounts_alertnetwork_ID > allowedRates_alertnetwork_ID {
		requestCounts_alertnetwork_ID = 0
		return true
	}
	return false
}
func GetNetworkInfo(c *fiber.Ctx) error {
	// Parse the request body
	info := &models.TransacAlertRequest{}

	if checkRateLimits_alertnetwork_ID() {
		//429
		return errorhandling.Rate_Limit_Exceeded(c, "You have exceeded the service rate limit. Maximum allowed: 5 transactions per day")
	}

	if c.Method() != fiber.MethodPost {
		// 405
		return errorhandling.Method_Not_Allowed(c, "Method_Not_Allowed")
	}

	if err := c.BodyParser(&info); err != nil {
		//400
		return errorhandling.Bad_Request(c, "The request contains a bad payload")
	}

	var resp1 models.TransactionAlert
	if err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.trace_alerts WHERE id = ?`, info.Transactionalertid).Scan(&resp1).Error; err != nil {
		//400
		return errorhandling.Bad_Request(c, "Error fetching data from trace_alerts table")
	}

	var resp2 models.Alertnetwork
	if err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.trace_account_alerts WHERE id = ?`, info.Transactionalertid).Scan(&resp2).Error; err != nil {
		//400
		return errorhandling.Bad_Request(c, "Error fetching data from trace_account_alerts table")
	}

	var resp3 models.NetworkResponse
	if err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.trace_vizurl WHERE parentalertid = ?`, info.Transactionalertid).Scan(&resp3).Error; err != nil {
		//400
		return errorhandling.Bad_Request(c, "Error fetching data from trace_vizurl table")
	}

	finalResp := struct {
		TraceAlerts        models.TransactionAlert
		TraceAccountAlerts models.Alertnetwork
		TraceVizurl        models.NetworkResponse
	}{
		TraceAlerts:        resp1,
		TraceAccountAlerts: resp2,
		TraceVizurl:        resp3,
	}

	return c.JSON(finalResp)
}

//200/
//400/
//404/
//405/
//429/

//401
//403
