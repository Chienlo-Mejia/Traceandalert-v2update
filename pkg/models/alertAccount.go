package models

import "time"

type RequestBodyalert struct {
	Since           string `json:"since"`
	Limit           int    `json:"limit"`
	PaginationToken string `json:"paginationToken"`
	Filter          string `json:"filter"`
}

// Response represents the structure of the response

// ErrorResponse represents the structure of an error response
type ErrorResponse struct {
	Errors ErrorDetail `json:"Errors"`
}

// ErrorDetail represents the structure of error details
type ErrorDetailalert struct {
	Error []ErrorInfo `json:"Error"`
}

// ErrorInfo represents detailed information about an error
type ErrorInfo struct {
	Source      string `json:"Source"`
	ReasonCode  string `json:"ReasonCode"`
	Description string `json:"Description"`
	Recoverable bool   `json:"Recoverable"`
	Details     string `json:"Details"`
}

type AlertStruct struct {
	TotalRecords            int     `json:"totalRecords"`
	DisplayedRecords        int     `json:"displayedRecords"`
	NextPaginationToken     string  `json:"nextPaginationToken"`
	PreviousPaginationToken string  `json:"previousPaginationToken"`
	Alerts                  []Alert `json:"alerts"`
}

type Alert struct {
	ID                           string    `json:"id"`
	NetworkalertId               string    `json:"networkalertId"`
	AccountId                    string    `json:"accountId"`
	NetworkId                    string    `json:"networkId"`
	OwningbankId                 string    `json:"owningbankId"`
	OwningbankName               string    `json:"owningbankName"`
	Time                         time.Time `json:"time"`
	Name                         string    `json:"name"`
	MuleScore                    float64   `json:"muleScore"`
	SourcetransactionValue       int       `json:"sourcetransactionValue"`
	EndpointFlag                 bool      `json:"endpointFlag"`
	NumoutboundRelationships     int       `json:"numoutboundRelationships"`
	NuminboundRelationships      int       `json:"numinboundRelationships"`
	NumscheduledMandates         int       `json:"numscheduledMandates"`
	FirstAppearance              time.Time `json:"firstAppearance"`
	MostrecentAppearance         time.Time `json:"mostrecentAppearance"`
	FirstTransactiontime         time.Time `json:"firstTransactiontime"`
	MostrecentTransactiontime    time.Time `json:"mostrecentTransactiontime"`
	ReceivesSalary               bool      `json:"receivesSalary"`
	DwellTime                    string    `json:"dwellTime"`
	NumNetworks                  int       `json:"numNetworks"`
	NumtracedNetworks            int       `json:"numtracedNetworks"`
	Generation                   int       `json:"generation"`
	Tracetype                    string    `json:"traceType"`
	TotalsuspiciousValueinbound  int       `json:"totalsuspiciousValueinbound"`
	TotalsuspiciousValueoutbound int       `json:"totalsuspiciousValueoutbound"`
	TotalvalueInbound            int       `json:"totalvalueInbound"`
	TotalvalueOutbound           int       `json:"totalvalueOutbound"`
	Generations                  []int     `json:"generations"`
	MostrecentFeedback           string    `json:"mostrecentFeedback"`
	ParentalertId                string    `json:"parentalertId"`
	DecisionDate                 time.Time `json:"decisionDate"`
	NextpaginationToken          string    `json:"nextpaginationToken"`
	PreviouspaginationToken      string    `json:"previouspaginationToken"`
}
