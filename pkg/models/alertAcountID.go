package models

import "time"

type (
	AccountAlertRequest1 struct {
		AccountId string `json:"accountId"`
	}

	ResponseInfo1 struct {
		ID                           string    `json:"Id"`
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
		NumInboundRelationships      int       `json:"numinboundRelationships"`
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
		TraceType                    string    `json:"traceType"`
		TotalsuspiciousvalueInbound  int       `json:"totalsuspiciousvalueInbound"`
		TotalsuspiciousvalueOutbound int       `json:"totalsuspiciousvalueOutbound"`
		TotalvalueInbound            int       `json:"totalvalueInbound"`
		TotalvalueOutbound           int       `json:"totalvalueOutbound"`
		Generations                  []int     `json:"generations"`
		MostrecentFeedback           string    `json:"mostrecentFeedback"`
		ParentalertId                string    `json:"parentalertId"`
		DecisionDate                 time.Time `json:"decisionDate"`
	}
	Token struct {
		NextpaginationToken     string `json:"nextpaginationToken"`
		PreviouspaginationToken string `json:"previouspaginationToken"`
	}
)
