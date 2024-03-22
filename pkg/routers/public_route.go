package routers

import (
	"fmt"
	"strings"
	"time"
	"tracealert/controller"
	"tracealert/credittransfer"
	"tracealert/middleware/loggers"
	notifications "tracealert/notifications"
	"tracealert/pkg/controllers/healthchecks"
	"tracealert/traceandalert"

	"github.com/gofiber/fiber/v2"
)

func SetupPublicRoutes(app *fiber.App) {

	//--------------------------------------------- Trace_Transaction_(POSTMAN,MOBILE_PHONE,BROWSER)
	var browserDetected bool
	app.Use(func(c *fiber.Ctx) error {
		userAgent := c.Get("User-Agent")
		requestPath := c.Path()
		requestTrigger := time.Now()

		// Postman
		if (requestPath == "/Email" || requestPath == "/Alert_credit") && userAgent != "" && (strings.Contains(userAgent, "Postman") || strings.Contains(userAgent, "PostmanRuntime")) {
			fmt.Printf("Request from Postman detected for the %s endpoint!\n", requestPath)

			loggers.Detectpostmanlogs(requestPath, "folderName", fmt.Sprintf("Transaction detected in POSTMAN from the endpoint of %s!", requestPath), requestTrigger)
			c.SendString(fmt.Sprintf("Transaction detected in Go for the %s endpoint!", requestPath))
		}

		// Flutter app
		if strings.Contains(userAgent, "Dart/3.1 (dart:io)") {
			deviceId := c.Get("Device-ID")
			if deviceId != "" {
				fmt.Printf("Request from Flutter app detected for the %s endpoint! Device ID: %s\n", requestPath, deviceId)
				loggers.Detectpostmanlogs(requestPath, "folderName", fmt.Sprintf("Transaction detected in MOBILEPHONE from the Device-ID %s!", deviceId), requestTrigger)
			} else {
				fmt.Printf("Request from Flutter app detected for the %s endpoint, but no Device-ID header found! %s\n", requestPath, deviceId)
				loggers.Detectpostmanlogs(requestPath, "folderName", fmt.Sprintf("Transaction detected in MOBILEPHONE, but no Device-ID header found for endpoint %s!", requestPath), requestTrigger)
			}
		}

		// Browser
		if (requestPath == "/Email" || requestPath == "/Alert_credit") &&
			(strings.Contains(userAgent, "Mozilla") || strings.Contains(userAgent, "Chrome") || strings.Contains(userAgent, "Safari")) &&
			!browserDetected {
			fmt.Printf("Request from a browser detected for the %s endpoint!\n", requestPath)
			loggers.Detectpostmanlogs(requestPath, "folderName", fmt.Sprintf("Transaction detected in BROWSER for the %s endpoint!", requestPath), requestTrigger)
			browserDetected = true // Set the flag to true
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

	//------------------------- Feedback --------------------------------//
	//feedback
	app.Post("/feedback", healthchecks.Feedback)
	//--------------------------- Trace ---------------------------------//

	//trace
	tracenetwork.Post("/tracenetwork", healthchecks.Tracenetwork)
	tracevisualisations.Post("/network", healthchecks.NetworkAlertID)
	//--------------------------- Alerts ---------------------------------//

	//alert
	alerts.Post("/accounts", healthchecks.Alertsaccount)
	alerts.Post("/transactions", healthchecks.Alerttransaction)
	alerts.Post("/networks", healthchecks.Alertnetwork)
	//alertID
	alertaccount.Post("/account_alert_id", healthchecks.GetAccInfo)
	alerttransaction.Post("/txn_alert_id", healthchecks.GetTransacInfo)
	alertnetworks.Post("/network_alert_id", healthchecks.GetNetworkInfo)
	//--------------------------- Match ---------------------------------//
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

	//----------------------------- CREDIT TRANSFER
	app.Post("/TransferAccount", credittransfer.TransferAccount)
	app.Post("/Token", credittransfer.Generateandsettoken)

	app.Post("/try", credittransfer.RunTraceService)

	app.Post("/Trace_credit", traceandalert.Tracenetworkcred)
	app.Post("/Alert_credit", traceandalert.Alertnetworkcredit)
	app.Post("/Matchesid_credit", traceandalert.Matchesidcredit)
	app.Post("/Feedbackid_credit", traceandalert.Feedbackcredit)

	app.Post("/Alertaccount_credit", traceandalert.Alertaccount_Credittransfer)
	app.Post("/Alerttransaction_credit", traceandalert.Alerttransaction_Credittransfer)
	app.Post("/trace_trans_PostmanMobilephoneBrowser", traceandalert.TracetransPostmanMobilephoneBrowser)

	//------------------------------Notifications
	app.Post("/Email", notifications.Email)
	app.Post("/Sms", notifications.Sms)
	app.Post("/Sms_telerivet", notifications.Sms_telerivet)
	// app.Post("/trace_location", notifications.Getclientlocation)

	// v1Endpoint.Post("/feedback", healthchecks.Feedback)

	// v1Endpoint.Post("/accounts", healthchecks.Alertsaccount)
	// v1Endpoint.Post("/transactions", healthchecks.Alerttransaction)
	// v1Endpoint.Post("/networks", healthchecks.Alertnetwork)

	// v1Endpoint.Post("/txn_alert_id", healthchecks.GetTransacInfo)
	// v1Endpoint.Post("/account_alert_id", healthchecks.GetAccInfo)
	// v1Endpoint.Post("/network_alert_id", healthchecks.GetNetworkInfo)

	// v1Endpoint.Post("/network", healthchecks.NetworkAlertID)
	// v1Endpoint.Post("/tracenetwork", healthchecks.Tracenetwork)

	// v1Endpoint.Post("/Transactionidmatches", healthchecks.Transactionidmatches)

	//--------------- Encrypt & decrypt -------------
	app.Post("/Encrypt", notifications.Encryptdata)
	app.Post("/Decrypt", notifications.Decryptdata)

}

func SetupPublicRoutesB(app *fiber.App) {

	// Endpoints
	apiEndpoint := app.Group("/api")
	publicEndpoint := apiEndpoint.Group("/public")
	v1Endpoint := publicEndpoint.Group("/v1")

	// Service health check
	v1Endpoint.Get("/", healthchecks.CheckServiceHealthB)
}
