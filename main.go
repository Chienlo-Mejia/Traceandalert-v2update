package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"
	"time"
	"tracealert/pkg/config"
	routers "tracealert/pkg/routers"
	middleware "tracealert/pkg/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
)

func main() {

	fmt.Print("Hello world")

	h1 := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.Execute(w, nil)
	}
	http.HandleFunc("/", h1)
	// log.Fatal(http.ListenAndServe(":8000", nil))

	//----------------------------------------------

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error Loading Env File: ", err)
	}
	envi := os.Getenv("ENVIRONMENT")

	err = godotenv.Load(fmt.Sprintf("project_env_files/.env-%v", envi)) //
	if err != nil {
		log.Fatal("Error Loading Env File: ", err)
	}

	// Initialize DB here
	config.CreateConnection()
	// Declare & initialize fiber
	app := fiber.New(fiber.Config{
		Views:        html.New("./template", ".gohtml"),
		UnescapePath: true,
	})

	app.Static("/", "./asset/image")
	app.Static("/", "./asset/js")
	app.Static("/", "./asset/css")

	// For GoRoutine implementation
	// appb := fiber.New(fiber.Config{
	// 	UnescapePath: true,
	// })

	// Configure application CORS

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",                                           // means all user can access the API or you can just specify the user that can use the system (example: facebook.com) only
		AllowHeaders: "Origin, Content-Type, Accept, Authorization", // header that we only accept``
	}),
		logger.New(logger.Config{ //Modified Logs
			Format:     "${cyan}${time} ${white}| ${green}${status} ${white}| ${ip} | ${host} | ${method} | ${magenta}${path} ${white}| ${red}${latency} ${white}\n",
			TimeFormat: "01/02/2006 3:04 PM",
			// TimeZone:   envRouting.PostgresTimeZone,
		}))

	// For GoRoutine implementation
	// appb.Use(cors.New(cors.Config{
	// 	AllowOrigins: "*",
	// 	AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	// }))

	// Declare & initialize logger
	app.Use(logger.New())

	// For GoRoutine implementation
	// appb.Use(logger.New())

	//---------------- Call Endpoint ---------------
	routers.SetupPublicRoutes(app)
	routers.SetupPrivateRoutes(app)
	//----------------------------------------------
	// For GoRoutine implementation
	// routers.SetupPublicRoutesB(appb)
	// go func() {
	// 	err := appb.Listen(fmt.Sprintf(":8002"))
	// 	if err != nil {
	// 		log.Fatal(err.Error())
	// 	}
	// }()

	app.Use(func(c *fiber.Ctx) error {
		if err := c.Next(); err != nil {
			if e, ok := err.(*fiber.Error); ok && e.Code == fiber.StatusNotFound {
				return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
					"Errors": fiber.Map{
						"Error": []fiber.Map{
							{
								"Source":      "Gateway",
								"ReasonCode":  "NOT_FOUND",
								"Description": "URL not found",
								"Recoverable": false,
								"Details":     nil,
							},
						},
					},
				})
			}
			return err
		}
		return nil
	})
	//----------

	//----------

	fmt.Println("Port:", middleware.GetEnv("PORT"))
	// Serve the application
	if middleware.GetEnv("SSL") == "enabled" {
		log.Fatal(app.ListenTLS(
			fmt.Sprintf(":%s", middleware.GetEnv("PORT")),
			middleware.GetEnv("SSL_CERTIFICATE"),
			middleware.GetEnv("SSL_KEY"),
		))
	} else {
		err := app.Listen(fmt.Sprintf(":%s", middleware.GetEnv("PORT")))
		if err != nil {
			log.Fatal(err.Error())
		}
	}

}

func checkForFraud() {
	for {

		currentTime := time.Now()
		fmt.Println("Running fraud tracing at:", currentTime)

		suspiciousTransactions := []struct {
			Amount    float64
			Recipient string
		}{
			{Amount: 5000.0, Recipient: "Fraudster1"},
			{Amount: 10000.0, Recipient: "Fraudster2"},
		}

		// Check if there are any suspicious transactions
		if len(suspiciousTransactions) > 0 {
			fmt.Println("Suspicious transactions found:")
			for _, transaction := range suspiciousTransactions {
				fmt.Printf("Amount: %.2f, Recipient: %s\n", transaction.Amount, transaction.Recipient)
			}
			// Here you could take appropriate action, such as sending alerts, freezing accounts, etc.
		} else {
			fmt.Println("No suspicious transactions found.")
		}

		// Delay for one second before checking again
		time.Sleep(time.Second)
	}
}
