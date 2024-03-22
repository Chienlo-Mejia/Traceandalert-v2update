package tracenetwork

type (
	Trans_Postman struct {
		StartDate string `json:"start_date"`
		EndDate   string `json:"end_date"`
	}

	Trans_response struct {
		Message        string `json:"message"`
		Requesttrigger string `json:"requesttrigger"`
	}
)
