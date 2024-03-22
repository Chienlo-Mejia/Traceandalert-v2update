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

var requestCountsalert_transaction int

const allowedRatesalert_transaction = 5

func checkRateLimitsalert_transaction() bool {
	requestCountsalert_transaction++
	if requestCountsalert_transaction > allowedRatesalert_transaction {
		requestCountsalert_transaction = 0
		return true
	}
	return false
}
func Alerttransaction(c *fiber.Ctx) error {
	var userRequest models.Transaction_Body
	Uniqueidalertaccount := Iftgeneratealerttransaction(32)

	if checkRateLimitsalert_transaction() {
		//429
		loggers.Alerttransaction(c.Path(), "folderName", Uniqueidalertaccount, "You have exceeded the service rate limit. Maximum allowed: 5 transactions per day", "(Rate Limit Exceeded - 429)", 0, "null", "null", "null", time.Now(), time.Now(), "null", "null", "null", "null", "null", "null", 0, "null", 0, "null", "null", "null", "null", 0, "null", time.Now(), "null")
		return errorhandling.Rate_Limit_Exceeded(c, "Rate limit exceeded ")
	}

	if err := c.BodyParser(&userRequest); err != nil {
		//400
		loggers.Alerttransaction(c.Path(), "folderName", Uniqueidalertaccount, "The request contains a bad payload", "(Bad Request - 400)", 0, "null", "null", "null", time.Now(), time.Now(), "null", "null", "null", "null", "null", "null", 0, "null", 0, "null", "null", "null", "null", 0, "null", time.Now(), "null")
		return errorhandling.Bad_Request(c, "The request contains a bad payload")
	}

	var transactions []models.Transaction_Response

	Query := `SELECT * FROM trace_alerts.trace_alerts WHERE 1 = 1`
	CountQuery := "SELECT COUNT(*) FROM trace_alerts.trace_alerts WHERE 1=1"
	CountFilteredQuery := "SELECT COUNT(*) FROM trace_alerts.trace_alerts WHERE 1=1"

	if userRequest.Since != "" {
		Query += " AND DATE_TRUNC('day', decisiondate) = $1"
		CountFilteredQuery += " AND DATE_TRUNC('day', decisiondate) = $1"
	}

	if userRequest.Filter != "" {
		if userRequest.Filter == "%3D%3D" {
			Query += " AND (networkid = '1234567890' OR id ='1234567890' OR txnid='1234567890' OR networkalertid = '1234567890' OR sourceid ='12345678' OR destid ='87654321' OR sourcebankid = '123456' OR destbankid = '654321' OR parentalertid = '12345678' OR remitinfo = 'Reference' OR currency ='USD' OR service = 'IPS') OR EXTRACT(HOUR FROM txntime) = 12 AND EXTRACT(MINUTE FROM txntime) = 0 OR EXTRACT(HOUR FROM time) = 12 AND EXTRACT(MINUTE FROM time) = 0 OR tracetype = 'FRAUD' OR mostrecentfeedback = 'ACTIONED_MULE' OR DATE(decisiondate) = '2020-01-01' "
			CountFilteredQuery += " AND (networkid = '1234567890' OR id ='1234567890' OR txnid='1234567890' OR networkalertid = '1234567890' OR sourceid ='12345678' OR destid ='87654321' OR sourcebankid = '123456' OR destbankid = '654321' OR parentalertid = '12345678' OR remitinfo = 'Reference' OR currency ='USD' OR service = 'IPS') OR EXTRACT(HOUR FROM txntime) = 12 AND EXTRACT(MINUTE FROM txntime) = 0 OR EXTRACT(HOUR FROM time) = 12 AND EXTRACT(MINUTE FROM time) = 0 OR tracetype = 'FRAUD' OR mostrecentfeedback = 'ACTIONED_MULE' OR DATE(decisiondate) = '2020-01-01'"
		} else if userRequest.Filter == "%3E%3D" {
			Query += " AND (mulescore >= 0.7)"
			CountFilteredQuery += " AND (mulescore >= 0.7)"
		} else if userRequest.Filter == "%3C%3D" {
			Query += " AND (value <= 28070.82 OR value <= 11228.33)"
			CountFilteredQuery += " AND (value <= 28070.82 OR value <= 11228.33)"
		} else if userRequest.Filter == "%3E" {
			Query += " AND (mulescore > 0.8)"
			CountFilteredQuery += " AND (mulescore > 0.8)"
		} else if userRequest.Filter == "%3C" {
			Query += " AND (value < 56141.63 OR generation < 10 OR EXTRACT(EPOCH FROM dwelltime) < 300)"
			CountFilteredQuery += " AND (value < 56141.63 OR generation < 10 OR EXTRACT(EPOCH FROM dwelltime) < 300)"
		} else {
			Query += " AND (" + userRequest.Filter + ")"
			CountFilteredQuery += " AND (" + userRequest.Filter + ")"
		}
	}

	var filteredCount int64

	var count int64
	resultCount := database.DBConn.Raw(CountQuery).Count(&count)
	if resultCount.Error != nil {
		//400
		loggers.Alerttransaction(c.Path(), "folderName", Uniqueidalertaccount, "The request contains a bad payload", "(Bad Request - 400)", 0, "null", "null", "null", time.Now(), time.Now(), "null", "null", "null", "null", "null", "null", 0, "null", 0, "null", "null", "null", "null", 0, "null", time.Now(), "null")
		return errorhandling.Bad_Request(c, "Error counting transactions")
	}

	resultFilteredCount := database.DBConn.Raw(CountFilteredQuery, userRequest.Since).Count(&filteredCount)
	if resultFilteredCount.Error != nil {
		//400
		loggers.Alerttransaction(c.Path(), "folderName", Uniqueidalertaccount, "The request contains a bad payload", "(Bad Request - 400)", 0, "null", "null", "null", time.Now(), time.Now(), "null", "null", "null", "null", "null", "null", 0, "null", 0, "null", "null", "null", "null", 0, "null", time.Now(), "null")
		return errorhandling.Bad_Request(c, "Error counting filtered transactions")
	}

	if userRequest.Limit > 0 {
		if int64(userRequest.Limit) > filteredCount {
			//400
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Limit exceeds the total number of filtered transactions",
			})
		}
		Query += fmt.Sprintf(" LIMIT %d", userRequest.Limit)
	}

	err := database.DBConn.Debug().Raw(Query, userRequest.Since).Scan(&transactions).Error

	if err != nil {
		//400
		loggers.Alerttransaction(c.Path(), "folderName", Uniqueidalertaccount, "The request contains a bad payload", "(Bad Request - 400)", 0, "null", "null", "null", time.Now(), time.Now(), "null", "null", "null", "null", "null", "null", 0, "null", 0, "null", "null", "null", "null", 0, "null", time.Now(), "null")
		return errorhandling.Bad_Request(c, "The request contains a bad payload")
	}

	if c.Method() != fiber.MethodPost {
		//405
		loggers.Alerttransaction(c.Path(), "folderName", Uniqueidalertaccount, "Only POST method allowed", "(Method Not Allowed - 405)", 0, "null", "null", "null", time.Now(), time.Now(), "null", "null", "null", "null", "null", "null", 0, "null", 0, "null", "null", "null", "null", 0, "null", time.Now(), "null")
		return errorhandling.Method_Not_Allowed(c, "Method_Not_Allowed")
	}

	if len(transactions) == 0 {
		//400
		loggers.Alerttransaction(c.Path(), "folderName", Uniqueidalertaccount, "The request contains a bad payload", "(Bad Request - 400)", 0, "null", "null", "null", time.Now(), time.Now(), "null", "null", "null", "null", "null", "null", 0, "null", 0, "null", "null", "null", "null", 0, "null", time.Now(), "null")
		return errorhandling.Url_Not_Found(c, "No data found for the specified date")

	}

	var errorresp errors.Errorresp

	// --------------------------- logs ---------------------------

	userLimit := userRequest.Limit
	if userLimit < 0 {
		return errorhandling.Bad_Request(c, "Invalid limit")
	}

	for i := 0; i < userLimit && i < len(transactions); i++ {

		ID := transactions[i].ID
		Errors := errorresp.Errorresponse
		Count := count
		Txnid := transactions[i].Txnid
		Networkalertid := transactions[i].Networkalertid
		Networkid := transactions[i].Networkid
		Time := transactions[i].Time
		Txntime := transactions[i].Txntime
		Sourceid := transactions[i].Sourceid
		Destid := transactions[i].Destid
		Sourcebankid := transactions[i].Sourcebankid
		Sourcebankname := transactions[i].Sourcebankname
		Destbankid := transactions[i].Destbankid
		Destbankname := transactions[i].Destbankname
		Value := transactions[i].Value
		Remitinfo := transactions[i].Remitinfo
		Generation := transactions[i].Generation
		Currency := transactions[i].Currency
		Service := transactions[i].Service
		Dwelltime := transactions[i].Dwelltime
		Tracetype := transactions[i].Tracetype
		Mulescore := transactions[i].Mulescore
		Parentalertid := transactions[i].Parentalertid
		Decisiondate := transactions[i].Decisiondate
		Mostrecentfeedback := transactions[i].Mostrecentfeedback

		message := fmt.Sprintf(" Success %s ", Errors)
		loggers.Alerttransaction(c.Path(), "folderName", Uniqueidalertaccount, ID, message, Count, Txnid, Networkalertid, Networkid, Time, Txntime, Sourceid, Destid, Sourcebankid, Sourcebankname, Destbankid, Destbankname, Value, Remitinfo, Generation, Currency, Service, Dwelltime, Tracetype, Mulescore, Parentalertid, Decisiondate, Mostrecentfeedback)
	}

	// ---------------------------

	// Referenceid := Referenceid(ID)
	// Generateid := Generateid(ID)
	type Alert struct {
		Alerts       fiber.Map                     `json:"alerts"`
		Transactions []models.Transaction_Response `json:"transactions"`
	}

	response := Alert{
		Alerts: fiber.Map{
			"totalRecords":            count,
			"displayedRecords":        userRequest.Limit,
			"nextPaginationToken":     transactions[0].Nextpaginationtoken,
			"previousPaginationToken": transactions[0].Previouspaginationtoken,
			// "referenceID":             Referenceid,
			// "generateID":              Generateid,
		},

		Transactions: transactions,
	}

	return c.JSON(response)
}

func Iftgeneratealerttransaction(length int) string {
	rand.Seed(time.Now().UnixNano())

	hexDigits := "0123456789abcdef"
	result := make([]byte, length)
	for i := range result {
		result[i] = hexDigits[rand.Intn(len(hexDigits))]
	}
	return string(result)
}

func GenerateIBFTReferenceID(instructionID string) string {
	var ipsReference string
	for ctr := 1; ctr <= 6; ctr++ {
		ipsReference = instructionID[len(instructionID)-ctr:]
	}
	return fmt.Sprintf("%v", ipsReference)
}
