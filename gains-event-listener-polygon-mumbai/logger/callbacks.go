package logger

import (
	"log"
	"os"
)

// WriteToLog writes specific data to a log file named with logName
func WriteToLog(data, logName string) {
	// Specify the log file path
	logFilePath := "logs/" + logName + ".log"

	// Open the log file for appending. Create it if it doesn't exist.
	logFile, err := os.OpenFile(logFilePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal("Error opening log file:", err)
	}
	defer logFile.Close()

	// Set the log output to the file
	log.SetOutput(logFile)

	// Log the specific data
	log.Println(data)
}
