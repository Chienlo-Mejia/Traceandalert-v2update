package tracefraud

import "time"

type Logsfeedback struct {
	UniqueidFeedback string `json:"uniqueidFeedback"`
	AlertId          string `json:"alertId"`
	FeedbackId       string `json:"feedbackId"`
	Errors           string `json:"errors"`
}

type LogsAlert struct {
	UniqueidAlertaccount string `json:"uniqueidAlertaccount"`
	ID                   string `json:"id"`
	Message              string `json:"message"`
	NetworkalertId       string `json:"networkalertId"`
	AccountId            string `json:"accountId"`
	NetworkId            string `json:"networkId"`
	OwningbankId         string `json:"owningbankId"`
	OwningbankName       string `json:"owningbankName"`
	DecisionDate         string `json:"decisionDate"`
	ParentalertId        string `json:"parentalertId"`
}

type Logstracevisualisation struct {
	UniqueidVisualisation string `json:"uniqueidVisualisation"`
	NetworkalertId        string `json:"networkalertId"`
	TraceId               string `json:"traceId"`
	Errors                string `json:"errors"`
}

type Logstracenetwork struct {
	UniqueidNetwork string    `json:"uniqueidNetwork"`
	TxnidRb         string    `json:"txnidRb"`
	TxnId           string    `json:"txnId"`
	NetworkId       string    `json:"networkId"`
	SourcetxnId     string    `json:"sourcetxnId"`
	DecisionDate    time.Time `json:"decisionDate"`
	Errors          string    `json:"errors"`
}

type Logsalerttransaction struct {
	UniqueidAlertaccount string    `json:"uniqueidAlertaccount"`
	ID                   string    `json:"id"`
	Errors               string    `json:"errors"`
	TxnId                string    `json:"txnId"`
	NetworkalertId       string    `json:"networkalertId"`
	NetworkId            string    `json:"networkId"`
	SourceId             string    `json:"sourceId"`
	DestId               string    `json:"destId"`
	SourcebankId         string    `json:"sourcebankId"`
	SourcebankName       string    `json:"sourcebankName"`
	DestbankId           string    `json:"destbankId"`
	DestbankName         string    `json:"destbankName"`
	ParentalertId        string    `json:"parentalertId"`
	DecisionDate         time.Time `json:"decisionDate"`
}

type Logscredittransfer struct {
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
	RequestTrigger         string `json:"requestTrigger"`
	AlertType              string `json:"alertType"`
	Feedback               string `json:"feedback"`
	SenderAccount          string `json:"senderAccount"`
}

type LogsFeedback_credittransfer struct {
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
	RequestTrigger         string `json:"requestTrigger"`
	AlertType              string `json:"alertType"`
	Feedback               string `json:"feedback"`
	Lock                   string `json:"lock"`
	SenderAccount          string `json:"senderAccount"`
}
type TraceInfo struct {
	Id             string `json:"id"`
	Message        string `json:"message"`
	Requesttrigger string `json:"requesttrigger"`
}
