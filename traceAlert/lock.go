package traceandalert

import (
	"time"
	"tracealert/pkg/utils/go-utils/database"

	"github.com/gofiber/fiber/v2"
)

type Lock struct {
	InstructionId string `json:"instructionId"`
}

type Lockdata struct {
	CtID                       int       `json:"ct_id"`
	CtTransactionType          string    `json:"ct_transaction_type"`
	CtStatus                   string    `json:"ct_status"`
	CtReasonCode               string    `json:"ct_reason_code"`
	CtLocalInstrument          string    `json:"ct_local_instrument"`
	CtInstructionID            string    `json:"ct_instruction_id"`
	CtTransactionID            string    `json:"ct_transaction_id"`
	CtReferenceID              string    `json:"ct_reference_id"`
	CtSenderBIC                string    `json:"ct_sender_bic"`
	CtSenderName               string    `json:"ct_sender_name"`
	CtSenderAccount            string    `json:"ct_sender_account"`
	CtCurrency                 string    `json:"ct_currency"`
	CtAmount                   float64   `json:"ct_amount"`
	CtReceivingBIC             string    `json:"ct_receiving_bic"`
	CtReceivingName            string    `json:"ct_receiving_name"`
	CtReceivingAccount         string    `json:"ct_receiving_account"`
	CtTransactionDateTime      string    `json:"ct_transaction_datetime"`
	CtApplication              string    `json:"ct_application"`
	CtFTDate                   time.Time `json:"ct_ft_date"`
	CtPACSEight                string    `json:"ct_pacs_eight"`
	CtPACSTwo                  string    `json:"ct_pacs_two"`
	CtDevice                   string    `json:"ct_device"`
	LogsUniqueIDCreditTransfer string    `json:"logs_uniqueid_credittransfer"`
	LogsMessage                string    `json:"logs_message"`
	LogsTraceAlert             string    `json:"logs_trace_alert"`
	LogsSourceTxnType          string    `json:"logs_source_txn_type"`
	LogsTransactionType        string    `json:"logs_transaction_type"`
	LogsStatus                 string    `json:"logs_status"`
	LogsReasonCode             string    `json:"logs_reason_code"`
	LogsLocalInstrument        string    `json:"logs_local_instrument"`
	LogsInstructionID          string    `json:"logs_instruction_id"`
	LogsReferenceID            string    `json:"logs_reference_id"`
	LogsRequestTrigger         string    `json:"logs_request_trigger"`
	LogsAlertType              string    `json:"logs_alert_type"`
	LogsFeedback               string    `json:"logs_feedback"`
	LogsLock                   string    `json:"logs_lock"`
}

func Lockaccount(c *fiber.Ctx) error {
	Lockacc := &Lock{}
	ctResult := []Lockdata{}

	if parsErr := c.BodyParser(Lockacc); parsErr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": parsErr.Error(),
		})
	}

	for _, data := range ctResult {
		if data.LogsLock == "lock" {
			return c.Status(fiber.StatusLocked).JSON(fiber.Map{
				"error": "Account is locked. Please contact the bank to unlock the account.",
			})
		}
	}

	if err := database.DBConn.Raw("SELECT * FROM trace_alerts.join_feedback_credittransfer(?)", Lockacc.InstructionId).Scan(&ctResult).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve account data.",
		})
	}

	return c.JSON(ctResult)
}
