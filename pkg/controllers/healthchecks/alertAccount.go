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

var requestCounts_alertaccount int

const allowedRates_alertaccount = 5

func checkRateLimits_alertaccount() bool {
	requestCounts_alertaccount++
	if requestCounts_alertaccount > allowedRates_alertaccount {
		requestCounts_alertaccount = 0
		return true
	}
	return false
}
func Alertsaccount(c *fiber.Ctx) error {
	var userRequest models.RequestBodyalert

	Uniqueidalertaccount := Iftgeneratealertaccount(32)

	if checkRateLimits_alertaccount() {
		//429
		loggers.Alertaccount(c.Path(), "folderName", Uniqueidalertaccount, "You have exceeded the service rate limit. Maximum allowed: 5 transactions per day", "(Rate Limit Exceeded - 429)", "null", "null", "null", "null", "null", time.Now(), "null", 0, 0, false, 0, 0, 0, time.Now(), time.Now(), false, "null", 0, 0, 0, "null", 0, 0, 0, 0, []int{0}, "null", "null", time.Now(), "null", "null")
		return errorhandling.Rate_Limit_Exceeded(c, "Rate limit exceeded ")
	}

	if err := c.BodyParser(&userRequest); err != nil {
		//400
		loggers.Alertaccount(c.Path(), "folderName", Uniqueidalertaccount, " The request body is expecting an array", "(Bad Request - 400)", "null", "null", "null", "null", "null", time.Now(), "null", 0, 0, false, 0, 0, 0, time.Now(), time.Now(), false, "null", 0, 0, 0, "null", 0, 0, 0, 0, []int{0}, "null", "null", time.Now(), "null", "null")
		return errorhandling.Bad_Request(c, "The request contains a bad payload")
	}

	var transactions []models.Alert

	Query := `SELECT * FROM trace_alerts.trace_account_alerts WHERE 1 = 1`
	CountQuery := "SELECT COUNT(*) FROM trace_alerts.trace_account_alerts WHERE 1=1"
	CountFilteredQuery := "SELECT COUNT(*) FROM trace_alerts.trace_account_alerts WHERE 1=1"

	if userRequest.Since != "" {
		Query += " AND DATE_TRUNC('day', decisiondate) = $1"
		CountFilteredQuery += " AND DATE_TRUNC('day', decisiondate) = $1"
	}

	if userRequest.Filter != "" {
		if userRequest.Filter == "%3D%3D" {
			Query += " AND (id = '12345678' OR networkalertid = '1234567890' OR accountid = '12345678' OR networkid = '1234567890' OR parentalertid = '12345678' OR owningbankname = 'OwningBank' OR name = 'John Doe' OR EXTRACT(HOUR FROM time) = 12 AND EXTRACT(MINUTE FROM time) = 0 OR tracetype = 'FRAUD' OR endpointflag = 'true' OR DATE(firstappearance) = '2020-01-01' OR DATE(mostrecentappearance) = '2020-01-01' OR EXTRACT(HOUR FROM firsttransactiontime) = 10 AND EXTRACT(MINUTE FROM firsttransactiontime) = 0 OR EXTRACT(HOUR FROM mostrecenttransactiontime) = 12 AND EXTRACT(MINUTE FROM mostrecenttransactiontime) = 0 OR mostrecentfeedback = 'ACTIONED_MULE' OR DATE(decisiondate) = '2020-01-01') "
			CountFilteredQuery += " AND (id = '12345678' OR networkalertid = '1234567890' OR accountid = '12345678' OR networkid = '1234567890' OR parentalertid = '12345678' OR owningbankname = 'OwningBank' OR name = 'John Doe' OR EXTRACT(HOUR FROM time) = 12 AND EXTRACT(MINUTE FROM time) = 0 OR tracetype = 'FRAUD' OR endpointflag = 'true' OR DATE(firstappearance) = '2020-01-01' OR DATE(mostrecentappearance) = '2020-01-01' OR EXTRACT(HOUR FROM firsttransactiontime) = 10 AND EXTRACT(MINUTE FROM firsttransactiontime) = 0 OR EXTRACT(HOUR FROM mostrecenttransactiontime) = 12 AND EXTRACT(MINUTE FROM mostrecenttransactiontime) = 0 OR  mostrecentfeedback = 'ACTIONED_MULE' OR DATE(decisiondate) = '2020-01-01')"
		} else if userRequest.Filter == "%3E%3D" {
			Query += " AND (mulescore >= 0.7 OR mulescore <= 0.75)"
			CountFilteredQuery += " AND (mulescore >= 0.7 OR mulescore <= 0.75)"
		} else if userRequest.Filter == "%3C%3D" {
			Query += " AND (value <= 28070.82 OR value <= 11228.33 OR numnetworks < 3 OR numtracednetworks < 3)"
			CountFilteredQuery += " AND (value <= 28070.82 OR value <= 11228.33 OR numnetworks < 3 OR numtracednetworks < 3)"
		} else if userRequest.Filter == "%3E" {
			Query += " AND (mulescore > 0.8)"
			CountFilteredQuery += " AND (mulescore > 0.8)"
		} else if userRequest.Filter == "%3C" {
			Query += " AND (sourcetransactionvalue > 56141.63 OR totalsuspiciousvalueinbound > 56141.63 OR totalsuspiciousvalueoutbound > 56141.63 OR totalvalueinbound > 56141.63 OR totalvalueoutbound > 56141.63 OR generation < 10 OR EXTRACT(EPOCH FROM dwelltime) < 300 OR numinboundrelationships > 2 OR numoutboundrelationships > 3 OR numscheduledmandates < 1  )"
			CountFilteredQuery += " AND (sourcetransactionvalue > 56141.63 OR totalsuspiciousvalueinbound > 56141.63 OR totalsuspiciousvalueoutbound > 56141.63 OR totalvalueinbound > 56141.63 OR totalvalueoutbound > 56141.63 OR generation < 10 OR EXTRACT(EPOCH FROM dwelltime) < 300)"
		} else {
			Query += " AND (" + userRequest.Filter + ")"
			CountFilteredQuery += " AND (" + userRequest.Filter + ")"
		}
	}

	// var totalCount int64
	var filteredCount int64

	var count int64
	resultCount := database.DBConn.Raw(CountQuery).Count(&count)
	if resultCount.Error != nil {
		//400
		loggers.Alertaccount(c.Path(), "folderName", Uniqueidalertaccount, " Error counting transactions", "(Bad Request - 400)", "null", "null", "null", "null", "null", time.Now(), "null", 0, 0, false, 0, 0, 0, time.Now(), time.Now(), false, "null", 0, 0, 0, "null", 0, 0, 0, 0, []int{0}, "null", "null", time.Now(), "null", "null")
		return errorhandling.Bad_Request(c, "Error counting transactions")
	}

	resultFilteredCount := database.DBConn.Raw(CountFilteredQuery, userRequest.Since).Count(&filteredCount)
	if resultFilteredCount.Error != nil {
		//400
		loggers.Alertaccount(c.Path(), "folderName", Uniqueidalertaccount, "Error counting filtered transactions", "(Bad Request - 400)", "null", "null", "null", "null", "null", time.Now(), "null", 0, 0, false, 0, 0, 0, time.Now(), time.Now(), false, "null", 0, 0, 0, "null", 0, 0, 0, 0, []int{0}, "null", "null", time.Now(), "null", "null")
		return errorhandling.Bad_Request(c, "Error counting filtered transactions")
	}

	////limit count and display data

	if userRequest.Limit > 0 {
		if int64(userRequest.Limit) > filteredCount {
			//400
			loggers.Alertaccount(c.Path(), "folderName", Uniqueidalertaccount, "Limit exceeds the total number of filtered transactions", "(Bad Request - 400)", "null", "null", "null", "null", "null", time.Now(), "null", 0, 0, false, 0, 0, 0, time.Now(), time.Now(), false, "null", 0, 0, 0, "null", 0, 0, 0, 0, []int{0}, "null", "null", time.Now(), "null", "null")
			return errorhandling.Bad_Request(c, "Limit exceeds the total number of filtered transactions")
		}
		Query += fmt.Sprintf(" LIMIT %d", userRequest.Limit)
	}

	err := database.DBConn.Debug().Raw(Query, userRequest.Since).Scan(&transactions).Error

	if err != nil {
		//400
		loggers.Alertaccount(c.Path(), "folderName", Uniqueidalertaccount, "The request contains a bad payload", "(Bad Request - 400)", "null", "null", "null", "null", "null", time.Now(), "null", 0, 0, false, 0, 0, 0, time.Now(), time.Now(), false, "null", 0, 0, 0, "null", 0, 0, 0, 0, []int{0}, "null", "null", time.Now(), "null", "null")
		return errorhandling.Bad_Request(c, "The request contains a bad payload")
	}

	if c.Method() != fiber.MethodPost {
		//405
		loggers.Alertaccount(c.Path(), "folderName", Uniqueidalertaccount, "Only POST method allowed", "(Method Not Allowed - 405)", "null", "null", "null", "null", "null", time.Now(), "null", 0, 0, false, 0, 0, 0, time.Now(), time.Now(), false, "null", 0, 0, 0, "null", 0, 0, 0, 0, []int{0}, "null", "null", time.Now(), "null", "null")
		return errorhandling.Method_Not_Allowed(c, "Method_Not_Allowed")
	}

	if len(transactions) == 0 {
		//400
		loggers.Alertaccount(c.Path(), "folderName", Uniqueidalertaccount, " No data found for the specified date", "", "null", "null", "null", "null", "null", time.Now(), "null", 0, 0, false, 0, 0, 0, time.Now(), time.Now(), false, "null", 0, 0, 0, "null", 0, 0, 0, 0, []int{0}, "null", "null", time.Now(), "null", "null")
		return errorhandling.Url_Not_Found(c, "No data found for the specified date")

	}

	var errorresp errors.Errorresp

	//--------------------------- logs ---------------------------

	userLimit := userRequest.Limit
	if userLimit < 0 {
		return errorhandling.Bad_Request(c, "Invalid limit")
	}

	for i := 0; i < userLimit && i < len(transactions); i++ {

		ID := transactions[i].ID
		Errors := errorresp.Errorresponse
		Networkalertid := transactions[i].NetworkalertId
		Accountid := transactions[i].AccountId
		Networkid := transactions[i].NetworkId
		Owningbankid := transactions[i].OwningbankId
		Owningbankname := transactions[i].OwningbankName
		Time := transactions[i].Time
		Name := transactions[i].Name
		Mulescore := transactions[i].MuleScore
		Sourcetransactionvalue := transactions[i].SourcetransactionValue
		Endpointflag := transactions[i].EndpointFlag
		Numoutboundrelationships := transactions[i].NumoutboundRelationships
		Numinboundrelationships := transactions[i].NuminboundRelationships
		Numscheduledmandates := transactions[i].NumscheduledMandates
		Firstappearance := transactions[i].FirstAppearance
		Mostrecentappearance := transactions[i].MostrecentAppearance
		Receivessalary := transactions[i].ReceivesSalary
		Dwelltime := transactions[i].DwellTime
		Numnetworks := transactions[i].NumNetworks
		Numtracednetworks := transactions[i].NumtracedNetworks
		Generation := transactions[i].Generation
		Tracetype := transactions[i].Tracetype
		Totalsuspiciousvalueinbound := transactions[i].TotalsuspiciousValueinbound
		Totalsuspiciousvalueoutbound := transactions[i].TotalsuspiciousValueoutbound
		Totalvalueinbound := transactions[i].TotalvalueInbound
		Totalvalueoutbound := transactions[i].TotalvalueOutbound
		Generations := transactions[i].Generations
		Mostrecentfeedback := transactions[i].MostrecentFeedback
		Parentalertid := transactions[i].ParentalertId
		Decisiondate := transactions[i].DecisionDate
		Nextpaginationtoken := transactions[i].NextpaginationToken
		Previouspaginationtoken := transactions[i].PreviouspaginationToken

		message := fmt.Sprintf(" Success %s ", Errors)
		loggers.Alertaccount(c.Path(), "folderName", Uniqueidalertaccount, ID, message, Networkalertid, Accountid, Networkid, Owningbankid, Owningbankname, Time, Name, Mulescore, Sourcetransactionvalue, Endpointflag, Numoutboundrelationships, Numinboundrelationships, Numscheduledmandates, Firstappearance, Mostrecentappearance, Receivessalary, Dwelltime, Numnetworks, Numtracednetworks, Generation, Tracetype, Totalsuspiciousvalueinbound, Totalsuspiciousvalueoutbound, Totalvalueinbound, Totalvalueoutbound, Generations, Mostrecentfeedback, Parentalertid, Decisiondate, Nextpaginationtoken, Previouspaginationtoken)
	}
	//---------------------------

	// Referenceid := Referenceid(ID)
	// Generateid := Generateid(ID)
	type Alert struct {
		Alerts       fiber.Map      `json:"alerts"`
		Transactions []models.Alert `json:"transactions"`
	}

	response := Alert{
		Alerts: fiber.Map{
			"totalRecords":            count,
			"displayedRecords":        userRequest.Limit,
			"nextPaginationToken":     transactions[0].NextpaginationToken,
			"previousPaginationToken": transactions[0].PreviouspaginationToken,
			// "referenceID":             Referenceid,
			// "generateID":              Generateid,
		},

		Transactions: transactions,
	}

	return c.JSON(response)
}

func Iftgeneratealertaccount(length int) string {
	rand.Seed(time.Now().UnixNano())

	hexDigits := "0123456789abcdef"
	result := make([]byte, length)
	for i := range result {
		result[i] = hexDigits[rand.Intn(len(hexDigits))]
	}
	return string(result)
}

// func Referenceid(instructionID string) string {
// 	var ipsReference string
// 	for ctr := 1; ctr <= 10; ctr++ {
// 		ipsReference = instructionID[len(instructionID)-ctr:]
// 	}
// 	return fmt.Sprintf("IFT%v-%v", ipsReference, time.Now().UnixMilli())
// }

// func Generateid(instructionID string) string {
// 	var ipsReference string
// 	for ctr := 1; ctr <= 10; ctr++ {
// 		ipsReference = instructionID[len(instructionID)-ctr:]
// 	}
// 	return fmt.Sprintf("%v", ipsReference)
// }

//404
//400
//200
//429
//405

//401
//403
