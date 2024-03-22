package models

type (
	RequestBodyArray1 struct {
		Items []RequestBody `json:"items"`
	}
	Tracevisualisationsalert struct {
		Networkalertid string `json:"networkalertid"`
		Format         string `json:"format"`
		Width          int64  `json:"width"`
		Height         int64  `json:"height"`
		Legend         bool   `json:"legend"`
		Type           string `json:"type"`
		Colourmode     string `json:"colourmode"`
		Status         bool   `json:"status"`

		// SomeData         string json:"somedata"
	}

	Traceresponse struct {
		Traceid string `json:"traceid2"`
	}

	ErrorDetailtracevisualisation struct {
		Source      string      `json:"Source"`
		ReasonCode  string      `json:"ReasonCode"`
		Description string      `json:"Description"`
		Recoverable bool        `json:"Recoverable"`
		Details     interface{} `json:"Details,omitempty"`
	}

	Responsetracevisualisation struct {
		Networkalertid string `json:"networkalertid"`
	}
	ErrorResponsestracevisualisation struct {
		Errors struct {
			Error []ErrorDetail `json:"Error"`
		} `json:"Errors"`
	}
)

type TraceFeedback struct {
	AlertID string `json:"alertID"`
	Status  bool   `json:"status"`
	// Add other fields as needed
}
