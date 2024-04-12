package credittransfer

// import "github.com/gofiber/fiber/v2"

// type BlockTransactionRequest struct {
// 	InstructionId string `json:"instructionId"`
// }

// func BlockTransaction(c *fiber.Ctx) error {
// 	var req BlockTransactionRequest
// 	if err := c.BodyParser(&req); err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"error": err.Error(),
// 		})
// 	}

// 	BlockedTransactions.Lock()
// 	BlockedTransactions.transactions[req.InstructionId] = true
// 	BlockedTransactions.Unlock()

// 	return c.JSON(fiber.Map{
// 		"message": "Transaction blocked successfully",
// 	})
// }
// func UnblockTransaction(c *fiber.Ctx) error {
// 	var req BlockTransactionRequest
// 	if err := c.BodyParser(&req); err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"error": err.Error(),
// 		})
// 	}

// 	BlockedTransactions.Lock()
// 	delete(BlockedTransactions.transactions, req.InstructionId)
// 	BlockedTransactions.Unlock()

// 	return c.JSON(fiber.Map{
// 		"message": "Transaction unblocked successfully",
// 	})
// }
