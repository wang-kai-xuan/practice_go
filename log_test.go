package main

import (
	"log"
	"os"
	"testing"
	"time"
)

// TestLogExample TestLogExample
func TestLogExample(t *testing.T) {
	name := "./" + time.Now().Format(time.RFC3339)[:10] + ".log"
	logFile, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if nil != err {
		panic(err)
	}

	logger := log.New(logFile, "", log.Ltime|log.Lshortfile)
	logger.Println("wkx test")
}
