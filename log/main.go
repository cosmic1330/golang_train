package main

import (
	"fmt"
	"log"
	"os"
)

var (
	INFO *log.Logger
	WARNING *log.Logger
	ERROR *log.Logger
)

func main() {
	logfile, err := os.OpenFile("log/.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Failed to open log file")
		return
	}
	INFO = log.New(logfile, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WARNING = log.New(logfile, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	ERROR = log.New(logfile, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	INFO.Println("This is a info log")
	WARNING.Println("This is a warning log")
	ERROR.Println("This is a error log")

}
