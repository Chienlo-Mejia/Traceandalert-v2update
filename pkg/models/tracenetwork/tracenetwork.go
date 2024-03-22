package tracenetwork

import "time"

type (
	Trans_Body struct {
		Since  string `json:"since"`
		Limit  int    `json:"limit"`
		Filter string `json:"filter"`
	}
	Trans_Request struct {
		// Instructionid    string    `json:"instructionid"`
		// Time             time.Time `josn:"time"`
		// Transactiontype  string    `json:"transactiontype"`
		// Status           string    `json:"status"`
		// Reasoncode       string    `json:"reasoncode"`
		// Description      string    `json:"description"`
		// Localinstrument  string    `json:"localinstrument"`
		// Referenceid      string    `json:"referenceid"`
		// Senderbic        string    `json:"senderbic"`
		// Sendername       string    `json:"sendername"`
		// Senderaccount    string    `json:"senderaccount"`
		// Amountcurrency   string    `json:"amountcurrency"`
		// Senderamount     float64   `json:"senderamount"`
		// Receivingbic     string    `json:"receivingbic"`
		// Receivingname    string    `json:"receivingname"`
		// Receivingaccount string    `json:"receivingaccount"`

		Instruction_id       string  `json:"instruction_id"`
		Transaction_datetime string  `josn:"transaction_datetime"`
		Transaction_type     string  `json:"transaction_type"`
		Status               string  `json:"status"`
		Reason_code          string  `json:"reason_code"`
		Local_instrument     string  `json:"local_instrument"`
		Reference_id         string  `json:"reference_id"`
		Sender_bic           string  `json:"sender_bic"`
		Sender_name          string  `json:"sender_name"`
		Sender_account       string  `json:"sender_account"`
		Amount               float64 `json:"amount"`
		Currency             string  `json:"currency"`
		Receiving_bic        string  `json:"receiving_bic"`
		Receiving_name       string  `json:"receiving_name"`
		Receiving_account    string  `json:"receiving_account"`
	}

	Credit_transfer struct {
		// Instruction_id    string `json:"instruction_id"`
		// Reference_id      string `json:"reference_id"`
		// Receiving_account string `json:"receiving_account"`
		// Receiving_name    string `json:"receiving_name"`
		// Sender_account    string `json:"sender_account"`
		// Sender_name       string `json:"sender_name"`
		// Filter           string `json:"filter"`
		Instruction_id    string `json:"instruction_id"`
		Reference_id      string `json:"reference_id"`
		Receiving_account string `json:"receiving_account"`
		Receiving_name    string `json:"receiving_name"`
		Sender_account    string `json:"sender_account"`
		Sender_name       string `json:"sender_name"`
		Filter            string `json:"filter"`
	}

	Credittransfer_feedback struct {
		Uniqueidcredittransfer string `json:"uniqueidcredittransfer"`
		Alerttype              string `json:"alerttype"`
		Tracetype              string `json:"tracetype"`
		Trace_alert            string `json:"trace_alert"`
	}

	Credittransfer_feedback_response struct {
		Uniqueidcredittransfer string `json:"uniqueidcredittransfer"`
		Message                string `json:"message"`
		Trace_alert            string `json:"trace_alert"`
		Sourcetxntype          string `json:"sourcetxntype"`
		Transaction_type       string `json:"transaction_type"`
		Status                 string `json:"status"`
		Reason_code            string `json:"reason_code"`
		Local_instrument       string `json:"local_instrument"`
		Instruction_id         string `json:"instruction_id"`
		Reference_id           string `json:"reference_id"`
		Alerttype              string `json:"alerttype"`
		Requesttrigger         string `json:"requesttrigger"`
	}

	Errorlog struct {
		ID          int       `json:"id"`
		ErrorCode   string    `json:"errorcode"`
		Source      string    `json:"source"`
		ReasonCode  string    `json:"reasoncode"`
		Description string    `json:"description"`
		Recoverable bool      `json:"recoverable"`
		Details     string    `json:"details"`
		CreatedAt   time.Time `json:"createdat"`
	}
)
