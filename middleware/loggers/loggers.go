package loggers

import (
	"fmt"
	"log"
	"os"
	"time"
	"tracealert/pkg/utils/go-utils/database"
)

var (
	WarningLogger *log.Logger
	ErrorLogger   *log.Logger
	Separator     *log.Logger
)

var InfoLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

func OpenLogFile(path string) (*os.File, error) {
	logFile, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	return logFile, err
}

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

// feedback

func Feedbacklogs(class, folder, Uniqueidfeedback string, Feedbackid string, Errors string, Alertid string) {
	currentTime := time.Now()
	folderName := fmt.Sprintf("./logs/feedback/%s/%s", folder, currentTime.Format("01-January"))
	if err := os.MkdirAll(folderName, os.ModePerm); err != nil {
		log.Fatal(err)
	}

	logFilename := fmt.Sprintf("%s/alerts_feedback_%s.log", folderName, currentTime.Format("2006-01-02"))
	logfile, err := os.OpenFile(logFilename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Error opening log file: %v", err)
	}
	defer func() {
		if closeErr := logfile.Close(); closeErr != nil {
			log.Printf("Error closing log file: %v", closeErr)
		}
	}()

	InfoLogger := log.New(logfile, "INFO: ", log.Ldate|log.Ltime)
	Separator := log.New(logfile, "", log.Ldate|log.Ltime)

	Separator.Println("")
	InfoLogger.Println("               - - - - Feedback - - - -   ")
	InfoLogger.Printf("Unique-ID: %+v\n", Uniqueidfeedback)
	InfoLogger.Printf("FEEDBACK-REQUEST: %+v\n", Feedbackid)
	InfoLogger.Printf("ALERTID: %+v\n", Alertid)
	InfoLogger.Printf("RESPONSE: %+v\n", Errors)

	Separator.Println("")
	log.Printf("               - - - - Feedback - - - -   ")
	log.Printf("Unique-IDz: %+v\n", Uniqueidfeedback)
	log.Printf("FEEDBACK-REQUEST: %+v\n", Feedbackid)
	log.Printf("ALERTID: %+v\n", Alertid)
	log.Printf("RESPONSE: %+v\n", Errors)

	database.DBConn.Exec("SELECT * FROM trace_alerts.insert_logs_feedback(?,?,?, ? )", Uniqueidfeedback, Alertid, Feedbackid, Errors)
}

// alerttransaction
func Alertaccount(class, folder, Uniqueidalertaccount string, ID string, message string, Networkalertid string, Accountid string, Networkid string, Owningbankid string, Owningbankname string, Time time.Time, Name string, Mulescore float64, Sourcetransactionvalue int, Endpointflag bool, Numoutboundrelationships int, Numinboundrelationships int, Numscheduledmandates int, Firstappearance time.Time, Mostrecentappearance time.Time, Receivessalary bool, Dwelltime string, Numnetworks int, Numtracednetworks int, Generation int, Tracetype string, Totalsuspiciousvalueinbound int, Totalsuspiciousvalueoutbound int, Totalvalueinbound int, Totalvalueoutbound int, Generations []int, Mostrecentfeedback string, Parentalertid string, Decisiondate time.Time, Nextpaginationtoken string, Previouspaginationtoken string) {
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
	InfoLogger := log.New(logfile, "INFO: ", log.Ldate|log.Ltime)
	Separator := log.New(logfile, "", log.Ldate|log.Ltime)

	Separator.Println("")
	InfoLogger.Printf("               - - - - Alert-Transaction - - - -   ")
	InfoLogger.Printf("Unique-ID: %+v\n", Uniqueidalertaccount)
	InfoLogger.Printf("ALERTACCOUNT-REQUEST: %+v\n", ID)
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
	InfoLogger.Printf("DECISIONDATE: %+v\n", Decisiondate)
	InfoLogger.Printf("NEXTPAGINATIONTOKEN: %+v\n", Nextpaginationtoken)
	InfoLogger.Printf("PREVIOUSPAGINATIONTOKEN: %+v\n", Previouspaginationtoken)

	Separator.Println("")
	log.Printf("               - - - - Alert-Transaction - - - -   ")
	log.Printf("Unique-ID: %+v\n", Uniqueidalertaccount)
	log.Printf("ALERTACCOUNT-REQUEST: %+v\n", ID)
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
	log.Printf("DECISIONDATE: %+v\n", Decisiondate)
	log.Printf("NEXTPAGINATIONTOKEN: %+v\n", Nextpaginationtoken)
	log.Printf("PREVIOUSPAGINATIONTOKEN: %+v\n", Previouspaginationtoken)

	database.DBConn.Exec("SELECT * FROM trace_alerts.insert_logsalertaccount(?,?, ?,?,?,?,?,?,?,?)", Uniqueidalertaccount, ID, message, Networkalertid, Accountid, Networkid, Owningbankid, Owningbankname, Decisiondate, Parentalertid)
}

// alerttransaction
func Alerttransaction(class, folder, Uniqueidalertaccount string, ID string, Errors string, count_no int64, Txnid string, Networkalertid string, Networkid string, Time time.Time, Txntime time.Time, Sourceid string, Destid string, Sourcebankid string, Sourcebankname string, Destbankid string, Destbankname string, Value int, Remitinfo string, Generation int, Currency string, Service string, Dwelltime string, Tracetype string, Mulescore float64, Parentalertid string, Decisiondate time.Time, Mostrecentfeedback string) {
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
	InfoLogger := log.New(logfile, "INFO: ", log.Ldate|log.Ltime)
	Separator := log.New(logfile, "", log.Ldate|log.Ltime)

	Separator.Println("")
	InfoLogger.Printf("               - - - - Alert-Transaction - - - -   ")
	InfoLogger.Printf("Unique-ID: %+v\n", Uniqueidalertaccount)
	InfoLogger.Printf("ID-: %+v\n", ID)
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
	InfoLogger.Printf("MOSTRECENTFEEDBACK: %+v\n", Mostrecentfeedback)
	InfoLogger.Printf("DECISIONDATE: %+v\n", Decisiondate)

	Separator.Println("")
	log.Printf("               - - - -  Alert-Transaction - - - -   ")
	log.Printf("Unique-ID: %+v\n", Uniqueidalertaccount)
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
	log.Printf("MOSTRECENTFEEDBACK: %+v\n", Mostrecentfeedback)
	log.Printf("DECISIONDATE: %+v\n", Decisiondate)

	database.DBConn.Exec("SELECT * FROM trace_alerts.insert_log_alerttransaction(?,?,?,?,?,?, ?,?,?,?,?,?,?,?)", Uniqueidalertaccount, ID, Errors, Txnid, Networkalertid, Networkid, Sourceid, Destid, Sourcebankid, Sourcebankname, Destbankid, Destbankname, Parentalertid, Decisiondate)
}

func Tracenetwork(class, folder, Uniqueidnetwork string, Txnid_RB string, Vizurl string, Errors string, Sourcetxnid string, Sourcetxntype string, Parentalertid string, Decisiondate time.Time, Networkalertid string, Id string, Networkalertid_traceacc string, Accountid string, Networkid string, Owningbankid string, Owningbankname string, Time time.Time, ID_alert string, Txnid string, Networkalertid_alert string, Networkid_alert string, Sourceid string, Destid string, Sourcebankid string, Sourcebankname string, Destbankid string, Destbankname string, Tracetype string, Parentalertid_alert string, Decisiondate_alert time.Time, Status bool, Nextpaginationtoken string, Previouspaginationtoken string) {
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
	InfoLogger := log.New(logfile, "INFO: ", log.Ldate|log.Ltime)
	Separator := log.New(logfile, "", log.Ldate|log.Ltime)

	Separator.Println("")
	InfoLogger.Printf("               - - - - Trace-Network - - - -   ")
	InfoLogger.Printf("Unique-ID: %+v\n", Uniqueidnetwork)
	InfoLogger.Printf("--- ALERT")
	InfoLogger.Printf("TXNID: %+v\n", Txnid_RB)
	InfoLogger.Printf("ID: %+v\n", ID_alert)
	InfoLogger.Printf("TXNID: %+v\n", Txnid)
	InfoLogger.Printf("NETWORKALERT: %+v\n", Networkalertid_alert)
	InfoLogger.Printf("NETWORKID: %+v\n", Networkid_alert)
	InfoLogger.Printf("SOURCEID: %+v\n", Sourceid)
	InfoLogger.Printf("DESTID: %+v\n", Destid)
	InfoLogger.Printf("SOURCEBANKID: %+v\n", Sourcebankid)
	InfoLogger.Printf("SOURCEBANKNAME: %+v\n", Sourcebankname)
	InfoLogger.Printf("DESTBANKID: %+v\n", Destbankid)
	InfoLogger.Printf("DESTBANKNAME: %+v\n", Destbankname)
	InfoLogger.Printf("TRACETYPE: %+v\n", Tracetype)
	InfoLogger.Printf("PARENTID: %+v\n", Parentalertid_alert)
	InfoLogger.Printf("DECISIONDATE: %+v\n", Decisiondate_alert)
	InfoLogger.Printf("STATUS: %+v\n", Status)
	InfoLogger.Printf("NEXTPAGINATIONTOKEN: %+v\n", Nextpaginationtoken)
	InfoLogger.Printf("PREVIOUSPAGINATIONTOKEN: %+v\n", Previouspaginationtoken)
	InfoLogger.Printf("--- TRACE_ACC_ALERT")
	InfoLogger.Printf("ID: %+v\n", Id)
	InfoLogger.Printf("NETWORKALERTID: %+v\n", Networkalertid_traceacc)
	InfoLogger.Printf("ACCOUNTID: %+v\n", Accountid)
	InfoLogger.Printf("NETWORKID: %+v\n", Networkid)
	InfoLogger.Printf("OWNINGBANKID: %+v\n", Owningbankid)
	InfoLogger.Printf("OWNINGBANKNAME: %+v\n", Owningbankname)
	InfoLogger.Printf("TIME: %+v\n", Time)
	InfoLogger.Printf("--- TRACE_VIZURL ")
	InfoLogger.Printf("VIZURL: %+v\n", Vizurl)
	InfoLogger.Printf("RESPONSE: %+v\n", Errors)
	InfoLogger.Printf("SOURCETXNID: %+v\n", Sourcetxnid)
	InfoLogger.Printf("SOURCETXNTYPE: %+v\n", Sourcetxntype)
	InfoLogger.Printf("PARENTALERTID: %+v\n", Parentalertid)
	InfoLogger.Printf("DECISIONDATE: %+v\n", Decisiondate)
	InfoLogger.Printf("NETWORKALERTID: %+v\n", Networkalertid)

	Separator.Println("")
	log.Printf("               - - - - Trace-Network - - - -   ")
	log.Printf("Unique-ID: %+v\n", Uniqueidnetwork)
	log.Printf("--- ALERT")
	log.Printf("TXNID: %+v\n", Txnid_RB)
	log.Printf("RESPONSE: %+v\n", Errors)
	log.Printf("ID: %+v\n", ID_alert)
	log.Printf("TXNID: %+v\n", Txnid)
	log.Printf("NETWORKALERT: %+v\n", Networkalertid_alert)
	log.Printf("NETWORKID: %+v\n", Networkid_alert)
	log.Printf("SOURCEID: %+v\n", Sourceid)
	log.Printf("DESTID: %+v\n", Destid)
	log.Printf("SOURCEBANKID: %+v\n", Sourcebankid)
	log.Printf("SOURCEBANKNAME: %+v\n", Sourcebankname)
	log.Printf("DESTBANKID: %+v\n", Destbankid)
	log.Printf("DESTBANKNAME: %+v\n", Destbankname)
	log.Printf("TRACETYPE: %+v\n", Tracetype)
	log.Printf("PARENTID: %+v\n", Parentalertid_alert)
	log.Printf("DECISIONDATE: %+v\n", Decisiondate_alert)
	log.Printf("STATUS: %+v\n", Status)
	log.Printf("NEXTPAGINATIONTOKEN: %+v\n", Nextpaginationtoken)
	log.Printf("PREVIOUSPAGINATIONTOKEN: %+v\n", Previouspaginationtoken)
	log.Printf("--- TRACE_ACC_ALERT")
	log.Printf("ID: %+v\n", Id)
	log.Printf("NETWORKALERTID_TRACEACC: %+v\n", Networkalertid_traceacc)
	log.Printf("ACCOUNTID: %+v\n", Accountid)
	log.Printf("NETWORKID: %+v\n", Networkid)
	log.Printf("OWNINGBANKID: %+v\n", Owningbankid)
	log.Printf("OWNINGBANKNAME: %+v\n", Owningbankname)
	log.Printf("TIME: %+v\n", Time)
	log.Printf("--- TRACE_VIZURL ")
	log.Printf("VIZURL: %+v\n", Vizurl)
	log.Printf("SOURCETXNID: %+v\n", Sourcetxnid)
	log.Printf("SOURCETXNTYPE: %+v\n", Sourcetxntype)
	log.Printf("PARENTALERTID: %+v\n", Parentalertid)
	log.Printf("DECISIONDATE: %+v\n", Decisiondate)
	log.Printf("MOSTRECENTFEEDBACK: %+v\n", Networkalertid)

	database.DBConn.Exec("SELECT * FROM trace_alerts.insert_logs_tracenetwork(?,?,?,?,?,?,?)", Uniqueidnetwork, Txnid_RB, Txnid, Networkid, Sourcetxnid, Decisiondate, Errors)
}

func Tracevisuallogs(class, folder, Uniqueidvisualisation string, Traceid string, Errors string, Networkalertid string) {
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
	InfoLogger := log.New(logfile, "INFO: ", log.Ldate|log.Ltime)
	Separator := log.New(logfile, "", log.Ldate|log.Ltime)

	Separator.Println("")
	InfoLogger.Printf("               - - - - Trace-ID - - - -   ")
	InfoLogger.Printf("Unique-ID: %+v\n", Uniqueidvisualisation)
	InfoLogger.Printf("TRACEID: %+v\n", Traceid)
	InfoLogger.Printf("NETWORKALERTID: %+v\n", Networkalertid)
	InfoLogger.Printf("RESPONSE: %+v\n", Errors)

	Separator.Println("")
	log.Printf("               - - - - Trace-ID - - - -   ")
	log.Printf("Unique-ID: %+v\n", Uniqueidvisualisation)
	log.Printf("TRACEID: %+v\n", Traceid)
	log.Printf("NETWORKALERTID: %+v\n", Networkalertid)
	log.Printf("RESPONSE: %+v\n", Errors)

	database.DBConn.Exec("SELECT * FROM trace_alerts.insert_logs_tracevisualisation(?,?,?,? )", Uniqueidvisualisation, Networkalertid, Traceid, Errors)

}

// Creditransferalertslogs
func CreditransferAlertsLogs(class, folder, UniqueidCredittransfer string, message string, InstructionId string, requestTrigger, Transaction_type string, Status string, ReasonCode string, LocalInstrument string, ReferenceId string, SenderBic string, SenderName string, SenderAccount string, Amount float64, Currency string, ReceivingBic string, ReceivingName string, ReceivingAccount string, TraceAlert string, SourceTxnType string, AlertType string) {
	currentTime := time.Now()
	folderName := fmt.Sprintf("./logs/credit_transfer_alert/%s/%s", folder, currentTime.Format("01-January"))
	if err := os.MkdirAll(folderName, os.ModePerm); err != nil {
		log.Fatal(err)
	}

	logFilename := fmt.Sprintf("%s/credit_transfer_alert%s.log", folderName, currentTime.Format("2006-01-02"))
	logfile, err := os.OpenFile(logFilename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Error opening log file: %v", err)
	}
	defer func() {
		if closeErr := logfile.Close(); closeErr != nil {
			log.Printf("Error closing log file: %v", closeErr)
		}
	}()

	InfoLogger := log.New(logfile, "INFO: ", log.Ldate|log.Ltime)
	Separator := log.New(logfile, "", log.Ldate|log.Ltime)

	Separator.Println("")
	InfoLogger.Println("               - - - - Trace_Credittransfer - - - -   ")
	InfoLogger.Printf("Unique-ID: %+v\n", UniqueidCredittransfer)
	InfoLogger.Printf("DECISIONDATE: %+v\n", requestTrigger)
	InfoLogger.Printf("RESPONSE: %+v\n", message)
	InfoLogger.Printf("INSTRUCTION_ID: %+v\n", InstructionId)
	InfoLogger.Printf("TRANSACTIONTYPE: %+v\n", Transaction_type)
	InfoLogger.Printf("STATUS: %+v\n", Status)
	InfoLogger.Printf("TRACE_ALERT: %+v\n", TraceAlert)
	InfoLogger.Printf("TRACE_TYPE: %+v\n", SourceTxnType)
	InfoLogger.Printf("ALERT_TYPE: %+v\n", AlertType)
	InfoLogger.Printf("REASONCODE: %+v\n", ReasonCode)
	InfoLogger.Printf("LOCALINSTRUMENT: %+v\n", LocalInstrument)
	InfoLogger.Printf("REFERENCEID: %+v\n", ReferenceId)
	InfoLogger.Printf("SENDERBIC: %+v\n", SenderBic)
	InfoLogger.Printf("SENDERNAME: %+v\n", SenderName)
	InfoLogger.Printf("SENDERACCOUNT: %+v\n", SenderAccount)
	InfoLogger.Printf("AMOUNT: %+v\n", Amount)
	InfoLogger.Printf("AMOUNTCURRENCY: %+v\n", Currency)
	InfoLogger.Printf("RECEIVINGBIC: %+v\n", ReceivingBic)
	InfoLogger.Printf("RECEIVINGNAME: %+v\n", ReceivingName)
	InfoLogger.Printf("RECEIVINGACCOUNT: %+v\n", ReceivingAccount)

	Separator.Println("")
	log.Printf("               - - - - Trace_Credittransfer - - - -   ")
	log.Printf("Unique-IDz: %+v\n", UniqueidCredittransfer)
	log.Printf("DECISIONDATE: %+v\n", requestTrigger)
	log.Printf("RESPONSE: %+v\n", message)
	log.Printf("INSTRUCTION_ID: %+v\n", InstructionId)
	log.Printf("TRANSACTIONTYPE: %+v\n", Transaction_type)
	log.Printf("STATUS: %+v\n", Status)
	log.Printf("TRACE_ALERT: %+v\n", TraceAlert)
	log.Printf("TRACE_TYPE: %+v\n", SourceTxnType)
	log.Printf("ALERT_TYPE: %+v\n", AlertType)
	log.Printf("REASONCODE: %+v\n", ReasonCode)
	log.Printf("LOCALINSTRUMENT: %+v\n", LocalInstrument)
	log.Printf("REFERENCEID: %+v\n", ReferenceId)
	log.Printf("SENDERBIC: %+v\n", SenderBic)
	log.Printf("SENDERNAME: %+v\n", SenderName)
	log.Printf("SENDERACCOUNT: %+v\n", SenderAccount)
	log.Printf("AMOUNT: %+v\n", Amount)
	log.Printf("AMOUNTCURRENCY: %+v\n", Currency)
	log.Printf("RECEIVINGBIC: %+v\n", ReceivingBic)
	log.Printf("RECEIVINGNAME: %+v\n", ReceivingName)
	log.Printf("RECEIVINGACCOUNT: %+v\n", ReceivingAccount)
	database.DBConn.Exec("SELECT * FROM trace_alerts.insert_logs_credittransfer(?,?,?,?,?,?,?,?,?,?,?,?)", UniqueidCredittransfer, message, TraceAlert, SourceTxnType, Transaction_type, Status, ReasonCode, LocalInstrument, InstructionId, ReferenceId, requestTrigger, AlertType)

}

func CreditransferFeedbackLogs(class, folder, UniqueidCredittransfer string, message string, InstructionId string, requestTrigger string, TransactionType string, Status string, ReasonCode string, LocalInstrument string, ReferenceId string, TraceAlert string, SourceTxnType string, AlertType string, FeedBack string) {
	currentTime := time.Now()
	folderName := fmt.Sprintf("./logs/Feedback_credit_transfer_alert/%s/%s", folder, currentTime.Format("01-January"))
	if err := os.MkdirAll(folderName, os.ModePerm); err != nil {
		log.Fatal(err)
	}

	logFilename := fmt.Sprintf("%s/Feedback_credit_transfer_alert%s.log", folderName, currentTime.Format("2006-01-02"))
	logfile, err := os.OpenFile(logFilename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Error opening log file: %v", err)
	}
	defer func() {
		if closeErr := logfile.Close(); closeErr != nil {
			log.Printf("Error closing log file: %v", closeErr)
		}
	}()

	InfoLogger := log.New(logfile, "INFO: ", log.Ldate|log.Ltime)
	Separator := log.New(logfile, "", log.Ldate|log.Ltime)

	Separator.Println("")
	InfoLogger.Println("               - - - - Trace_Credittransfer - - - -   ")
	InfoLogger.Printf("Unique-ID: %+v\n", UniqueidCredittransfer)
	InfoLogger.Printf("DECISIONDATE: %+v\n", requestTrigger)
	InfoLogger.Printf("RESPONSE: %+v\n", message)
	InfoLogger.Printf("INSTRUCTION_ID: %+v\n", InstructionId)
	InfoLogger.Printf("TRANSACTIONTYPE: %+v\n", TransactionType)
	InfoLogger.Printf("STATUS: %+v\n", Status)
	InfoLogger.Printf("TRACE_ALERT: %+v\n", TraceAlert)
	InfoLogger.Printf("TRACE_TYPE: %+v\n", SourceTxnType)
	InfoLogger.Printf("ALERT_TYPE: %+v\n", AlertType)
	InfoLogger.Printf("FEEDBACK: %+v\n", FeedBack)
	InfoLogger.Printf("REASONCODE: %+v\n", ReasonCode)
	InfoLogger.Printf("LOCALINSTRUMENT: %+v\n", LocalInstrument)
	InfoLogger.Printf("REFERENCEID: %+v\n", ReferenceId)

	Separator.Println("")
	log.Printf("               - - - - Trace_Credittransfer - - - -   ")
	log.Printf("Unique-ID: %+v\n", UniqueidCredittransfer)
	log.Printf("DECISIONDATE: %+v\n", requestTrigger)
	log.Printf("RESPONSE: %+v\n", message)
	log.Printf("INSTRUCTION_ID: %+v\n", InstructionId)
	log.Printf("TRANSACTIONTYPE: %+v\n", TransactionType)
	log.Printf("STATUS: %+v\n", Status)
	log.Printf("TRACE_ALERT: %+v\n", TraceAlert)
	log.Printf("TRACE_TYPE: %+v\n", SourceTxnType)
	log.Printf("ALERT_TYPE: %+v\n", AlertType)
	log.Printf("FEEDBACK: %+v\n", FeedBack)
	log.Printf("REASONCODE: %+v\n", ReasonCode)
	log.Printf("LOCALINSTRUMENT: %+v\n", LocalInstrument)
	log.Printf("REFERENCEID: %+v\n", ReferenceId)

	database.DBConn.Exec("SELECT * FROM trace_alerts.insert_logsfeedback_credittransfer(?,?,?,?,?,?,?,?,?,?,?,?,?)", UniqueidCredittransfer, message, TraceAlert, SourceTxnType, TransactionType, Status, ReasonCode, LocalInstrument, InstructionId, ReferenceId, requestTrigger, AlertType, FeedBack)

}

//Detectpostmanlogs

func Detectpostmanlogs(class, folder, message string, requestTrigger time.Time) {
	currentTime := time.Now()
	folderName := fmt.Sprintf("./logs/DetectPostman/%s/%s", folder, currentTime.Format("01-January"))
	if err := os.MkdirAll(folderName, os.ModePerm); err != nil {
		log.Fatal(err)
	}

	logFilename := fmt.Sprintf("%s/DetectPostman%s.log", folderName, currentTime.Format("2006-01-02"))
	logfile, err := os.OpenFile(logFilename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Error opening log file: %v", err)
	}
	defer func() {
		if closeErr := logfile.Close(); closeErr != nil {
			log.Printf("Error closing log file: %v", closeErr)
		}
	}()

	InfoLogger := log.New(logfile, "INFO: ", log.Ldate|log.Ltime)
	Separator := log.New(logfile, "", log.Ldate|log.Ltime)

	Separator.Println("")
	InfoLogger.Println("               - - - - Trace_Transaction_(POSTMAN,MOBILE_PHONE,BROWSER) - - - -   ")
	InfoLogger.Printf("Message: %+v\n", message)
	InfoLogger.Printf("Date_Time: %+v\n", requestTrigger)

	Separator.Println("")
	log.Printf("               - - - - Trace_Transaction_(POSTMAN,MOBILE_PHONE,BROWSER) - - - -   ")
	log.Printf("message: %+v\n", message)
	log.Printf("Date_Time: %+v\n", requestTrigger)

	database.DBConn.Exec("SELECT * FROM trace_alerts.insert_logs_tracetrans_PostmanMobilephoneBrowser(?,? )", message, requestTrigger)
}

// // creditransferfeedbacklogs
// func Creditransferfeedbacklogs(class, folder, Uniqueidcredittransfer string, message string, Instructionid string, requestTrigger, Transactiontype string, Status string, Reasoncode string, Description string, Localinstrument string, Referenceid string, Senderbic string, Sendername string, Senderaccount string, Amountcurrency string, Senderamount float64, Receivingbic string, Receivingname string, Receivingaccount string, Trace_alert string, Sourcetxntype string) {
// 	currentTime := time.Now()
// 	folderName := "./logs/credit_transfer_trace/" + folder + "/" + currentTime.Format("01-January")
// 	CreateDirectory(folderName)
// 	err := os.MkdirAll(folderName, os.ModePerm)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	logFilename := fmt.Sprintf("%s/alerts_credittransfer_trace %s.log", folderName, currentTime.Format("2006-01-02"))

// 	// Open the log file in append mode.
// 	logfile, err := os.OpenFile(logFilename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
// 	if err != nil {
// 		log.Printf("Error opening log file: %v", err)
// 		return
// 	}
// 	defer func() {
// 		if closeErr := logfile.Close(); closeErr != nil {
// 			log.Printf("Error closing log file: %v", closeErr)
// 		}
// 	}()

// 	// Set the output of the log to the file.
// 	InfoLogger := log.New(logfile, "INFO: ", log.Ldate|log.Ltime)
// 	Separator := log.New(logfile, "", log.Ldate|log.Ltime)

// 	Separator.Println("")
// 	InfoLogger.Printf("               - - - - Alert_Credittransfer - - - -   ")
// 	InfoLogger.Printf("Unique-ID: %+v\n", Uniqueidcredittransfer)
// 	InfoLogger.Printf("DECISIONDATE: %+v\n", requestTrigger)
// 	InfoLogger.Printf("RESPONSE: %+v\n", message)
// 	// InfoLogger.Printf("LIMIT: %+v\n", Limit)
// 	// InfoLogger.Printf("TOTALRECORDS: %+v\n", count)
// 	InfoLogger.Printf("TRACE_ALERT: %+v\n", Trace_alert)
// 	InfoLogger.Printf("TRACE_TYPE: %+v\n", Sourcetxntype)
// 	InfoLogger.Printf("INSTRUCTION_ID: %+v\n", Instructionid)
// 	InfoLogger.Printf("TRANSACTIONTYPE: %+v\n", Transactiontype)
// 	InfoLogger.Printf("STATUS: %+v\n", Status)
// 	InfoLogger.Printf("REASONCODE: %+v\n", Reasoncode)
// 	InfoLogger.Printf("DESCRIPTION: %+v\n", Description)
// 	InfoLogger.Printf("LOCALINSTRUMENT: %+v\n", Localinstrument)
// 	InfoLogger.Printf("REFERENCEID: %+v\n", Referenceid)
// 	InfoLogger.Printf("SENDERBIC: %+v\n", Senderbic)
// 	InfoLogger.Printf("SENDERNAME: %+v\n", Sendername)
// 	InfoLogger.Printf("SENDERACCOUNT: %+v\n", Senderaccount)
// 	InfoLogger.Printf("AMOUNTCURRENCY: %+v\n", Amountcurrency)
// 	InfoLogger.Printf("SENDERAMOUNT: %+v\n", Senderamount)
// 	InfoLogger.Printf("RECEIVINGBIC: %+v\n", Receivingbic)
// 	InfoLogger.Printf("RECEIVINGNAME: %+v\n", Receivingname)
// 	InfoLogger.Printf("RECEIVINGACCOUNT: %+v\n", Receivingaccount)
// 	InfoLogger.Printf("=========================================================================================")

// 	Separator.Println("")
// 	log.Printf("               - - - - Alert_Credittransfer - - - -   ")
// 	log.Printf("Unique-ID: %+v\n", Uniqueidcredittransfer)
// 	log.Printf("DECISIONDATE: %+v\n", requestTrigger)
// 	log.Printf("RESPONSE: %+v\n", message)
// 	// log.Printf("LIMIT: %+v\n", Limit)
// 	// log.Printf("TOTALRECORDS: %+v\n", count)
// 	log.Printf("TRACE_ALERT: %+v\n", Trace_alert)
// 	log.Printf("TRACE_TYPE: %+v\n", Sourcetxntype)
// 	log.Printf("INSTRUCTION_ID: %+v\n", Instructionid)
// 	log.Printf("TRANSACTIONTYPE: %+v\n", Transactiontype)
// 	log.Printf("STATUS: %+v\n", Status)
// 	log.Printf("REASONCODE: %+v\n", Reasoncode)
// 	log.Printf("DESCRIPTION: %+v\n", Description)
// 	log.Printf("LOCALINSTRUMENT: %+v\n", Localinstrument)
// 	log.Printf("REFERENCEID: %+v\n", Referenceid)
// 	log.Printf("SENDERBIC: %+v\n", Senderbic)
// 	log.Printf("SENDERNAME: %+v\n", Sendername)
// 	log.Printf("SENDERACCOUNT: %+v\n", Senderaccount)
// 	log.Printf("AMOUNTCURRENCY: %+v\n", Amountcurrency)
// 	log.Printf("SENDERAMOUNT: %+v\n", Senderamount)
// 	log.Printf("RECEIVINGBIC: %+v\n", Receivingbic)
// 	log.Printf("RECEIVINGNAME: %+v\n", Receivingname)
// 	log.Printf("RECEIVINGACCOUNT: %+v\n", Receivingaccount)
// 	log.Printf("=========================================================================================")
// 	//Uniqueidcredittransfer, message, Tracealert ,trace_type ,transaction status reasoncode localinstrument, instructionid,reference ,datetime

// 	database.DBConn.Exec("SELECT * FROM trace_alerts.insert_logs_credittransfer(?,?,?,?,?,?,?,?,?,?,?)", Uniqueidcredittransfer, message, Trace_alert, Sourcetxntype, Transactiontype, Status, Reasoncode, Localinstrument, Instructionid, Referenceid, requestTrigger)

// }
