package credittransfer

import (
	"fmt"
	"math/rand"
	"time"

	"tracealert/pkg/models/credittransfer"
	errorhandling "tracealert/pkg/models/errorHandling"
	"tracealert/pkg/utils/go-utils/database"

	"github.com/gofiber/fiber/v2"
)

var requestCounts int

const allowedRates = 5

func checkRateLimits() bool {
	requestCounts++
	if requestCounts > allowedRates {
		requestCounts = 0
		return true
	}
	return false
}
func TransferAccount(c *fiber.Ctx) error {
	trans := &credittransfer.Trans_Body{}

	if parsErr := c.BodyParser(trans); parsErr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": parsErr.Error(),
		})
	}
	if checkRateLimits() {
		// 429
		return errorhandling.Rate_Limit_Exceeded(c, "Rate limit exceeded ")
	}

	// Generate ID
	instructionid := Referenceid()
	referenceid := Iftgenerate()
	requestTrigger := time.Now().Format("2006-01-02 03:04:05")

	response := &credittransfer.Trans_Request{
		Instructionid: instructionid,
		Referenceid:   referenceid,
	}

	if trans.Transactionamount > 100 {
		response.Transactiontype = "SENDING"
		response.Reasoncode = "ACTC"
		response.Status = "SUCCESS"
		response.Localinstrument = "ICRT"
		response.Instructionid = instructionid
		response.Referenceid = referenceid
		response.Senderbic = "CAMZPHM2XXX"
		response.Receivingbic = "CBMFPHM1XXX"
		response.Description = "Transaction processed successfully"
		response.Amountcurrency = "PHP"
	} else {
		response.Transactiontype = "RECEIVING"
		response.Reasoncode = "AC01"
		response.Status = "FAILED-RJCT"
		response.Localinstrument = "ICRT"
		response.Instructionid = instructionid
		response.Referenceid = referenceid
		response.Senderbic = "CAMZPHM2XXX"
		response.Receivingbic = "CBMFPHM1XXX"
		response.Description = "Invalid transaction amount"
		response.Amountcurrency = "PHP"
	}

	insertQuery := `INSERT INTO trace_alerts.credit_transfer( instructionid,amountcurrency, description,  localinstrument, reasoncode, receivingaccount, receivingbic, receivingname, referenceid, senderaccount, senderamount, senderbic, sendername, status, transactiontype,datetime)
		VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?,?)`

	transacSave := database.DBConn.Debug().Exec(insertQuery,
		instructionid, response.Amountcurrency, response.Description, response.Localinstrument, response.Reasoncode,
		trans.Recipientaccountnumber, response.Receivingbic, trans.Recipientaccountname, referenceid,
		trans.Senderaccountnumber, trans.Transactionamount, response.Senderbic, trans.Senderaccountname,
		response.Status, response.Transactiontype, requestTrigger).Error

	if transacSave != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}

	err := database.DBConn.Raw("SELECT * FROM trace_alerts.credit_transfer WHERE instructionid = ?", instructionid).Scan(&response).Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error retrieving inserted data"})
	}
	return c.JSON(response)
}

func Referenceid() string {
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
