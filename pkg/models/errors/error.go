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
	Code        string `json:"reasoncode"`
	Description string `json:"description"`
	Service     string `json:"service"`
}
