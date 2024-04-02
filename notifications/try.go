package email

type Transaction struct {
	Amount   float64
	Location string
}

func DetectGeographicalAnomalies(transactions []Transaction) []Transaction {
	var unusualTransactions []Transaction
	locationCounts := make(map[string]int)

	// Count occurrences of each location
	for _, transaction := range transactions {
		locationCounts[transaction.Location]++
	}

	// Identify locations with significantly fewer transactions
	for location, count := range locationCounts {
		// You can adjust the threshold as needed
		if count <= 1 {
			for _, transaction := range transactions {
				if transaction.Location != location {
					unusualTransactions = append(unusualTransactions, transaction)
				}
			}
		}
	}

	return unusualTransactions
}
