package credittransfer

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"tracealert/pkg/models/creditTransfer"
	"tracealert/pkg/utils/go-utils/database"

	"github.com/gofiber/fiber/v2"
)

type BlockedID struct {
	Senderaccountnumber string `json:"senderaccountnumber"`
	ExpiresAt           int64  `json:"expires_at"`
}

var blockedIDs map[string]*BlockedID = make(map[string]*BlockedID)
var accountLocks sync.Mutex

func BlockID(c *fiber.Ctx) error {
	var blockRequest struct {
		Senderaccountnumber string `json:"Senderaccountnumber"`
		ExpiresIn           int64  `json:"expires_in,omitempty"`
	}

	if err := c.BodyParser(&blockRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	accountLocks.Lock()
	defer accountLocks.Unlock()

	var expiresAt int64
	if blockRequest.ExpiresIn > 0 {
		expiresAt = time.Now().Add(time.Duration(blockRequest.ExpiresIn) * time.Second).UnixMilli()
	} else {
		expiresAt = time.Now().Add(5 * time.Second).UnixMilli()
	}

	blockedIDs[blockRequest.Senderaccountnumber] = &BlockedID{Senderaccountnumber: blockRequest.Senderaccountnumber, ExpiresAt: expiresAt}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"Senderaccountnumber": fmt.Sprintf(" %s blocked for transactions (expires at %d)", blockRequest.Senderaccountnumber, expiresAt),
	})
}

func TransferAccount(c *fiber.Ctx) error {
	trans := &creditTransfer.Trans_Body{}
	if parsErr := c.BodyParser(trans); parsErr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": parsErr.Error(),
		})
	}
	accountLocks.Lock()
	defer accountLocks.Unlock()

	if blockedID, ok := blockedIDs[trans.Senderaccountnumber]; ok {
		if blockedID.ExpiresAt > time.Now().UnixMilli() {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"ALERT": fmt.Sprintf("This Account is Maintanance for Senderaccountnumber: %s (expires at %d)", trans.Senderaccountnumber, blockedID.ExpiresAt),
			})
		} else {

			delete(blockedIDs, trans.Senderaccountnumber)
		}
	}

	instructionId := ReferenceId()
	referenceId := Iftgenerate()
	requestTrigger := time.Now().Format("2006-01-02 03:04:05")

	response := &creditTransfer.Trans_Request{
		InstructionId:   instructionId,
		ReferenceId:     referenceId,
		TransactionType: "RECEIVING",
		SenderBic:       "CAMZPHM2XXX",
		ReceivingBic:    "CBMFPHM1XXX",
		AmountCurrency:  "PHP",
		LocalInstrument: "ICRT",
		Description:     "Invalid transaction amount",
		ReasonCode:      "AC01",
		Status:          "FAILED-RJCT",
	}

	if trans.Transactionamount > 100 {
		response.TransactionType = "SENDING"
		response.ReasonCode = "ACTC"
		response.Status = "SUCCESS"
		response.Description = "Transaction processed successfully"
	}

	insertQuery := `INSERT INTO trace_alerts.credit_transfer( instruction_id, amount_currency, description,  local_instrument, reason_code, receiving_account, receiving_bic, receiving_name, reference_id, sender_account, sender_amount, sender_bic, sender_name, status, transaction_type, date_time)
        VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	transacSave := database.DBConn.Debug().Exec(insertQuery,
		instructionId, response.AmountCurrency, response.Description, response.LocalInstrument, response.ReasonCode,
		trans.Recipientaccountnumber, response.ReceivingBic, trans.Recipientaccountname, referenceId,
		trans.Senderaccountnumber, trans.Transactionamount, response.SenderBic, trans.Senderaccountname,
		response.Status, response.TransactionType, requestTrigger).Error

	if transacSave != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}

	logEntry := Log{UserID: 1, Action: "access", Timestamp: time.Now()}
	MonitorForDataBreaches(logEntry)
	go cleanupBlockedIDs()
	return c.JSON(response)
}

type Log struct {
	UserID    int
	Action    string
	Timestamp time.Time
}

func MonitorForDataBreaches(logEntry Log) {

	if logEntry.Action == "access" && logEntry.Timestamp.Hour() < 6 {
		alertMsg := fmt.Sprintf("ALERT: Potential unauthorized access by UserID %d at %s", logEntry.UserID, logEntry.Timestamp)
		fmt.Println(alertMsg)

	}
}

func cleanupBlockedIDs() {
	for {
		accountLocks.Lock()
		for id, blockedID := range blockedIDs {
			if blockedID.ExpiresAt <= time.Now().UnixMilli() {
				delete(blockedIDs, id)
			}
		}
		accountLocks.Unlock()
		time.Sleep(time.Minute)
	}
}

func ReferenceId() string {
	rand.Seed(time.Now().UnixNano())
	uniqueDigit := rand.Intn(10000)
	return fmt.Sprintf("20240219CAMZPHM2XXXB0000000000%d", uniqueDigit)
}

func Iftgenerate() string {
	rand.Seed(time.Now().UnixNano())
	uniqueDigit := rand.Intn(1000)
	uniqueDigit1 := rand.Intn(10000000000000)
	return fmt.Sprintf("IFT%v-%v", uniqueDigit, uniqueDigit1)
}
