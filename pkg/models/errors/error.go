package errors

type ErrorModel struct {
	Message   string `json:"message"`
	IsSuccess bool   `json:"isSuccess"`
	Error     error  `json:"error"`
}

type Errorresp struct {
	Errorresponse string `json:"error"`
}

type Errorcode struct {
	Id          int    `json:"id"`
	Source      string `json:"source"`
	Reasoncode  string `json:"reasoncode"`
	Description string `json:"description"`
	Recoverable bool   `json:"recoverable"`
	Details     string `json:"details"`
	Errorcode   string `json:"errorcode"`
}
