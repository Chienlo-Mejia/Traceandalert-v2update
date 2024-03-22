package tracenetwork

import "time"

type (
	Trans_Body struct {
		Since  string `json:"since"`
		Limit  int    `json:"limit"`
		Filter string `json:"filter"`
	}
	Trans_Request struct {
		InstructionId       string  `json:"instructionId"`
		TransactionDatetime string  `josn:"transactionDatetime"`
		TransactionType     string  `json:"transactionType"`
		Status              string  `json:"status"`
		ReasonCode          string  `json:"reasonCode"`
		LocalInstrument     string  `json:"localInstrument"`
		ReferenceId         string  `json:"referenceId"`
		SenderBic           string  `json:"senderBic"`
		SenderName          string  `json:"senderName"`
		SenderAccount       string  `json:"senderAccount"`
		Amount              float64 `json:"amount"`
		Currency            string  `json:"currency"`
		ReceivingBic        string  `json:"receivingBic"`
		ReceivingName       string  `json:"receivingName"`
		ReceivingAccount    string  `json:"receivingAccount"`
	}

	Credit_transfer struct {
		InstructionId    string `json:"instructionId"`
		ReferenceId      string `json:"referenceId"`
		ReceivingAccount string `json:"receivingAccount"`
		ReceivingName    string `json:"receivingName"`
		SenderAccount    string `json:"senderAccount"`
		SenderName       string `json:"senderName"`
		Filter           string `json:"filter"`
	}

	Credittransfer_feedback struct {
		UniqueidCredittransfer string `json:"uniqueidCredittransfer"`
		AlertType              string `json:"alertType"`
		TraceType              string `json:"traceType"`
		TraceAlert             string `json:"traceAlert"`
	}

	Credittransfer_feedback_response struct {
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
		AlertType              string `json:"alertType"`
		RequestTrigger         string `json:"requestTrigger"`
	}

	Errorlog struct {
		ID          int       `json:"id"`
		ErrorCode   string    `json:"errorCode"`
		Source      string    `json:"source"`
		ReasonCode  string    `json:"reasonCode"`
		Description string    `json:"description"`
		Recoverable bool      `json:"recoverable"`
		Details     string    `json:"details"`
		CreatedAt   time.Time `json:"createdAt"`
	}
)
