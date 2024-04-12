package email

import (
	"fmt"
	"net/smtp"
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/robfig/cron/v3"
)

// type EmailMessage struct {
// 	To           []string `json:"to"`
// 	Message      string   `json:"message"`
// 	Detail       string   `json:"detail"`
// 	ScheduleTime string   `json:"schedule_time"`
// }

// func Email(c *fiber.Ctx) error {
// 	var input EmailMessage

// 	if err := c.BodyParser(&input); err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"message": "Error parsing the request body.",
// 			"error":   err.Error(),
// 		})
// 	}

// 	pass := "ooqz lcqj rfki dyer"
// 	from := "chichimejia660@gmail.com"

// 	smtpHost := "smtp.gmail.com"
// 	smtpPort := "587"

// 	auth := smtp.PlainAuth("", from, pass, smtpHost)

// 	errCh := make(chan error, len(input.To))

// 	for _, to := range input.To {
// 		go func(to string) {

// 			msg := "From: " + from + "\r\n" +
// 				"To: " + to + "\r\n" +
// 				"Subject: Customized Alert: Action Required\r\n" +
// 				"MIME-version: 1.0;\r\n" +
// 				"Content-Type: text/html; charset=\"UTF-8\";\r\n\r\n" +
// 				"<html><head><style>" +
// 				"body { font-family: 'Arial', sans-serif; background-color: #f7f7f7; color: #333; margin: 0; padding: 0; }" +
// 				"#container { max-width: 600px; margin: 0 auto; padding: 20px; background-color: #fff; border-radius: 10px; box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1); }" +
// 				"h1 { color: #003366; font-size: 24px; margin-bottom: 20px; }" +
// 				"p { font-size: 16px; margin-bottom: 10px; }" +
// 				".urgent { color: #e74c3c; font-weight: bold; }" +
// 				".highlight { color: #1f497d; }" +
// 				".signature { font-style: italic; }" +
// 				"</style></head><body>" +
// 				"<div id='container'>" +
// 				"<h1 style='text-align: center;'>Customized Alert: Action Required</h1>" +
// 				"<p>Dear: " + to + ",</p>" +
// 				"<p class='urgent'>Urgent action is needed:</p>" +
// 				"<p class='highlight'>Message: " + input.Message + "</p>" +
// 				"<p class='highlight'>Privacy Notice: " + input.Detail + "</p>" +
// 				"<p style='text-align: right;' class='signature'>Best regards,<br/>INSTAPAY Team</p>" +
// 				"</div></body></html>"

// 			err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, []byte(msg))
// 			if err != nil {
// 				fmt.Println("Email sending failed for recipient", to, ":", err)
// 				errCh <- err
// 				return
// 			}
// 			fmt.Println("Email sent successfully to", to)
// 			errCh <- nil
// 		}(to)
// 	}

// 	for range input.To {
// 		err := <-errCh
// 		if err != nil {

// 			close(errCh)
// 			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 				"message": "Error sending emails.",
// 				"error":   err.Error(),
// 			})
// 		}
// 	}

// 	return c.JSON(fiber.Map{
// 		"message": "Emails sent successfully.",
// 	})
// }

// ----------------------------------------

// type EmailMessage struct {
// 	To      []string `json:"to"`
// 	Message string   `json:"message"`
// 	Detail  string   `json:"detail"`
// }

// func Email(c *fiber.Ctx) error {
// 	emailData := EmailMessage{
// 		To:      []string{"cpe.anista.bienmarc@gmail.com", "chienlo.mejia@fortress-asya.com", "rico.vergara@fortress-asya.com", "chienlomejia660@gmail.com", "chichimejia660@gmail.com", "rosendo.ocubillo@fortress-asya.com", "josephmejia660@gmail.com", "mapeebun@gmail.com", "arren.agravante@fortress-asya.com", "bien.anista@fortress-asya.com"},
// 		Message: "Urgent Problem",
// 		Detail:  "We are currently experiencing technical difficulties with our payment processing system. This may cause delays in processing transactions. Our team is actively working to resolve the issue as quickly as possible. We apologize for any inconvenience this may cause.",
// 	}

// 	pass := "ooqz lcqj rfki dyer"
// 	from := "chichimejia660@gmail.com"

// 	smtpHost := "smtp.gmail.com"
// 	smtpPort := "587"

// 	auth := smtp.PlainAuth("", from, pass, smtpHost)

// 	var mu sync.Mutex
// 	var errors []error
// 	emailSize := 30000
// 	numEmail := (len(emailData.To) + emailSize - 1) / emailSize

// 	for i := 0; i < numEmail; i++ {
// 		startIdx := i * emailSize
// 		endIdx := (i + 1) * emailSize
// 		if endIdx > len(emailData.To) {
// 			endIdx = len(emailData.To)
// 		}

// 		batch := emailData.To[startIdx:endIdx]

// 		var wg sync.WaitGroup

// 		for _, to := range batch {
// 			wg.Add(1)
// 			go func(to string) {
// 				defer wg.Done()

// 				msg := "From: " + from + "\r\n" +
// 					"To: " + to + "\r\n" +
// 					"Subject: Customized Alert: Action Required\r\n" +
// 					"MIME-version: 1.0;\r\n" +
// 					"Content-Type: text/html; charset=\"UTF-8\";\r\n\r\n" +
// 					"<html><head><style>" +
// 					"body { font-family: 'Arial', sans-serif; background-color: #f7f7f7; color: #333; margin: 0; padding: 0; }" +
// 					"#container { max-width: 600px; margin: 0 auto; padding: 20px; background-color: #fff; border-radius: 10px; box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1); }" +
// 					"h1 { color: #003366; font-size: 24px; margin-bottom: 20px; }" +
// 					"p { font-size: 16px; margin-bottom: 10px; }" +
// 					".urgent { color: #e74c3c; font-weight: bold; }" +
// 					".highlight { color: #1f497d; }" +
// 					".signature { font-style: italic; }" +
// 					"</style></head><body>" +
// 					"<div id='container'>" +
// 					"<h1 style='text-align: center;'>Customized Alert: Action Required</h1>" +
// 					"<p>Dear: " + to + ",</p>" +
// 					"<p class='urgent'>Urgent action is needed:</p>" +
// 					"<p class='highlight'>Message: " + emailData.Message + "</p>" +
// 					"<p class='highlight'>Privacy Notice: " + emailData.Detail + "</p>" +
// 					"<p style='text-align: right;' class='signature'>Best regards,<br/>INSTAPAY Team</p>" +
// 					"</div></body></html>"

// 				err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, []byte(msg))
// 				if err != nil {
// 					fmt.Printf("Email sending failed for recipient %s: %s\n", to, err)
// 					mu.Lock()
// 					errors = append(errors, err)
// 					mu.Unlock()
// 				} else {
// 					fmt.Println("Email sent successfully to", to)
// 				}
// 			}(to)
// 		}

// 		wg.Wait()
// 	}

// 	if len(errors) > 0 {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"message": "Error sending emails.",
// 			"errors":  errors,
// 		})
// 	}

// 	return c.JSON(fiber.Map{
// 		"message": "Emails sent successfully.",
// 	})
// }

type EmailMessage struct {
	Email   []string `json:"email"`
	Message string   `json:"message"`
	Detail  string   `json:"detail"`
}

func Email(c *fiber.Ctx, emailData EmailMessage) error {
	pass := "ooqz lcqj rfki dyer"
	from := "chichimejia660@gmail.com"

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	auth := smtp.PlainAuth("", from, pass, smtpHost)

	var mu sync.Mutex
	var errors []error
	emailSize := 30000
	numEmail := (len(emailData.Email) + emailSize - 1) / emailSize

	for i := 0; i < numEmail; i++ {
		startIdx := i * emailSize
		endIdx := (i + 1) * emailSize
		if endIdx > len(emailData.Email) {
			endIdx = len(emailData.Email)
		}

		batch := emailData.Email[startIdx:endIdx]

		var wg sync.WaitGroup

		for _, to := range batch {
			wg.Add(1)
			go func(to string) {
				defer wg.Done()

				msg := "From: " + from + "\r\n" +
					"To: " + to + "\r\n" +
					"Subject: Customized Alert: Action Required\r\n" +
					"MIME-version: 1.0;\r\n" +
					"Content-Type: text/html; charset=\"UTF-8\";\r\n\r\n" +
					"<html><head><style>" +
					"body { font-family: 'Arial', sans-serif; background-color: #f7f7f7; color: #333; margin: 0; padding: 0; }" +
					"#container { max-width: 600px; margin: 0 auto; padding: 20px; background-color: #fff; border-radius: 10px; box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1); }" +
					"h1 { color: #003366; font-size: 24px; margin-bottom: 20px; }" +
					"p { font-size: 16px; margin-bottom: 10px; }" +
					".urgent { color: #e74c3c; font-weight: bold; }" +
					".highlight { color: #1f497d; }" +
					".signature { font-style: italic; }" +
					"</style></head><body>" +
					"<div id='container'>" +
					"<h1 style='text-align: center;'>Customized Alert: Action Required</h1>" +
					"<p>Dear: " + to + ",</p>" +
					"<p class='urgent'>Urgent action is needed:</p>" +
					"<p class='highlight'>Message: " + emailData.Message + "</p>" +
					"<p class='highlight'>Privacy Notice: " + emailData.Detail + "</p>" +
					"<p style='text-align: right;' class='signature'>Best regards,<br/>INSTAPAY Team</p>" +
					"</div></body></html>"

				err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, []byte(msg))
				if err != nil {
					fmt.Printf("Email sending failed for recipient %s: %s\n", to, err)
					mu.Lock()
					errors = append(errors, err)
					mu.Unlock()
				} else {
					fmt.Println("Email sent successfully to", to)
				}
			}(to)
		}

		wg.Wait()
	}

	if len(errors) > 0 {
		return fmt.Errorf("error sending emails: %v", errors)
	}

	return nil
}

var Feedback string

func CronJobEmail() {
	c := cron.New()
	_, err := c.AddFunc("*/5 * * * *", func() {

		if Feedback == "ACTIONED_MULE" {
			emailData := EmailMessage{
				Email:   []string{"cpe.anista.bienmarc@gmail.com", "chienlo.mejia@fortress-asya.com", "rico.vergara@fortress-asya.com", "chienlomejia660@gmail.com", "chichimejia660@gmail.com", "rosendo.ocubillo@fortress-asya.com", "josephmejia660@gmail.com", "mapeebun@gmail.com", "arren.agravante@fortress-asya.com", "bien.anista@fortress-asya.com"},
				Message: "Urgent Problem",
				Detail:  "We are currently experiencing technical difficulties with our payment processing system. This may cause delays in processing transactions. Our team is actively working to resolve the issue as quickly as possible. We apologize for any inconvenience this may cause.",
			}

			if err := Email(nil, emailData); err != nil {
				fmt.Println("Failed to send email:", err)
			}
		}
	})
	if err != nil {
		fmt.Println("Error scheduling cron job:", err)
	}

	c.Start()
}

// ----------------------

// func Email(c *fiber.Ctx) error {
// 	var input EmailMessage
// 	// requestTrigger := time.Now().Format("2006-01-02 15:04:05")
// 	if err := c.BodyParser(&input); err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"message": "Error parsing the request body.",
// 			"error":   err.Error(),
// 		})
// 	}

// 	cronJob := cron.New()

// 	_, err := cronJob.AddFunc(input.ScheduleTime, func() {
// 		//minutes
// 		//24hrs
// 		//days
// 		//month
// 		if strings.Contains(input.ScheduleTime, strconv.Itoa(time.Now().Year())) {
// 			sendEmail(input)
// 		} else {

// 			sendEmail(input)
// 		}
// 	})

// 	if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"message": "Error scheduling cron job.",
// 			"error":   err.Error(),
// 		})
// 	}
// 	cronJob.Start()

// 	return c.JSON(fiber.Map{
// 		"message": "Message scheduled successfully.",
// 	})
// }

// func sendEmail(input EmailMessage) {
// 	pass := "ooqz lcqj rfki dyer"
// 	from := "chichimejia660@gmail.com"

// 	smtpHost := "smtp.gmail.com"
// 	smtpPort := "587"

// 	auth := smtp.PlainAuth("", from, pass, smtpHost)

// 	var mu sync.Mutex
// 	var errors []error

// 	// Email size can receive 30000 message
// 	emailSize := 30000
// 	numEmail := (len(input.To) + emailSize - 1) / emailSize

// 	for i := 0; i < numEmail; i++ {
// 		startIdx := i * emailSize
// 		endIdx := (i + 1) * emailSize
// 		if endIdx > len(input.To) {
// 			endIdx = len(input.To)
// 		}

// 		batch := input.To[startIdx:endIdx]

// 		var wg sync.WaitGroup

// 		for _, to := range batch {
// 			wg.Add(1)
// 			go func(to string) {
// 				defer wg.Done()

// 				msg := "From: " + from + "\r\n" +
// 					"To: " + strings.Join(input.To, ",") + "\r\n" +
// 					"Subject: Customized Alert: Action Required\r\n" +
// 					"MIME-version: 1.0;\r\n" +
// 					"Content-Type: text/html; charset=\"UTF-8\";\r\n\r\n" +
// 					"<html><head><style>" +
// 					"body { font-family: 'Arial', sans-serif; background-color: #f7f7f7; color: #333; margin: 0; padding: 0; }" +
// 					"#container { max-width: 600px; margin: 0 auto; padding: 20px; background-color: #fff; border-radius: 10px; box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1); }" +
// 					"h1 { color: #003366; font-size: 24px; margin-bottom: 20px; }" +
// 					"p { font-size: 16px; margin-bottom: 10px; }" +
// 					".urgent { color: #e74c3c; font-weight: bold; }" +
// 					".highlight { color: #1f497d; }" +
// 					".signature { font-style: italic; }" +
// 					"</style></head><body>" +
// 					"<div id='container'>" +
// 					"<h1 style='text-align: center;'>Customized Alert: Action Required</h1>" +
// 					"<p>Dear: " + to + "</p>" +
// 					"<p class='urgent'>Urgent action is needed:</p>" +
// 					"<p class='highlight'>Message: " + input.Message + "</p>" +
// 					"<p class='highlight'>Privacy Notice: " + input.Detail + "</p>" +
// 					"<p style='text-align: right;' class='signature'>Best regards,<br/>INSTAPAY Team</p>" +
// 					"</div></body></html>"

// 				err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, []byte(msg))
// 				if err != nil {
// 					fmt.Printf("Email sending failed for recipient %s: %s\n", to, err)
// 					mu.Lock()
// 					errors = append(errors, err)
// 					mu.Unlock()
// 				} else {
// 					fmt.Println("Email sent successfully to", to)
// 				}
// 			}(to)
// 		}

// 		wg.Wait()
// 	}

// 	if len(errors) > 0 {
// 		fmt.Println("Error sending emails:", errors)
// 		return
// 	}

// 	fmt.Println("Emails sent successfully.")
// }
