package Log

import (
	"fmt"
	"log"
	"os"
	"time"

	d "github.com/VidarSolutions/Lib/Dates"
)

// Define LogLvl type
type LogLvl int

// Define constants for logging levels
const (
	DEBUG LogLvl = iota
	INFO
	WARN
	ERROR
	FATAL
)

// String method to convert LogLvl to string for easier logging
func (lvl LogLvl) String() string {
	return [...]string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}[lvl]
}

type LogEntry struct {
	Message string    `json:"message"`
	Date    time.Time `json:"date"`
	LogLvl  LogLvl
}

type AppLog struct {
	AppName     string
	AuthCode    string
	LogDir      string
	LogFileName string
	DateFormat  string
	LogLvl      LogLvl
	LogEntry    []LogEntry
	logger      *os.File // Custom logger instance
}

func (appLog AppLog) Say(msg string) {
	fmt.Printf("Logging : %s ", msg)
}

func (appLog AppLog) Write(msg string, logLvl LogLvl, console bool) {
	if logLvl < appLog.LogLvl {
		return // Skip logging if the provided log level is lower than the application's log level
	}
	var logmsg string
	switch logLvl {
	case DEBUG:
		logmsg = fmt.Sprintf("DEBUG: %s", msg)
	case INFO:
		logmsg = fmt.Sprintf("INFO: %s", msg)
	case WARN:
		logmsg = fmt.Sprintf("WARN: %s", msg)
	case ERROR:
		logmsg = fmt.Sprintf("ERROR: %s", msg)
	case FATAL:
		logmsg = fmt.Sprintf("FATAL: %s", msg)
	default:
		logmsg = fmt.Sprintf("UNKNOWN: %s", msg)
	}
	appLog.logger.WriteString(fmt.Sprintf("%s : %s", d.GetDate(appLog.DateFormat), msg))

	date := d.GetDate(appLog.DateFormat)
	entry := LogEntry{Message: logmsg, Date: date, LogLvl: logLvl}
	if console {
		appLog.Say(entry.Print())
	}
	appLog.LogEntry = append(appLog.LogEntry, entry)
}

func (entry LogEntry) Print() string {
	return fmt.Sprintf("%s  :: %s ", entry.Date, entry.Message)
}
func GetLogger(fileName string) (*os.File, error) {

	log.Printf("Saving App Logs to File : %s\n", fileName)
	// Open log file
	fileLogger, flerr := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if flerr != nil {
		fmt.Println("Error opening log file:", flerr)
		return nil, flerr
	}
	fileLogger.WriteString("Saving Import Key Program Results To This File\n\n")
	return fileLogger, nil
}
