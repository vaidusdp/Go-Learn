package main

import (
	// "fmt"
	"log"
	"os"
)

func main() {
	// log.Println("This is a log msg")

	// log.SetPrefix("Info: ")
	// log.Println("This is a info msg")

	// // Flags
	// log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	// log.Println("This is a date only log")

	infoLogger.Println("INFO MSG")
	warnLogger.Println("WARN MSG")
	errorLogger.Println("ERROR MSG")

	file, err := os.OpenFile("app.log", os.O_CREATE | os.O_WRONLY | os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file error")
	}

	defer file.Close()

	debugLogger := log.New(file, "Debug: ", log.Ldate|log.Ltime|log.Lshortfile)
	debugLogger.Println("This is a debug message")

	warnLogger1 := log.New(file, "Warn: ", log.Ldate|log.Ltime|log.Lshortfile)
	infoLogger1  := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger1  := log.New(os.Stdout, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile)

	infoLogger1.Println("INFO MSG")
	warnLogger1.Println("WARN MSG")
	errorLogger1.Println("ERROR MSG")
	debugLogger.Println("Debug Msg")
}

var (
	infoLogger  = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	warnLogger  = log.New(os.Stdout, "WARN: ", log.Ldate|log.Ltime)
	errorLogger = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
)