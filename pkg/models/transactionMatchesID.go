package models

import "time"

type (
	TransacIDmatches struct {
		Service         string  `json:"service"`
		Sourceid        string  `json:"sourceid"`
		Beneficiaryid   string  `json:"beneficiaryid"`
		Transactiontime string  `json:"transactiontime"`
		Amount          float64 `json:"amount"`
	}

	TransactionMatches struct {
		Id            string    `json:"id"`
		Sourceid      string    `json:"sourceid"`
		Beneficiaryid string    `json:"beneficiaryid"`
		Reftime       time.Time `json:"reftime"`
		Amount        float64   `json:"amount"`
		Time          time.Time `josn:"time "`
		Txnid         string    `josn:"txnid"`
	}
	RequestBodyArraymatches struct {
		Items []RequestBody `json:"items"`
	}
)
