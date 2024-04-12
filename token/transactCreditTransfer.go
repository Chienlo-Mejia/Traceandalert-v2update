package credittransfer

// type Transaction struct {
// 	ID     string  `json:"id"`
// 	Type   string  `json:"type"`
// 	Amount float64 `json:"amount"`
// }

// var (
// 	errorLogger         *log.Logger
// 	transactionLogger   *log.Logger
// 	BlockedTransactions struct {
// 		sync.RWMutex
// 		transactions map[string]bool
// 	}
// )

// func init() {
// 	BlockedTransactions.transactions = make(map[string]bool)

// 	// Initialize loggers
// 	errorLog, err := os.OpenFile("error.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
// 	if err != nil {
// 		log.Fatalf("Failed to open error log file: %v", err)
// 	}
// 	errorLogger = log.New(errorLog, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

// 	transactionLog, err := os.OpenFile("transaction.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
// 	if err != nil {
// 		log.Fatalf("Failed to open transaction log file: %v", err)
// 	}
// 	transactionLogger = log.New(transactionLog, "TRANSACTION: ", log.Ldate|log.Ltime)
// }

// func Transactions(c *fiber.Ctx) error {
// 	var transaction Transaction
// 	if err := c.BodyParser(&transaction); err != nil {
// 		errorLogger.Printf("Error parsing transaction body: %v", err)
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"error": err.Error(),
// 		})
// 	}

// 	BlockedTransactions.RLock()
// 	defer BlockedTransactions.RUnlock()
// 	if BlockedTransactions.transactions[transaction.ID] {
// 		errorLogger.Printf("Transaction blocked: %s", transaction.ID)
// 		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
// 			"error": "Transaction is blocked",
// 		})
// 	}

// 	// Assuming saveTransaction is a function that saves the transaction to the database
// 	if err := saveTransaction(&transaction); err != nil {
// 		errorLogger.Printf("Error saving transaction: %v", err)
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"error": "Failed to save transaction",
// 		})
// 	}

// 	transactionLogger.Printf("Date: %s", time.Now().Format("2006-01-02 15:04:05"))
// 	transactionLogger.Printf("ID: %s", transaction.ID)
// 	transactionLogger.Printf("Type: %s", transaction.Type)
// 	transactionLogger.Printf("Amount: %.2f", transaction.Amount)

// 	return c.JSON(fiber.Map{
// 		"message": "Transaction saved successfully",
// 	})
// }

// func saveTransaction(transaction *Transaction) error {
// 	return nil
// }
