package credittransfer

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

type Transaction struct {
	ID          string
	Amount      float64
	Source      string
	Destination string
	Timestamp   time.Time
}

type Trace struct {
	TransactionID string
	TraceType     string
	Timestamp     time.Time
}

func createFraudTrace(transaction Transaction) Trace {
	// Create a fraud trace for the given transaction
	fraudTrace := Trace{
		TransactionID: transaction.ID,
		TraceType:     "Fraud",
		Timestamp:     time.Now(),
	}
	return fraudTrace
}

func RunTraceService(c *fiber.Ctx) error {
	// Simulating a fraudulent transaction
	fraudulentTransaction := Transaction{
		ID:          "123456789",
		Amount:      1000.00,
		Source:      "Hacker",
		Destination: "Victim",
		Timestamp:   time.Now(),
	}

	// Generate fraud trace for the fraudulent transaction
	fraudTrace := createFraudTrace(fraudulentTransaction)

	// Simulating a non-fraudulent transaction
	nonFraudulentTransaction := Transaction{
		ID:          "987654321",
		Amount:      500.00,
		Source:      "Legitimate",
		Destination: "Recipient",
		Timestamp:   time.Now(),
	}

	// Create a map to hold both fraud trace and transaction data
	responseData := map[string]interface{}{
		"fraudTrace":               fraudTrace,
		"fraudulentTransaction":    fraudulentTransaction,
		"nonFraudulentTransaction": nonFraudulentTransaction,
	}

	return c.JSON(responseData)
}
