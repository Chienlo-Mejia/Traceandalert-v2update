package utils

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
	Separator     *log.Logger
)

func CreateDirectory(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, 0755)
		if err != nil {
			return err
		}
		return nil
	} else {
		return err
	}
}

func SystemLoggerAPI(url string, body interface{}, class string, resp *http.Response, ret interface{}, ip string) {

	currentTime := time.Now()
	file, err := os.OpenFile("logs/"+currentTime.Format("01022006")+"_API.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime)
	Separator = log.New(file, "", log.Ldate|log.Ltime)

	strBody := fmt.Sprintf("%#v", body)
	strRet := fmt.Sprintf("%#v", ret)
	strStatus := fmt.Sprintf("%v", resp.Status)

	Separator.Println("")
	InfoLogger.Println(class + ": - - - - - - - - - - - - - - -")
	InfoLogger.Println(class + ": FROM: " + ip)
	InfoLogger.Println(class + ": URL REQUEST: " + url)
	InfoLogger.Println(class + ": BODY: " + strBody)
	InfoLogger.Println(class + ": RESPONSE: " + strRet)
	InfoLogger.Println(class + ": STATUS: " + strStatus)
	file.Close()
}

func SystemLoggerErrorAPI(url string, body interface{}, class string, resp *http.Response, ret interface{}, ip string) {

	currentTime := time.Now()
	file, err := os.OpenFile("logs/"+currentTime.Format("01022006")+"_API.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	InfoLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime)
	Separator = log.New(file, "", log.Ldate|log.Ltime)

	strBody := fmt.Sprintf("%#v", body)
	strRet := fmt.Sprintf("%#v", ret)
	strStatus := fmt.Sprintf("%v", resp.Status)

	Separator.Println("")
	InfoLogger.Println(class + ": - - - - - - - - - - - - - - -")
	InfoLogger.Println(class + ": FROM: " + ip)
	InfoLogger.Println(class + ": URL REQUEST: " + url)
	InfoLogger.Println(class + ": BODY: " + strBody)
	InfoLogger.Println(class + ": RESPONSE: " + strRet)
	InfoLogger.Println(class + ": STATUS: " + strStatus)
	file.Close()
}

func SystemLoggerDB(body interface{}, class string, status int, ret string, ip string) {

	currentTime := time.Now()
	file, err := os.OpenFile("logs/"+currentTime.Format("01022006")+"_DB.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime)
	Separator = log.New(file, "", log.Ldate|log.Ltime)

	strBody := fmt.Sprintf("%#v", body)
	strRet := fmt.Sprintf("%#v", ret)

	Separator.Println("")
	InfoLogger.Println(class + ": - - - - - - - - - - - - - - -")
	InfoLogger.Println(class + ": FROM: " + ip)
	InfoLogger.Println(class + ": BODY: " + strBody)
	InfoLogger.Println(class + ": RESPONSE: " + strRet)
	InfoLogger.Println(class + ": STATUS: " + strconv.Itoa(status))
	file.Close()
}

func SystemLoggerErrorDB(body interface{}, class string, status int, ret string, ip string) {

	currentTime := time.Now()
	file, err := os.OpenFile("logs/"+currentTime.Format("01022006")+"_DB.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	ErrorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime)
	Separator = log.New(file, "", log.Ldate|log.Ltime)

	strBody := fmt.Sprintf("%#v", body)
	strRet := fmt.Sprintf("%#v", ret)

	Separator.Println("")
	ErrorLogger.Println(class + ": - - - - - - - - - - - - - - -")
	ErrorLogger.Println(class + ": FROM: " + ip)
	ErrorLogger.Println(class + ": BODY: " + strBody)
	ErrorLogger.Println(class + ": RESPONSE: " + strRet)
	ErrorLogger.Println(class + ": STATUS: " + strconv.Itoa(status))
	file.Close()
}

func SystemLoggerError(class string, proccess string, errorData error) {

	currentTime := time.Now()
	file, err := os.OpenFile("logs/"+currentTime.Format("01022006")+"_SYSTEM_ERROR.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	ErrorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime)
	Separator = log.New(file, "", log.Ldate|log.Ltime)

	strError := fmt.Sprintf("%#v", errorData.Error())

	Separator.Println("")
	ErrorLogger.Println(class + ": - - - - - - - - - - - - - - -")
	ErrorLogger.Println(class + ": ERROR AT: " + proccess)
	ErrorLogger.Println(class + ": ERROR: " + strError)
	file.Close()
}

func SystemLogger(class string, data string, username string) {

	currentTime := time.Now()
	file, err := os.OpenFile("logs/"+currentTime.Format("01022006")+"_SYSTEM_ERROR.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime)
	Separator = log.New(file, "", log.Ldate|log.Ltime)

	strError := fmt.Sprintf("%#v", username)

	Separator.Println("")
	InfoLogger.Println(class + ": - - - - - - - - - - - - - - -")
	InfoLogger.Println(class + ": UPDATED RECORD is_active column: value(" + data + ")")
	InfoLogger.Println(class + ": ERROR: " + strError)
	file.Close()
}

// feedback
func Feedbacklogs(class, folder, Feedbackid string, Errors string) {
	currentTime := time.Now()
	folderName := "./logs/feedback/" + folder + "/" + currentTime.Format("01-January")
	CreateDirectory(folderName)
	err := os.MkdirAll(folderName, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	logFilename := fmt.Sprintf("%s/alerts_feedback %s.log", folderName, currentTime.Format("2006-01-02"))

	// Open the log file in append mode.
	logfile, err := os.OpenFile(logFilename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Printf("Error opening log file: %v", err)
		return
	}
	defer func() {
		if closeErr := logfile.Close(); closeErr != nil {
			log.Printf("Error closing log file: %v", closeErr)
		}
	}()

	// Set the output of the log to the file.
	log.SetOutput(logfile)

	InfoLogger.Printf("")
	InfoLogger.Printf("               - - - - Feedback - - - -   ")

	InfoLogger.Printf("FEEDBACK: %+v\n", Feedbackid)
	InfoLogger.Printf("RESPONSE: %+v\n", Errors)
	InfoLogger.Printf("=========================================================================================")

	log.Printf("")
	log.Printf("               - - - - Feedback - - - -   ")
	log.Printf("FEEDBACK: %+v\n", Feedbackid)
	log.Printf("RESPONSE: %+v\n", Errors)
	log.Printf("=========================================================================================")

}

// alerttransaction
func Alertaccount(class, folder, ID string, message string, Networkalertid string, Accountid string, Networkid string, Owningbankid string, Owningbankname string, Time time.Time, Name string, Mulescore float64, Sourcetransactionvalue int, Endpointflag bool, Numoutboundrelationships int, Numinboundrelationships int, Numscheduledmandates int, Firstappearance time.Time, Mostrecentappearance time.Time, Receivessalary bool, Dwelltime string, Numnetworks int, Numtracednetworks int, Generation int, Tracetype string, Totalsuspiciousvalueinbound int, Totalsuspiciousvalueoutbound int, Totalvalueinbound int, Totalvalueoutbound int, Generations []int, Mostrecentfeedback string, Parentalertid string, Decisiondate time.Time, Nextpaginationtoken string, Previouspaginationtoken string) {
	currentTime := time.Now()
	folderName := "./logs/alert_account/" + folder + "/" + currentTime.Format("01-January")
	CreateDirectory(folderName)
	err := os.MkdirAll(folderName, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	logFilename := fmt.Sprintf("%s/alerts_account %s.log", folderName, currentTime.Format("2006-01-02"))

	// Open the log file in append mode.
	logfile, err := os.OpenFile(logFilename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Printf("Error opening log file: %v", err)
		return
	}
	defer func() {
		if closeErr := logfile.Close(); closeErr != nil {
			log.Printf("Error closing log file: %v", closeErr)
		}
	}()

	// Set the output of the log to the file.
	log.SetOutput(logfile)

	InfoLogger.Printf("")
	InfoLogger.Printf("               - - - - Alert-Transaction - - - -   ")
	InfoLogger.Printf("FEEDBACK: %+v\n", ID)
	InfoLogger.Printf("RESPONSE: %+v\n", message)
	InfoLogger.Printf("NETWORKALERTID: %+v\n", Networkalertid)
	InfoLogger.Printf("ACCOUNTID: %+v\n", Accountid)
	InfoLogger.Printf("NETWORKID: %+v\n", Networkid)
	InfoLogger.Printf("OWNINGBANKID: %+v\n", Owningbankid)
	InfoLogger.Printf("OWNINGBANKNAME: %+v\n", Owningbankname)
	InfoLogger.Printf("TIME: %+v\n", Time)
	InfoLogger.Printf("NAME: %+v\n", Name)
	InfoLogger.Printf("MULESCORE: %+v\n", Mulescore)
	InfoLogger.Printf("SOURCETRANSACTIONVALUE: %+v\n", Sourcetransactionvalue)
	InfoLogger.Printf("ENDPOINTFLAG: %+v\n", Endpointflag)
	InfoLogger.Printf("NUMOUTBOUNDRELATIONSHIPS: %+v\n", Numoutboundrelationships)
	InfoLogger.Printf("NUMINBOUNDRELATIONSHIPS: %+v\n", Numinboundrelationships)
	InfoLogger.Printf("NUMSCHEDULEDMANDATES: %+v\n", Numscheduledmandates)
	InfoLogger.Printf("FIRSTAPPEARANCE: %+v\n", Firstappearance)
	InfoLogger.Printf("MOSTRECENTAPPEARANCE: %+v\n", Mostrecentappearance)
	InfoLogger.Printf("RECEIVESSALARY: %+v\n", Receivessalary)
	InfoLogger.Printf("DWELLTIME: %+v\n", Dwelltime)
	InfoLogger.Printf("NUMNETWORKS: %+v\n", Numnetworks)
	InfoLogger.Printf("NUMTRACEDNETWORKS: %+v\n", Numtracednetworks)
	InfoLogger.Printf("GENERATION: %+v\n", Generation)
	InfoLogger.Printf("TRACETYPE: %+v\n", Tracetype)
	InfoLogger.Printf("TOTALSUSPICIOUSVALUEINBOUND: %+v\n", Totalsuspiciousvalueinbound)
	InfoLogger.Printf("TOTALSUSPICIOUSVALUEOUTBOUND: %+v\n", Totalsuspiciousvalueoutbound)
	InfoLogger.Printf("TOTALVALUEINBOUND: %+v\n", Totalvalueinbound)
	InfoLogger.Printf("TOTALVALUEOUTBOUND: %+v\n", Totalvalueoutbound)
	InfoLogger.Printf("GENERATIONS: %+v\n", Generations)
	InfoLogger.Printf("MOSTRECENTFEEDBACK: %+v\n", Mostrecentfeedback)
	InfoLogger.Printf("PARENTALERTID: %+v\n", Parentalertid)
	InfoLogger.Printf("NEXTPAGINATIONTOKEN: %+v\n", Nextpaginationtoken)
	InfoLogger.Printf("PREVIOUSPAGINATIONTOKEN: %+v\n", Previouspaginationtoken)
	InfoLogger.Printf("=========================================================================================")

	log.Printf("")
	log.Printf("               - - - - Alert-Transaction - - - -   ")
	log.Printf("FEEDBACK: %+v\n", ID)
	log.Printf("RESPONSE: %+v\n", message)
	log.Printf("NETWORKALERTID: %+v\n", Networkalertid)
	log.Printf("ACCOUNTID: %+v\n", Accountid)
	log.Printf("NETWORKID: %+v\n", Networkid)
	log.Printf("OWNINGBANKID: %+v\n", Owningbankid)
	log.Printf("OWNINGBANKNAME: %+v\n", Owningbankname)
	log.Printf("TIME: %+v\n", Time)
	log.Printf("NAME: %+v\n", Name)
	log.Printf("MULESCORE: %+v\n", Mulescore)
	log.Printf("SOURCETRANSACTIONVALUE: %+v\n", Sourcetransactionvalue)
	log.Printf("ENDPOINTFLAG: %+v\n", Endpointflag)
	log.Printf("NUMOUTBOUNDRELATIONSHIPS: %+v\n", Numoutboundrelationships)
	log.Printf("NUMINBOUNDRELATIONSHIPS: %+v\n", Numinboundrelationships)
	log.Printf("NUMSCHEDULEDMANDATES: %+v\n", Numscheduledmandates)
	log.Printf("FIRSTAPPEARANCE: %+v\n", Firstappearance)
	log.Printf("MOSTRECENTAPPEARANCE: %+v\n", Mostrecentappearance)
	log.Printf("RECEIVESSALARY: %+v\n", Receivessalary)
	log.Printf("DWELLTIME: %+v\n", Dwelltime)
	log.Printf("NUMNETWORKS: %+v\n", Numnetworks)
	log.Printf("NUMTRACEDNETWORKS: %+v\n", Numtracednetworks)
	log.Printf("GENERATION: %+v\n", Generation)
	log.Printf("TRACETYPE: %+v\n", Tracetype)
	log.Printf("TOTALSUSPICIOUSVALUEINBOUND: %+v\n", Totalsuspiciousvalueinbound)
	log.Printf("TOTALSUSPICIOUSVALUEOUTBOUND: %+v\n", Totalsuspiciousvalueoutbound)
	log.Printf("TOTALVALUEINBOUND: %+v\n", Totalvalueinbound)
	log.Printf("TOTALVALUEOUTBOUND: %+v\n", Totalvalueoutbound)
	log.Printf("GENERATIONS: %+v\n", Generations)
	log.Printf("MOSTRECENTFEEDBACK: %+v\n", Mostrecentfeedback)
	log.Printf("PARENTALERTID: %+v\n", Parentalertid)
	log.Printf("NEXTPAGINATIONTOKEN: %+v\n", Nextpaginationtoken)
	log.Printf("PREVIOUSPAGINATIONTOKEN: %+v\n", Previouspaginationtoken)

	log.Printf("=========================================================================================")

}

// alerttransaction
func Alerttransaction(class, folder, ID string, Errors string, count_no int64, Txnid string, Networkalertid string, Networkid string, Time time.Time, Txntime time.Time, Sourceid string, Destid string, Sourcebankid string, Sourcebankname string, Destbankid string, Destbankname string, Value int, Remitinfo string, Generation int, Currency string, Service string, Dwelltime string, Tracetype string, Mulescore float64, Parentalertid string, Decisiondate time.Time, Mostrecentfeedback string) {
	currentTime := time.Now()
	folderName := "./logs/alert_transaction/" + folder + "/" + currentTime.Format("01-January")
	CreateDirectory(folderName)
	err := os.MkdirAll(folderName, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	logFilename := fmt.Sprintf("%s/alert_transaction %s.log", folderName, currentTime.Format("2006-01-02"))

	// Open the log file in append mode.
	logfile, err := os.OpenFile(logFilename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Printf("Error opening log file: %v", err)
		return
	}
	defer func() {
		if closeErr := logfile.Close(); closeErr != nil {
			log.Printf("Error closing log file: %v", closeErr)
		}
	}()

	// Set the output of the log to the file.
	log.SetOutput(logfile)

	InfoLogger.Printf("")
	InfoLogger.Printf("               - - - - Alert-Transaction - - - -   ")
	InfoLogger.Printf("ID: %+v\n", ID)
	InfoLogger.Printf("RESPONSE: %+v\n", Errors)
	InfoLogger.Printf("COUNT: %+v\n", count_no)
	InfoLogger.Printf("TXNID: %+v\n", Txnid)
	InfoLogger.Printf("NETWORKALERTID: %+v\n", Networkalertid)
	InfoLogger.Printf("NETWORKID: %+v\n", Networkid)
	InfoLogger.Printf("Time: %+v\n", Time)
	InfoLogger.Printf("TXNTIME: %+v\n", Txntime)
	InfoLogger.Printf("SOURCEID: %+v\n", Sourceid)
	InfoLogger.Printf("DESTID: %+v\n", Destid)
	InfoLogger.Printf("SOURCEBANKID: %+v\n", Sourcebankid)
	InfoLogger.Printf("SOURCEBANKNAME: %+v\n", Sourcebankname)
	InfoLogger.Printf("DESTBANKID: %+v\n", Destbankid)
	InfoLogger.Printf("DESTBANKNAME: %+v\n", Destbankname)
	InfoLogger.Printf("VALUE: %+v\n", Value)
	InfoLogger.Printf("REMITINFO: %+v\n", Remitinfo)
	InfoLogger.Printf("GENERATION: %+v\n", Generation)
	InfoLogger.Printf("CURRENCY: %+v\n", Currency)
	InfoLogger.Printf("SERVICE: %+v\n", Service)
	InfoLogger.Printf("DWELLTIME: %+v\n", Dwelltime)
	InfoLogger.Printf("TRACETYPE: %+v\n", Tracetype)
	InfoLogger.Printf("MULESCORE: %+v\n", Mulescore)
	InfoLogger.Printf("PARENTALERTID: %+v\n", Parentalertid)
	InfoLogger.Printf("DECISIONDATE: %+v\n", Decisiondate)
	InfoLogger.Printf("MOSTRECENTFEEDBACK: %+v\n", Mostrecentfeedback)
	InfoLogger.Printf("DECISIONDATE: %+v\n", Decisiondate)
	InfoLogger.Printf("=========================================================================================")

	log.Printf("")
	log.Printf("               - - - -  Alert-Transaction - - - -   ")
	log.Printf("ID: %+v\n", ID)
	log.Printf("RESPONSE: %+v\n", Errors)
	log.Printf("COUNT: %+v\n", count_no)
	log.Printf("TXNID: %+v\n", Txnid)
	log.Printf("NETWORKALERTID: %+v\n", Networkalertid)
	log.Printf("NETWORKID: %+v\n", Networkid)
	log.Printf("TIME: %+v\n", Time)
	log.Printf("TXNTIME: %+v\n", Txntime)
	log.Printf("SOURCEID: %+v\n", Sourceid)
	log.Printf("DESTID: %+v\n", Destid)
	log.Printf("SOURCEBANKID: %+v\n", Sourcebankid)
	log.Printf("SOURCEBANKNAME: %+v\n", Sourcebankname)
	log.Printf("DESTBANKID: %+v\n", Destbankid)
	log.Printf("DESTBANKNAME: %+v\n", Destbankname)
	log.Printf("VALUE: %+v\n", Value)
	log.Printf("REMITINFO: %+v\n", Remitinfo)
	log.Printf("GENERATION: %+v\n", Generation)
	log.Printf("CURRENCY: %+v\n", Currency)
	log.Printf("SERVICE: %+v\n", Service)
	log.Printf("DWELLTIME: %+v\n", Dwelltime)
	log.Printf("TRACETYPE: %+v\n", Tracetype)
	log.Printf("MULESCORE: %+v\n", Mulescore)
	log.Printf("PARENTALERTID: %+v\n", Parentalertid)
	log.Printf("DECISIONDATE: %+v\n", Parentalertid)
	log.Printf("MOSTRECENTFEEDBACK: %+v\n", Mostrecentfeedback)
	log.Printf("DECISIONDATE: %+v\n", Decisiondate)
	log.Printf("=========================================================================================")
}

// tracenetwork
func Tracenetwork(class, folder, Vizurl string, Errors string, Sourcetxnid string, Sourcetxntype string, Length int, Generations int, Totalvalue int, Sourcevalue int, Uniqueaccounts int, Meandwelltime string, Mediandwelltime string, Meanmulescore float64, Elapsedtime string, Numnotinvestigated int, Parentalertid string, Decisiondate time.Time, Mostrecentfeedback string) {
	currentTime := time.Now()
	folderName := "./logs/trace_network/" + folder + "/" + currentTime.Format("01-January")
	CreateDirectory(folderName)
	err := os.MkdirAll(folderName, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	logFilename := fmt.Sprintf("%s/trace_network %s.log", folderName, currentTime.Format("2006-01-02"))

	// Open the log file in append mode.
	logfile, err := os.OpenFile(logFilename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Printf("Error opening log file: %v", err)
		return
	}
	defer func() {
		if closeErr := logfile.Close(); closeErr != nil {
			log.Printf("Error closing log file: %v", closeErr)
		}
	}()

	// Set the output of the log to the file.
	log.SetOutput(logfile)

	InfoLogger.Printf("")
	InfoLogger.Printf("               - - - - Trace-Network - - - -   ")
	InfoLogger.Printf("VIZURL: %+v\n", Vizurl)
	InfoLogger.Printf("RESPONSE: %+v\n", Errors)
	InfoLogger.Printf("SOURCETXNID: %+v\n", Sourcetxnid)
	InfoLogger.Printf("SOURCETXNTYPE: %+v\n", Sourcetxntype)
	InfoLogger.Printf("LENGTH: %+v\n", Length)
	InfoLogger.Printf("GENERATIONS: %+v\n", Generations)
	InfoLogger.Printf("TOTALVALUE: %+v\n", Totalvalue)
	InfoLogger.Printf("SOURCEVALUE: %+v\n", Sourcevalue)
	InfoLogger.Printf("UNIQUEACCOUNTS: %+v\n", Uniqueaccounts)
	InfoLogger.Printf("MEANDWELLTIME: %+v\n", Meandwelltime)
	InfoLogger.Printf("MEDIANDWELLTIME: %+v\n", Mediandwelltime)
	InfoLogger.Printf("MEANMULESCORE: %+v\n", Meanmulescore)
	InfoLogger.Printf("ELAPSEDTIME: %+v\n", Elapsedtime)
	InfoLogger.Printf("NUMNOTINVESTIGATED: %+v\n", Numnotinvestigated)
	InfoLogger.Printf("PARENTALERTID: %+v\n", Parentalertid)
	InfoLogger.Printf("DECISIONDATE: %+v\n", Decisiondate)
	InfoLogger.Printf("MOSTRECENTFEEDBACK: %+v\n", Mostrecentfeedback)
	InfoLogger.Printf("=========================================================================================")

	log.Printf("")
	log.Printf("               - - - - Trace-Network - - - -   ")
	log.Printf("VIZURL: %+v\n", Vizurl)
	log.Printf("VIZURL: %+v\n", Errors)
	log.Printf("SOURCETXNID: %+v\n", Sourcetxnid)
	log.Printf("SOURCETXNTYPE: %+v\n", Sourcetxntype)
	log.Printf("LENGTH: %+v\n", Length)
	log.Printf("GENERATIONS: %+v\n", Generations)
	log.Printf("TOTALVALUE: %+v\n", Totalvalue)
	log.Printf("SOURCEVALUE: %+v\n", Sourcevalue)
	log.Printf("UNIQUEACCOUNTS: %+v\n", Uniqueaccounts)
	log.Printf("MEANDWELLTIME: %+v\n", Meandwelltime)
	log.Printf("MEDIANDWELLTIME: %+v\n", Mediandwelltime)
	log.Printf("MEANMULESCORE: %+v\n", Meanmulescore)
	log.Printf("ELAPSEDTIME: %+v\n", Elapsedtime)
	log.Printf("NUMNOTINVESTIGATED: %+v\n", Numnotinvestigated)
	log.Printf("PARENTALERTID: %+v\n", Parentalertid)
	log.Printf("DECISIONDATE: %+v\n", Decisiondate)
	log.Printf("MOSTRECENTFEEDBACK: %+v\n", Mostrecentfeedback)
	log.Printf("=========================================================================================")

}

func Tracevisuallogs(class, folder, Traceid string, Errors string) {
	currentTime := time.Now()
	folderName := "./logs/trace_visualisation/" + folder + "/" + currentTime.Format("01-January")
	CreateDirectory(folderName)
	err := os.MkdirAll(folderName, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	logFilename := fmt.Sprintf("%s/trace_visualisation %s.log", folderName, currentTime.Format("2006-01-02"))

	// Open the log file in append mode.
	logfile, err := os.OpenFile(logFilename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Printf("Error opening log file: %v", err)
		return
	}
	defer func() {
		if closeErr := logfile.Close(); closeErr != nil {
			log.Printf("Error closing log file: %v", closeErr)
		}
	}()

	// Set the output of the log to the file.
	log.SetOutput(logfile)

	InfoLogger.Printf("")
	InfoLogger.Printf("               - - - - Trace-ID - - - -   ")
	InfoLogger.Printf("TRACEID: %+v\n", Traceid)
	InfoLogger.Printf("RESPONSE: %+v\n", Errors)
	InfoLogger.Printf("=========================================================================================")

	log.Printf("")
	log.Printf("               - - - - Trace-ID - - - -   ")
	log.Printf("TRACEID: %+v\n", Traceid)
	log.Printf("RESPONSE: %+v\n", Errors)
	log.Printf("=========================================================================================")

}
