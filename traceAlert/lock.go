package traceandalert

import (
	"tracealert/pkg/utils/go-utils/database"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Lock struct {
	SenderAccount string `json:"senderAccount"`
}
type Lockdata struct {
	UniqueidCredittransfer string `json:"uniqueidCredittransfer"`
	Message                string `json:"message"`
	TraceAlert             string `json:"traceAlert"`
	SourceTxnType          string `json:"sourceTxnType"`
	TransactionType        string `json:"transactionType"`
	Status                 string `json:"status"`
	ReasonCode             string `json:"reasonCode"`
	LocalInstrument        string `json:"localInstrument"`
	InstructionId          string `json:"instructionId"`
	ReferenceId            string `json:"referenceId"`
	RequestTrigger         string `json:"requestTrigger"`
	AlertType              string `json:"alertType"`
	SenderAccount          string `json:"senderAccount"`
	Feedback               string `json:"feedback"`
	Lock                   string `json:"lock"`
}

func LockAccount(c *fiber.Ctx) error {
	Lockacc := &Lock{}
	var feedback string

	if parsErr := c.BodyParser(Lockacc); parsErr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": parsErr.Error(),
		})
	}

	if err := database.DBConn.Raw("SELECT feedback FROM trace_alerts.logsfeedback_credittransfer WHERE sender_account = ?", Lockacc.SenderAccount).Row().Scan(&feedback); err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Sender account not found.",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve account data.",
		})
	}

	switch feedback {
	case "ACTIONED_MULE":
		if err := database.DBConn.Exec("UPDATE trace_alerts.logsfeedback_credittransfer SET lock = ? WHERE sender_account = ?", "lock", Lockacc.SenderAccount).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to lock the account.",
			})
		}

		return c.Status(fiber.StatusLocked).JSON(fiber.Map{
			"error": "Account is locked due to fraudulent activity. Please contact the bank for assistance.",
		})
	case "NOT_INVESTIGATED":
		if err := database.DBConn.Exec("UPDATE trace_alerts.logsfeedback_credittransfer SET lock = ? WHERE sender_account = ?", "unlock", Lockacc.SenderAccount).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to unlock the account.",
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Account is unlocked.",
		})
	default:
		if err := database.DBConn.Exec("UPDATE trace_alerts.logsfeedback_credittransfer SET lock = ? WHERE sender_account = ?", "None", Lockacc.SenderAccount).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to unlock the account.",
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Account is unlocked.",
		})
	}
}

// --unlock
func UnlockAccount(c *fiber.Ctx) error {
	Lockacc := &Lock{}
	var lock string

	if parsErr := c.BodyParser(Lockacc); parsErr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": parsErr.Error(),
		})
	}

	if err := database.DBConn.Raw("SELECT lock FROM trace_alerts.logsfeedback_credittransfer WHERE sender_account = ?", Lockacc.SenderAccount).Row().Scan(&lock); err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Sender account not found.",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve account data.",
		})
	}

	// If the account is locked, unlock it
	if lock == "lock" {
		if err := database.DBConn.Exec("UPDATE trace_alerts.logsfeedback_credittransfer SET lock = ? WHERE sender_account = ?", "unlock", Lockacc.SenderAccount).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to unlock the account.",
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Account is unlocked.",
		})
	}

	// If the account is not locked, return an error
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"error": "Account is not locked or does not exist.",
	})
}

// func Lockaccount(c *fiber.Ctx) error {
// 	Lockacc := &Lock{}
// 	ctResult := []Lockdata{}

// 	if parsErr := c.BodyParser(Lockacc); parsErr != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"error": parsErr.Error(),
// 		})
// 	}

// 	for _, data := range ctResult {
// 		if data.LogsLock == "lock" {
// 			return c.Status(fiber.StatusLocked).JSON(fiber.Map{
// 				"error": "Account is locked. Please contact the bank to unlock the account.",
// 			})
// 		}
// 	}

// 	if err := database.DBConn.Raw("SELECT * FROM trace_alerts.join_feedback_credittransfer(?)", Lockacc.InstructionId).Scan(&ctResult).Error; err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"error": "Failed to retrieve account data.",
// 		})
// 	}

//		return c.JSON(ctResult)
//	}
