package routers

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"strings"
	"time"
	"tracealert/controller"
	"tracealert/middleware/loggers"
	notifications "tracealert/notifications"
	"tracealert/pkg/controllers/healthchecks"
	token "tracealert/token"
	traceandalert "tracealert/traceAlert"
	traceCode "tracealert/traceAlert/traceCode"

	"github.com/gofiber/fiber/v2"
)

var browserSessions = make(map[string]bool)

func SetupPublicRoutes(app *fiber.App) {

	//--------------------------------------------- Trace_Transaction_(POSTMAN,MOBILE_PHONE,BROWSER) ---------------------------
	app.Use(func(c *fiber.Ctx) error {
		userAgent := c.Get("User-Agent")
		requestPath := c.Path()
		requestTrigger := time.Now()

		getSessionID := func() string {
			id := make([]byte, 16)
			_, err := rand.Read(id)
			if err != nil {
				return time.Now().Format(time.RFC3339Nano)
			}
			return base64.URLEncoding.EncodeToString(id)
		}

		//--------------------------- Postman ---------------------------
		if (requestPath == "/Email" || requestPath == "/Alert_credit") && userAgent != "" && (strings.Contains(userAgent, "Postman") || strings.Contains(userAgent, "PostmanRuntime")) {
			fmt.Printf("Request from Postman detected for the %s endpoint!\n", requestPath)
			loggers.Detectpostmanlogs(requestPath, "folderName", fmt.Sprintf("Transaction detected in POSTMAN from the endpoint of %s!", requestPath), requestTrigger)
			c.SendString(fmt.Sprintf("Transaction detected in Go for the %s endpoint!", requestPath))
		}
		//--------------------------- Flutter app ---------------------------
		if strings.Contains(userAgent, "Dart/3.1 (dart:io)") {
			deviceId := c.Get("Device-ID")
			if deviceId != "" {
				fmt.Printf("Request from Flutter app detected for the %s endpoint! Device ID: %s\n", requestPath, deviceId)
				loggers.Detectpostmanlogs(requestPath, "folderName", fmt.Sprintf("Transaction detected in MOBILEPHONE from the Device-ID %s!", deviceId), requestTrigger)
			} else {
				fmt.Printf("Request from Flutter app detected for the %s endpoint, but no Device-ID header found!\n", requestPath)
				loggers.Detectpostmanlogs(requestPath, "folderName", fmt.Sprintf("Transaction detected in MOBILEPHONE, but no Device-ID header found for endpoint %s!", requestPath), requestTrigger)
			}
			return nil
		}

		//--------------------------- Browser
		if (requestPath == "/Email" || requestPath == "/Alert_credit") &&
			(strings.Contains(userAgent, "Mozilla") || strings.Contains(userAgent, "Chrome") || strings.Contains(userAgent, "Safari")) {
			sessionID := getSessionID()
			if !browserSessions[sessionID] {
				fmt.Printf("Request from a browser detected for the %s endpoint!\n", requestPath)
				loggers.Detectpostmanlogs(requestPath, "folderName", fmt.Sprintf("Transaction detected in BROWSER for the %s endpoint!", requestPath), requestTrigger)
				browserSessions[sessionID] = true
			}
			return nil
		}

		return c.Next()
	})

	//--------------------------- Endpoint ------------------------------//

	// Endpoints
	apiEndpoint := app.Group("/api")
	publicEndpoint := apiEndpoint.Group("/public")
	v1Endpoint := publicEndpoint.Group("/v1")

	//trace
	trace := app.Group("/financial-crime")
	tracenetwork := trace.Group("/networks")
	tracevisualisations := trace.Group("/visualisations")

	//alerts
	alerts := app.Group("/alerts")
	alertaccount := alerts.Group("/account")
	alerttransaction := alerts.Group("/transactions")
	alertnetworks := alerts.Group("/networks")

	//------------------------- Ex-Feedback --------------------------------//
	//feedback
	app.Post("/feedback", healthchecks.Feedback)
	//--------------------------- Ex-Trace ---------------------------------//

	//trace
	tracenetwork.Post("/tracenetwork", healthchecks.TraceNetwork)
	tracevisualisations.Post("/network", healthchecks.NetworkAlertID)
	//--------------------------- Ex-Alerts ---------------------------------//

	//alert
	alerts.Post("/accounts", healthchecks.AlertsAccount)
	alerts.Post("/transactions", healthchecks.AlertTransaction)
	alerts.Post("/networks", healthchecks.AlertNetwork)
	//alertID
	alertaccount.Post("/account_alert_id", healthchecks.GetAccInfo)
	alerttransaction.Post("/txn_alert_id", healthchecks.GetTransacInfo)
	alertnetworks.Post("/network_alert_id", healthchecks.GetNetworkInfo)
	//--------------------------- Ex-Match ---------------------------------//
	//match-id
	app.Post("/Transactionidmatches", healthchecks.Transactionidmatches)

	//---------------------  Service health check -----------------------//
	// Service health check
	v1Endpoint.Get("/", healthchecks.CheckServiceHealth)

	//--------------------------- WEB -------------------------//
	webEnpoint := app.Group("/web")
	apiSample := webEnpoint.Group("/trace_alert")

	apiSample.Get("xample/", controller.ShowPage)

	webEnpoint.Get("/Feedbackinfo", healthchecks.Feedbackinfo)           //
	webEnpoint.Get("/alertaccount", healthchecks.Alertaccinfo)           //
	webEnpoint.Get("/Tracevisualisationinfo", healthchecks.Tracevisinfo) //
	webEnpoint.Get("/Tracenetworkinfo", healthchecks.Tracenetworkinfo)
	webEnpoint.Get("/Alerttransactioninfo", healthchecks.Alerttransactioninfo)

	webEnpoint.Get("/Credit_transfer", healthchecks.Credittransferinfo)
	webEnpoint.Get("/Feedback_Credittransfer", healthchecks.FeedbackCredittransferinfo)
	webEnpoint.Get("/TracePostmanMobileBrowser", healthchecks.TracePostmanMobileBrowser)

	webEnpoint.Get("/login", controller.ShowPage1)
	webEnpoint.Get("/Location", controller.Location)
	app.Post("/TransferAccount", token.TransferAccount)
	app.Post("/LockAcc", token.BlockID)
	app.Post("/UnlockAcc", token.UnblockID)
	//----------------------------- Token -------------

	app.Post("/Token", token.Generateandsettoken)
	app.Post("/ip-details", token.Handledetails)

	app.Get("/MonitoringFraud", token.MonitoringFraud)
	// app.Get("/Location1", location.GetLocation)

	//----------------------------- Trace and Alert -------------
	// app.Post("/try", credittransfer.RunTraceService)
	app.Post("/Alertaccount_credit", traceCode.AlertaccountCredittransfer)
	app.Post("/Alerttransaction_credit", traceCode.AlerttransactionCredittransfer)

	app.Post("/Trace_credit", traceandalert.TracenetworkCred)
	app.Post("/alert-credit", traceandalert.AlertFraud)
	app.Post("/Matchesid_credit", traceandalert.MatchesidCredit)
	app.Post("/Feedbackid_credit", traceandalert.FeedbackCredit)
	app.Post("/trace_trans_PostmanMobilephoneBrowser", traceandalert.TracetransPostmanMobilephoneBrowser)

	//----------------------------- Lock account -------------
	// app.Post("/GenerateCreditTransfer", token.Transactions)
	// app.Post("/Lock", token.BlockTransaction)
	// app.Post("/UnlockLock", token.UnblockTransaction)
	app.Post("/Lockaccount", traceandalert.LockAccount)
	app.Post("/Unlockaccount", traceandalert.UnlockAccount)

	//------------------------------ Notifications -------------
	// app.Post("/Email", notifications.Email)
	app.Post("/Sms", notifications.Sms)
	// app.Post("/Sms_telerivet", notifications.Sms_telerivet)

	//------------------------------  Encrypt & decrypt -------------
	app.Post("/Encrypt", notifications.Encryptdata)
	app.Post("/Decrypt", notifications.Decryptdata)

	//------------------------------  Trace Fraud -------------
	app.Post("/TraceFraud", traceandalert.TraceFraud)
	app.Post("/FeedbackFraud", traceandalert.FeedbackFraud)

}

func SetupPublicRoutesB(app *fiber.App) {

	// Endpoints
	apiEndpoint := app.Group("/api")
	publicEndpoint := apiEndpoint.Group("/public")
	v1Endpoint := publicEndpoint.Group("/v1")

	// Service health check
	v1Endpoint.Get("/", healthchecks.CheckServiceHealthB)
}
