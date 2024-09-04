package logs

import (
	"log"
	"os"
)

// Error_Logger is a logger for error messages.
var Error_Logger = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime)

// Info_Logger is a logger for informational messages.
var Info_Logger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)

// Warning_Logger is a logger for warning messages.
var Warning_Logger = log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime)
