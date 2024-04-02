package credittransfer

import "time"

type (
	Trans_Body struct {
		Senderaccountnumber    string  `json:"senderaccountnumber"`
		Senderaccountname      string  `json:"senderaccountname"`
		Recipientbankcode      string  `json:"recipientbankcode"`
		Recipientaccountnumber string  `json:"recipientaccountnumber"`
		Recipientaccountname   string  `json:"recipientaccountname"`
		Transactionreference   string  `json:"transactionreference"`
		Transactionamount      float64 `json:"transactionamount"`
		Transactioncharge      float64 `json:"transactioncharge"`
	}

	Trans_Request struct {
		Instructionid    string    `json:"instructionid"`
		Transactiontype  string    `json:"transactiontype"`
		Status           string    `json:"status"`
		Reasoncode       string    `json:"reasoncode"`
		Description      string    `json:"description"`
		Localinstrument  string    `json:"localinstrument"`
		Referenceid      string    `json:"referenceid"`
		Senderbic        string    `json:"senderbic"`
		Sendername       string    `json:"sendername"`
		Senderaccount    string    `json:"senderaccount"`
		Amountcurrency   string    `json:"amountcurrency"`
		Senderamount     float64   `json:"senderamount"`
		Receivingbic     string    `json:"receivingbic"`
		Receivingname    string    `json:"receivingname"`
		Receivingaccount string    `json:"receivingaccount"`
		Datetime         time.Time `json:"datetime"`
	}
)

type (
	FDSAPRequestCreditTransfer struct {
		ReceivingBIC           string  `json:"receivingBIC,omitempty"`
		ReceivingAccountNumber string  `json:"receivingAccountNumber,omitempty"`
		ReceivingName          string  `json:"receivingName,omitempty"`
		SenderBIC              string  `json:"senderBIC,omitempty"`
		SenderName             string  `json:"senderName,omitempty"`
		SenderAccountNumber    string  `json:"senderAccountNumber,omitempty"`
		Amount                 float64 `json:"amount,omitempty"`
		Currency               string  `json:"currency,omitempty"`
		LocalInstrument        string  `json:"localInstrument,omitempty"`
		ReferenceID            string  `json:"referenceId,omitempty"`
		AppId                  string  `json:"appId,omitempty"`
		Application            string  `json:"application,omitempty"`
		// IDs
		BusinessMessageId string `json:"businessMessageId,omitempty"`
		MessageId         string `json:"messageId,omitempty"`
		InstructionId     string `json:"instructionId,omitempty"`
	}

	FDSIRequestCreditTransfer struct {
		SenderAccountNumber    string  `json:"SenderAccountNumber,omitempty"`
		SenderAccountName      string  `json:"SenderAccountName,omitempty"`
		RecipientBankCode      string  `json:"RecipientBankCode,omitempty"`
		RecipientAccountNumber string  `json:"RecipientAccountNumber,omitempty"`
		RecipientAccountName   string  `json:"RecipientAccountName,omitempty"`
		TransactionReference   string  `json:"TransactionReference,omitempty"`
		TransactionAmount      float64 `json:"TransactionAmount,omitempty"`
		TransactionCharge      float64 `json:"TransactionCharge,omitempty"`
		// IDs
		BusinessMessageId int `json:"businessMessageId,omitempty"`
		MessageId         int `json:"messageId,omitempty"`
		InstructionId     int `json:"instructionId,omitempty"`
	}

	ResponseCreditTransfer struct {
		TransactionType     string  `json:"transactionType,omitempty"`
		Status              string  `json:"status"`
		ReasonCode          string  `json:"reasonCode"`
		Description         string  `json:"description"`
		LocalInstrument     string  `json:"localInstrument,omitempty"`
		InstructionID       string  `json:"instructionId,omitempty"`
		TransactionID       string  `json:"transactionId,omitempty"`
		ReferenceID         string  `json:"referenceId,omitempty"`
		SenderBIC           string  `json:"senderBIC,omitempty"`
		SenderName          string  `json:"senderName,omitempty"`
		SenderAccount       string  `json:"senderAccount,omitempty"`
		AmountCurrency      string  `json:"amountCurrency,omitempty"`
		SenderAmount        float64 `json:"senderAmount,omitempty"`
		ReceivingBIC        string  `json:"receivingBIC,omitempty"`
		ReceivingName       string  `json:"receivingName,omitempty"`
		ReceivingAccount    string  `json:"receivingAccount,omitempty"`
		TransactionDateTime string  `json:"transactionDateTime,omitempty"`
		Pacs_008            string  `json:"pacs008,omitempty"`
		Pacs_002            string  `json:"pacs002,omitempty"`
	}

	SystemParameter struct {
		Parameter string `json:"parameter"`
		Value     string `json:"value"`
	}

	ReasonCode struct {
		Code        string `json:"code"`
		Description string `json:"description"`
	}
)

var (
	HeadMessageDefinition                    = "head.001.001.01" // Head
	SignOnRequestMessageDefinition           = "admn.001.001.01" // Sign-On Request - sr
	SignOnResponseMessageDefinition          = "admn.002.001.01" // Sign-On Response - rs
	SignOffRequestMessageDefinition          = "admn.003.001.01" // Sign-Off Request - fr
	SignOffResponseMessageDefinition         = "admn.004.001.01" // Sign-Off Response - rf
	SystemNotificationEventMessageDefinition = "admi.004.001.02" // System Event Notification - ne
	EchoRequestMessageDefinition             = "admn.005.001.01" // Echo Request - er
	EchoResponseMessageDefinition            = "admn.006.001.01" // Echo Response - re
	CreditTransferMessageDefinition          = "pacs.008.001.08" // Credit Transfer V08 - ct
	MessageStatusReportMessageDefinition     = "pacs.002.001.10" // Message Status Report V10 (Response to Business Messages) - ps
	MessageRejectMessageDefinition           = "admi.002.001.01" // Message Reject (Admin) - mr
	PaymentCancellationMessageDefinition     = "camt.056.001.08" // Payment Cancellation V08 (Request for Return of Funds or System Time-out) - rt
	RBIBankIdentifierCode                    = "CAMZPHM2XXX"
)
