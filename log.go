package main

import (
	"fmt"
	"os"
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
	return []byte(fmt.Sprintf("\"%s\" \"%s\" \"%s\" %v\n", timestamp, f.LevelDesc[entry.Level], entry.Message, entry.Data)), nil
}

func NewLogger(prefix string) *logrus.Logger {
	if Log != nil {
		return Log
	}

	date := time.Now().Format(time.RFC3339)[:10]

	pathMap := lfshook.PathMap{
		logrus.ErrorLevel: date + "-error.log",
		logrus.FatalLevel: date + "-fatal.log",
	}

	Log = logrus.New()

	name := "./" + prefix + "-" + date + ".log"
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
