package lib

import (
	"log"
	"net/http"
	"os"
	"time"
)

// Function to create a log file and return a new Logger
func Logger() (*log.Logger, *os.File) {
	// Create logfile if it doesn't exist, set file attributes
	f, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	// Handle error
	if err != nil {
		log.Fatal("Cannot open file ", err)
	}
	// Create a new Logger and pass the file to it
	logger := log.New(f, "", 0)

	return logger, f
}

// Decorator function for http HandlerFunc logging
func Decorator(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		logger, logFile := Logger()
		logger.Println(t.Format(time.RFC3339), r.RemoteAddr, r.Method, r.RequestURI)
		logFile.Close()
		f(w, r) // Execute function passed to the Decorator
	}
}
