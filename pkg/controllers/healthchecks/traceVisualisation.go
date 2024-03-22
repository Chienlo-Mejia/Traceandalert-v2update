package healthchecks

import (
	"fmt"
	"math/rand"
	"time"
	"tracealert/middleware/loggers"
	"tracealert/pkg/models"
	errorhandling "tracealert/pkg/models/errorHandling"
	"tracealert/pkg/utils/go-utils/database"

	"github.com/gofiber/fiber/v2"
)

var requestCounts_tracevisual int

const allowedRates_tracevisual = 5

func checkRateLimits_tracevisual() bool {
	requestCounts_tracevisual++
	if requestCounts_tracevisual > allowedRates_tracevisual {
		requestCounts_tracevisual = 0
		return true
	}
	return false
}
func NetworkAlertID(c *fiber.Ctx) error {
	var userRequest models.Tracevisualisationsalert
	Uniqueidvisualisation := Iftgeneratevisualisation(32)

	// Check rate limits
	if checkRateLimits_tracevisual() {
		loggers.Tracevisuallogs(c.Path(), "folderName", Uniqueidvisualisation, userRequest.Networkalertid, "You have exceeded the service rate limit. Maximum allowed: 5 transactions per day", "(Rate Limit Exceeded - 429)")
		return errorhandling.Rate_Limit_Exceeded(c, "Rate limit exceeded ")
	}

	if c.Method() != fiber.MethodPost {
		loggers.Tracevisuallogs(c.Path(), "folderName", Uniqueidvisualisation, userRequest.Networkalertid, "Only POST method allowed ", "(Method Not Allowed - 405)")
		return errorhandling.Method_Not_Allowed(c, "Method not allowed")
	}

	accountArray := &models.RequestBodyArray{}
	// Parse the request body
	if parsErr := c.BodyParser(accountArray); parsErr != nil {
		loggers.Tracevisuallogs(c.Path(), "folderName", Uniqueidvisualisation, userRequest.Networkalertid, "Unprocessable entity", "(Unprocessable Entity - 422)")
		return errorhandling.Unprocessable_Entity(c, "Unprocessable_Entity")
	}

	// Parse the request body into userRequest
	if err := c.BodyParser(&userRequest); err != nil {
		loggers.Tracevisuallogs(c.Path(), "folderName", Uniqueidvisualisation, userRequest.Networkalertid, "The request body is expecting an array ", "(Bad Request - 400)")
		return errorhandling.Bad_Request(c, "Bad_Request")
	}

	// Check if the network alert ID already exists in the database
	exists, err := checkNetworkAlertIDExists(userRequest.Networkalertid)
	if err != nil {
		loggers.Tracevisuallogs(c.Path(), "folderName", Uniqueidvisualisation, userRequest.Networkalertid, "Database error", "(Internal Server Error - 500)")
		return errorhandling.Internal_Server_Error(c, "Database error")
	}
	if exists {
		loggers.Tracevisuallogs(c.Path(), "folderName", Uniqueidvisualisation, userRequest.Networkalertid, "Network alert ID already exists", "(Conflict - 409)")
		return errorhandling.Conflict(c, "Network alert ID already exists")
	}

	// Build SQL query
	query := `SELECT * FROM trace_alerts.trace_visualisation WHERE 1 = 1`

	// Use parameters to avoid SQL injection
	var params []interface{}

	if userRequest.Networkalertid != "" && userRequest.Networkalertid != "0" {
		query += " AND networkalertid = ?"
		params = append(params, userRequest.Networkalertid)
	}

	if userRequest.Format != "" && userRequest.Format != "0" {
		query += " AND format = ?"
		params = append(params, userRequest.Format)
	}

	if userRequest.Type != "" && userRequest.Type != "0" {
		query += " AND type = ?"
		params = append(params, userRequest.Type)
	}

	if userRequest.Colourmode != "" && userRequest.Colourmode != "0" {
		query += " AND colourmode = ?"
		params = append(params, userRequest.Colourmode)
	}

	// Perform database query
	var feedbackList []models.Traceresponse
	result := database.DBConn.Raw(query+" ORDER BY id", params...).Scan(&feedbackList)

	if result.Error != nil {
		loggers.Tracevisuallogs(c.Path(), "folderName", Uniqueidvisualisation, userRequest.Networkalertid, "Database error", "(Internal Server Error - 500)")
		return errorhandling.Internal_Server_Error(c, "Database error")
	}

	if len(feedbackList) == 0 {
		loggers.Tracevisuallogs(c.Path(), "folderName", Uniqueidvisualisation, userRequest.Networkalertid, "No matching data found", "(Not Found - 404)")
		return errorhandling.Bad_Request(c, "The request contains a bad payload")
	}

	// Log success
	logMessage := fmt.Sprintf(" %s ", feedbackList[0].Traceid)
	message := "Success"
	loggers.Tracevisuallogs(c.Path(), "folderName", Uniqueidvisualisation, logMessage, message, userRequest.Networkalertid)

	return c.JSON(feedbackList)
}

func checkNetworkAlertIDExists(Networkalertid string) (bool, error) {
	var count int
	err := database.DBConn.Raw("SELECT COUNT(*) FROM trace_alerts.trace_visualisation WHERE networkalertid = ?", Networkalertid).Scan(&count).Error
	if err != nil {
		return false, err
	}
	return count > 1, nil
}

func Iftgeneratevisualisation(length int) string {
	rand.Seed(time.Now().UnixNano())

	hexDigits := "0123456789abcdef"
	result := make([]byte, length)
	for i := range result {
		result[i] = hexDigits[rand.Intn(len(hexDigits))]
	}
	return string(result)
}

// func NetworkAlertID(c *fiber.Ctx) error {
// 	var userRequest models.Tracevisualisationsalert

// 	if checkRateLimits_tracevisual() {
// 		// 429
// 		loggers.Tracevisuallogs(c.Path(), "folderName", userRequest.Networkalertid, "You have exceeded the service rate limit. Maximum allowed: ${rate_limit.output} TPS", "(Rate Limit Exceeded - 409)")
// 		return errorhandling.Rate_Limit_Exceeded(c, "Rate limit exceeded ")

// 	}

// 	if c.Method() != fiber.MethodPost {
// 		// 405
// 		loggers.Tracevisuallogs(c.Path(), "folderName", userRequest.Networkalertid, "Only POST method allowed ", "(Method Not Allowed - 405)")
// 		return errorhandling.Method_Not_Allowed(c, "Method not allowed")
// 	}

// 	accountArray := &models.RequestBodyArray{}
// 	// Parse the request body
// 	if parsErr := c.BodyParser(accountArray); parsErr != nil {
// 		// 422
// 		loggers.Tracevisuallogs(c.Path(), "folderName", userRequest.Networkalertid, "Unprocessable entity", "(Unprocessable Entity - 422)")
// 		return errorhandling.Unprocessable_Entity(c, "Unprocessable_Entity")
// 	}

// 	if err := c.BodyParser(&userRequest); err != nil {
// 		// 400
// 		loggers.Tracevisuallogs(c.Path(), "folderName", userRequest.Networkalertid, "The request body is expecting an array ", "(Bad Request - 400)")
// 		return errorhandling.Bad_Request(c, "Bad_Request")
// 	}

// 	query := ` SELECT * FROM trace_alerts.trace_visualisation WHERE 1 = 1`

// 	if userRequest.Networkalertid != "" && userRequest.Networkalertid != "0" {
// 		query += " AND networkalertid = ?"
// 	}

// 	if userRequest.Format != "" && userRequest.Format != "0" {
// 		query += " AND format = ?"
// 	}

// 	if userRequest.Type != "" && userRequest.Type != "0" {
// 		query += " AND type = ?"
// 	}

// 	if userRequest.Colourmode != "" && userRequest.Colourmode != "0" {
// 		query += " AND colourmode = ?"
// 	}

// 	var feedbackList []models.Traceresponse
// 	//error response
// 	var errorresp errors.Errorresp

// 	result := database.DBConn.Raw(query+" ORDER BY id", userRequest.Networkalertid, userRequest.Format, userRequest.Type, userRequest.Colourmode).Scan(&feedbackList)
// 	if result.Error != nil {
// 		// 409
// 		loggers.Tracevisuallogs(c.Path(), "folderName", userRequest.Networkalertid, " Alert ID does not match the specified entity", "(Conflict - 409)")
// 		return errorhandling.Conflict(c, "Alert ID does not match the specified entity")

// 	}

// 	if userRequest.Status {
// 		// 409
// 		loggers.Tracevisuallogs(c.Path(), "folderName", userRequest.Networkalertid, " Invalid customer for third party", "(Permision Denied - 403)")
// 		return errorhandling.Permision_Denied(c, "Invalid customer for third party")
// 	}

// 	if len(feedbackList) == 0 {
// 		loggers.Tracevisuallogs(c.Path(), "folderName", userRequest.Networkalertid, "The request body is expecting an array ", "(Bad Request - 400)")
// 		return errorhandling.Bad_Request(c, "The request contains a bad payload")
// 	}
// 	//------------------- logs -------------------
// 	Traceid := feedbackList[0].Traceid
// 	Errors := errorresp.Errorresponse
// 	Networkalertid := userRequest.Networkalertid

// 	logMessage := fmt.Sprintf("%s ", Traceid)
// 	message := fmt.Sprintf(" Success %s ", Errors)
// 	loggers.Tracevisuallogs(c.Path(), "folderName", logMessage, message, Networkalertid)

// 	return c.JSON(feedbackList)

// }

//429/
//400/
//405/
//500
//403/
//404/
//200/

//401/
