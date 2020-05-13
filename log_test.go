package main

import (
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/rifflock/lfshook"
	logrus "github.com/sirupsen/logrus"
)

var Log *logrus.Logger

type PlainFormatter struct {
	TimestampFormat string
	LevelDesc       []string
}

func (f *PlainFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	timestamp := fmt.Sprintf(entry.Time.Format(f.TimestampFormat))
	return []byte(fmt.Sprintf("%s %s \"%s\"\n", timestamp, f.LevelDesc[entry.Level], entry.Message)), nil
}

func NewLogger() *logrus.Logger {
	if Log != nil {
		return Log
	}

	date := "./" + time.Now().Format(time.RFC3339)[:10]

	pathMap := lfshook.PathMap{
		logrus.InfoLevel:  date + "-info.log",
		logrus.ErrorLevel: date + "-error.log",
	}

	Log = logrus.New()

	name := date + ".log"
	logFile, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if nil != err {
		panic(err)
	}
	plainFormatter := new(PlainFormatter)
	plainFormatter.TimestampFormat = "2006-01-02 15:04:05"
	plainFormatter.LevelDesc = []string{"PANC", "FATL", "ERRO", "WARN", "INFO", "DEBG"}
	Log.SetFormatter(plainFormatter)

	Log.SetLevel(logrus.TraceLevel)
	Log.SetOutput(logFile)
	Log.Hooks.Add(lfshook.NewHook(
		pathMap,
		plainFormatter,
	))
	return Log
}

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

func init() {
	// Log as JSON instead of the default ASCII formatter.
	logrus.SetFormatter(&logrus.TextFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	name := "./" + time.Now().Format(time.RFC3339)[:10] + ".log"
	logFile, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if nil != err {
		panic(err)
	}
	// Only log the warning severity or above.
	logrus.SetLevel(logrus.TraceLevel)
	logrus.SetOutput(logFile)
}

func TestLogrus(t *testing.T) {
	logrus.WithFields(logrus.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

	logrus.WithFields(logrus.Fields{
		"omg":    true,
		"number": 122,
	}).Warn("The group's number increased tremendously!")

	// A common pattern is to re-use fields between logging statements by re-using
	// the logrus.Entry returned from WithFields()
	contextLogger := logrus.WithFields(logrus.Fields{
		"common": "this is a common field",
		"other":  "I also should be logged always",
	})

	contextLogger.Info("I'll be logged with common and other field")
	contextLogger.Info("Me too")
	logrus.WithFields(logrus.Fields{
		"omg":    true,
		"number": 100,
	}).Fatal("The ice breaks!")
}

func TestLogrus1(t *testing.T) {
	Log = NewLogger()

	Log.WithFields(logrus.Fields{
		"omg": true,
	}).Info("nihao test")

	Log.WithFields(logrus.Fields{
		"omg": true,
	}).Debug("nihao test")

	Log.WithFields(logrus.Fields{
		"omg": true,
	}).Warn("nihao test")

	Log.WithFields(logrus.Fields{
		"omg": true,
	}).Error("nihao test")

	Log.WithFields(logrus.Fields{
		"omg": true,
	}).Fatal("nihao test")
}
