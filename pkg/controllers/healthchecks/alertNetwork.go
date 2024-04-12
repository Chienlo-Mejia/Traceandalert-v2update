package healthchecks

import (
	"fmt"
	"time"
	"tracealert/pkg/models"
	errorhandling "tracealert/pkg/models/errorHandling"
	"tracealert/pkg/utils/go-utils/database"

	"github.com/gofiber/fiber/v2"
)

var requestCounts_alertnetwork int

const allowedRates_alertnetwork = 5

func checkRateLimits_alertnetwork() bool {
	requestCounts_alertnetwork++
	if requestCounts_alertnetwork > allowedRates_alertnetwork {
		requestCounts_alertnetwork = 0
		return true
	}
	return false
}

func AlertNetwork(c *fiber.Ctx) error {
	network := &models.NetworkBody{}

	if checkRateLimits_alertnetwork() {
		//429
		return errorhandling.Rate_Limit_Exceeded(c, "You have exceeded the service rate limit. Maximum allowed: ${rate_limit.output} TPS")
	}

	if c.Method() != fiber.MethodPost {

		//405
		return errorhandling.Method_Not_Allowed(c, "Method_Not_Allowed")

	}

	if parsErr := c.BodyParser(&network); parsErr != nil {

		//400
		return errorhandling.Bad_Request(c, "The request contains a bad payload")
	} else if network == nil || (network.Since == "" && network.Limit == 0 && network.PaginationToken == "" && network.Filter == "" && !network.Include_all_alerts) {

		//400
		return errorhandling.Bad_Request(c, "The request body is empty")
	}

	// models
	account := models.TransactionAlert{}
	transaction := &models.Alertnetwork{}
	networks := &models.NetworkResponse{}

	// database
	if err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.trace_alerts`).Scan(&account).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"data": err.Error()})
	}
	if err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.trace_account_alerts`).Scan(transaction).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"data": err.Error()})
	}
	if err := database.DBConn.Debug().Raw(`SELECT * FROM trace_alerts.trace_vizurl`).Scan(networks).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"data": err.Error()})
	}

	var transactions []models.TransactionAlert

	query := `SELECT * FROM trace_alerts.trace_alerts WHERE 1 = 1`
	countQuery := "SELECT COUNT(*) FROM trace_alerts.trace_alerts WHERE 1=1"
	countFilteredQuery := "SELECT COUNT(*) FROM trace_alerts.trace_alerts WHERE 1=1"

	if network.Since != "" {
		// Validate the date format
		if _, err := time.Parse("2006-01-02", network.Since); err != nil {
			//400
			return errorhandling.Bad_Request(c, "Please provide a valid date format (YYYY-MM-DD)")
		}

		query += " AND DATE_TRUNC('day', decisiondate) = DATE_TRUNC('day', ?::date)"
		countFilteredQuery += " AND DATE_TRUNC('day', decisiondate) = DATE_TRUNC('day', ?::date)"
	}

	if network.Filter != "" {
		query += " AND " + network.Filter
		countFilteredQuery += " AND " + network.Filter
	}

	if network.Filter != "" {
		if network.Filter == "%3D%3D" {
			query += " AND (id='1234567890'  ) "
			countFilteredQuery += " AND (id='1234567890')"
		} else if network.Filter == "%3E%3D" {
			query += " AND (meanmulescore <= 0.75 OR numactionedmules < 2 OR numlegitimate < 2 OR numnotinvestigated < 2)"
			countFilteredQuery += " AND (meanmulescore <= 0.75 OR numactionedmules < 2 OR numlegitimate < 2 OR numnotinvestigated < 2)"
		} else if network.Filter == "%3C%3D" {
			query += " AND (length <= 100 OR totalvalue <= 2807544.66 OR sourcevalue <= 5615)"
			countFilteredQuery += " AND (length <= 100 OR totalvalue <= 2807544.66 OR sourcevalue <= 5615)"
		} else if network.Filter == "%3E" {
			query += " AND (generation < 10 OR elapsetime < 3600)"
			countFilteredQuery += " AND (generation < 10 OR elapsetime < 3600)"
		} else if network.Filter == "%3C" {
			query += " AND (uniqueaccounts > 10)"
			countFilteredQuery += " AND ( uniqueaccounts > 10 OR EXTRACT(EPOCH FROM meandwelltime) < 300 OR EXTRACT(EPOCH FROM mediandwelltime) < 3600)"
		} else {
			query += " AND (" + network.Filter + ")"
			countFilteredQuery += " AND (" + network.Filter + ")"
		}
	}

	var totalCount int64
	if err := database.DBConn.Raw(countQuery).Count(&totalCount).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error counting all transactions",
		})
	}

	var filteredCount int64
	if network.Since != "" {
		if err := database.DBConn.Raw(countFilteredQuery, network.Since).Count(&filteredCount).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error counting filtered transactions"})
		}
	} else {
		if err := database.DBConn.Raw(countFilteredQuery).Count(&filteredCount).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error counting filtered transactions"})
		}
	}

	// Limit
	if network.Limit > 0 {
		if int64(network.Limit) > filteredCount {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Limit exceeds the total number of filtered transactions",
			})
		}
		query += fmt.Sprintf(" LIMIT %d", network.Limit)
	}

	//query
	err := database.DBConn.Debug().Raw(query, network.Since).Scan(&transactions).Error

	if err != nil {
		//400
		return errorhandling.Bad_Request(c, "The request contains a bad payload")
	}

	if len(transactions) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "No data found for the specified date",
		})
	}
	//-----------------------------------------------------------------------------------

	var transactions1 []models.Alertnetwork

	query1 := `SELECT * FROM trace_alerts.trace_account_alerts WHERE 1 = 1`
	countQuery1 := "SELECT COUNT(*) FROM trace_alerts.trace_account_alerts WHERE 1=1"
	countFilteredQuery1 := "SELECT COUNT(*) FROM trace_alerts.trace_account_alerts WHERE 1=1"

	if network.Since != "" {
		if _, err := time.Parse("2006-01-02", network.Since); err != nil {
			//400
			return errorhandling.Bad_Request(c, "Please provide a valid date format (YYYY-MM-DD)")
		}

		query1 += " AND DATE_TRUNC('day', decisiondate) = DATE_TRUNC('day', ?::date)"
		countFilteredQuery1 += " AND DATE_TRUNC('day', decisiondate) = DATE_TRUNC('day', ?::date)"
	}

	if network.Filter != "" {
		query1 += " AND " + network.Filter
		countFilteredQuery1 += " AND " + network.Filter
	}
	var totalCount1 int64
	if err := database.DBConn.Raw(countQuery1).Count(&totalCount1).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error counting all transactions",
		})
	}
	var filteredCount1 int64
	if network.Since != "" {
		if err := database.DBConn.Raw(countFilteredQuery1, network.Since).Count(&filteredCount1).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error counting filtered transactions"})
		}
	} else {
		if err := database.DBConn.Raw(countFilteredQuery1).Count(&filteredCount1).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error counting filtered transactions"})
		}
	}

	// Limit
	if network.Limit > 0 {
		if int64(network.Limit) > filteredCount1 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Limit exceeds the total number of filtered transactions",
			})
		}
		query1 += fmt.Sprintf(" LIMIT %d", network.Limit)
	}

	err1 := database.DBConn.Debug().Raw(query1, network.Since).Scan(&transactions1).Error

	if err1 != nil {
		//400
		return errorhandling.Bad_Request(c, "The request contains a bad payload")
	}

	if len(transactions1) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "No data found for the specified date",
		})
	}

	//-----------------------------------------------------------------------------------

	var transactions2 []models.NetworkResponse

	query2 := `SELECT * FROM trace_alerts.trace_vizurl WHERE 1 = 1`
	countQuery2 := "SELECT COUNT(*) FROM trace_alerts.trace_vizurl WHERE 1=1"
	countFilteredQuery2 := "SELECT COUNT(*) FROM trace_alerts.trace_vizurl WHERE 1=1"

	if network.Since != "" {
		if _, err := time.Parse("2006-01-02", network.Since); err != nil {
			//400
			return errorhandling.Bad_Request(c, "Please provide a valid date format (YYYY-MM-DD)")
		}

		query2 += " AND DATE_TRUNC('day', decisiondate) = DATE_TRUNC('day', ?::date)"
		countFilteredQuery2 += " AND DATE_TRUNC('day', decisiondate) = DATE_TRUNC('day', ?::date)"
	}

	if network.Filter != "" {
		query1 += " AND " + network.Filter
		countFilteredQuery2 += " AND " + network.Filter
	}

	var totalCount2 int64
	if err := database.DBConn.Raw(countQuery2).Count(&totalCount2).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error counting all transactions",
		})
	}

	var filteredCount2 int64
	if network.Since != "" {
		if err := database.DBConn.Raw(countFilteredQuery2, network.Since).Count(&filteredCount2).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error counting filtered transactions"})
		}
	} else {
		if err := database.DBConn.Raw(countFilteredQuery2).Count(&filteredCount2).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error counting filtered transactions"})
		}
	}

	// Limit
	if network.Limit > 0 {
		if int64(network.Limit) > filteredCount2 {
			return errorhandling.Bad_Request(c, "Limit exceeds the total number of filtered transactions")
		}
		query1 += fmt.Sprintf(" LIMIT %d", network.Limit)
	}

	err2 := database.DBConn.Debug().Raw(query2, network.Since).Scan(&transactions2).Error

	if err2 != nil {
		//400
		return errorhandling.Bad_Request(c, "The request contains a bad payload")

	}

	if len(transactions1) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "No data found for the specified date",
		})
	}

	// Prepare response
	type AlertnetworkResponse struct {
		Alerts        fiber.Map               `json:"alerts"`
		Tracealerts   models.Alertnetwork     `json:"tracealerts"`
		AccountAlerts models.TransactionAlert `json:"accountAlerts"`
		Network       models.NetworkResponse  `json:"network"`
	}
	alerts := fiber.Map{
		"totalCount":    totalCount,
		"filteredCount": filteredCount,
		"id":            transactions[0].ID,
		"Time":          transactions1[0].Time,
		"networkID":     transactions1[0].Networkalertid,
	}

	response := AlertnetworkResponse{
		Alerts:        alerts,
		Tracealerts:   transactions1[0],
		AccountAlerts: transactions[0],
		Network:       transactions2[0],
	}

	return c.JSON(response)
}

//404/
//200/
//400/
//405/
//429/

//401
//403
