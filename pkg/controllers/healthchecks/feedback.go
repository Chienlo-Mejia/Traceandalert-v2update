package healthchecks

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
	"tracealert/middleware/loggers"
	"tracealert/pkg/models"
	errorhandling "tracealert/pkg/models/errorHandling"
	"tracealert/pkg/models/errors"
	"tracealert/pkg/utils/go-utils/database"

	"github.com/gofiber/fiber/v2"
)

var (
	mu                 sync.Mutex
	requestCountsPerID = make(map[string]map[string]int)
)

const maxTransactionsPerDay = 5

func checkRateLimits_feedback(ID string) bool {
	mu.Lock()
	defer mu.Unlock()
	today := time.Now().Format("2006-01-02")
	if _, ok := requestCountsPerID[today]; !ok {
		// If not, initialize it
		requestCountsPerID[today] = make(map[string]int)
	}

	// Increment the request count for the specific ID for today
	requestCountsPerID[today][ID]++

	// Return true if the request count exceeds the limit, else false
	return requestCountsPerID[today][ID] > maxTransactionsPerDay
}

// Function to reset requestCountsPerID daily at midnight
func ResetRequestCountsDaily() {
	for {
		// Get the current time
		now := time.Now()

		// Calculate the duration until the next midnight
		nextMidnight := now.Truncate(24 * time.Hour).Add(24 * time.Hour)
		timeToMidnight := nextMidnight.Sub(now)

		// Sleep until the next midnight
		time.Sleep(timeToMidnight)

		// Reset counts for the next day
		mu.Lock()
		delete(requestCountsPerID, now.Format("2006-01-02"))
		mu.Unlock()
	}
}

func Feedback(c *fiber.Ctx) error {
	var userRequest models.RequestBody
	Uniqueidfeedback := Iftgeneratealertfeedback(32)

	if c.Method() != fiber.MethodPost {
		// 405
		loggers.Feedbacklogs(c.Path(), "folderName", Uniqueidfeedback, userRequest.Alertid, "Only POST method allowed", "(Method Not Allowed - 405)")
		return errorhandling.Method_Not_Allowed(c, "Method_Not_Allowed")
	}

	if !c.Is("json") {
		// 415
		loggers.Feedbacklogs(c.Path(), "folderName", Uniqueidfeedback, userRequest.Alertid, "The request media type 'application/x-www-form-urlencoded' is not supported by this resource", "(Unsupported Media Type - 415)")
		return errorhandling.Unsupported_Media_Type(c, "Unsupported media type")
	}

	// Parse the request body
	if err := c.BodyParser(&userRequest); err != nil {
		// 400
		loggers.Feedbacklogs(c.Path(), "folderName", Uniqueidfeedback, userRequest.Alertid, "The request body is expecting an array", "(Bad Request - 400)")
		return errorhandling.Bad_Request(c, "The request body is expecting an array")
	}

	// Check rate limits
	if checkRateLimits_feedback(userRequest.Alertid) {
		// 429
		loggers.Feedbacklogs(c.Path(), "folderName", Uniqueidfeedback, userRequest.Alertid, "You have exceeded the service rate limit. Maximum allowed: 5 transactions per day", "(Rate Limit Exceeded - 429)")
		return errorhandling.Rate_Limit_Exceeded(c, "Rate limit exceeded")
	}

	accountArray := &models.RequestBodyArray{}
	// Parse the request body
	if parsErr := c.BodyParser(accountArray); parsErr != nil {
		// 422
		loggers.Feedbacklogs(c.Path(), "folderName", Uniqueidfeedback, userRequest.Alertid, "Expects a single JSON object and not an array", "(Unprocessable Entity - 422)")
		return errorhandling.Unprocessable_Entity(c, "Unprocessable_Entity")
	}

	if err := c.BodyParser(&userRequest); err != nil {
		// 400
		loggers.Feedbacklogs(c.Path(), "folderName", Uniqueidfeedback, userRequest.Alertid, "The request body is expecting an array", "(Bad Request - 400)")
		return errorhandling.Bad_Request(c, "The request body is expecting an array")
	}

	// query := ` SELECT id, alertid, alerttype, entityid, decisiondate, feedback, status, feedbackid FROM public.feedback WHERE 1 = 1`
	query := ` SELECT * FROM trace_alerts.feedback WHERE 1 = 1`

	if userRequest.Alertid != "" && userRequest.Alertid != "0" {
		query += " AND alertid = ?"
	}

	if userRequest.Alerttype != "" && userRequest.Alerttype != "0" {
		query += " AND alerttype = ?"
	}

	if userRequest.Entityid != "" && userRequest.Entityid != "0" {
		query += " AND entityid = ?"
	}

	if userRequest.Feedback != "" && userRequest.Feedback != "0" {
		query += " AND feedback = ?"
	}

	//check if exist
	exists, err := checkalertidexist(userRequest.Alertid)
	if err != nil {
		loggers.Feedbacklogs(c.Path(), "folderName", Uniqueidfeedback, userRequest.Alertid, "Database error", "(Internal Server Error - 500)")
		return errorhandling.Internal_Server_Error(c, "Database error")
	}
	if exists {
		loggers.Feedbacklogs(c.Path(), "folderName", Uniqueidfeedback, userRequest.Alertid, "Network alert ID already exists", "(Conflict - 409)")
		return errorhandling.Conflict(c, "Alert ID already exists")
	}
	var feedbackList []models.Response

	var errorresp errors.Errorresp

	result := database.DBConn.Raw(query+" ORDER BY id", userRequest.Alertid, userRequest.Alerttype, userRequest.Entityid, userRequest.Feedback).Scan(&feedbackList)
	if result.Error != nil {
		// 409
		loggers.Feedbacklogs(c.Path(), "folderName", Uniqueidfeedback, userRequest.Alertid, " Alert ID does not match the specified entity", "(Conflict - 409)")
		return errorhandling.Conflict(c, "Alert ID does not match the specified entity")
	}

	if len(feedbackList) == 0 {
		// 404
		loggers.Feedbacklogs(c.Path(), "folderName", Uniqueidfeedback, userRequest.Alertid, " The request contains a bad payload", "(NOT_FOUND - 400)")
		return errorhandling.Url_Not_Found(c, "URL_NOT_FOUND")
	}

	if userRequest.Status {
		// 403
		loggers.Feedbacklogs(c.Path(), "folderName", Uniqueidfeedback, userRequest.Alertid, " Invalid customer for third party", "(Permission Denied - 403)")
		return errorhandling.Permision_Denied(c, "Permission Denied")
	}

	//------------------- logs -------------------

	Feedbackid := feedbackList[0].Feedbackid
	Errors := errorresp.Errorresponse
	Alertid := userRequest.Alertid

	logMessage := fmt.Sprintf("%s ", Feedbackid)
	message := fmt.Sprintf(" Success %s ", Errors)
	loggers.Feedbacklogs(c.Path(), "folderName", Uniqueidfeedback, logMessage, message, Alertid)

	return c.JSON(fiber.Map{
		"feedbackID": Feedbackid,
	})
}

func checkalertidexist(Alertid string) (bool, error) {
	var count int
	err := database.DBConn.Raw("SELECT COUNT(*) FROM trace_alerts.feedback WHERE alertid = ?", Alertid).Scan(&count).Error
	if err != nil {
		return false, err
	}
	return count > 1, nil
}

func Iftgeneratealertfeedback(length int) string {
	rand.Seed(time.Now().UnixNano())

	hexDigits := "0123456789abcdef"
	result := make([]byte, length)
	for i := range result {
		result[i] = hexDigits[rand.Intn(len(hexDigits))]
	}
	return string(result)
}

//400
//403
//404
//200
//405
//409
//429
//422
//415

//401
