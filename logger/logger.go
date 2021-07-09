package logger

import (
	"fmt"
	"log"
	"os"
	"time"
)

const (
	TimeFormat = "01-02-2006 15:04:05.0000"
	LogFile    = "logs/logs.log"
)

var logFile *log.Logger

func init() {
	outFile, err := os.OpenFile(LogFile, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0775)
	if err != nil {
		panic(err.Error())
	}

	logFile = log.New(outFile, "", 0)
}

func ForInfo(v ...interface{}) {
	str := fmt.Sprintf("%s |INFO|: %v", time.Now().Format(TimeFormat), v)
	logFile.Println(str)
	log.Println(str)
}

func ForWarning(v ...interface{}) {
	str := fmt.Sprintf("%s |WARNING|: %v\n", time.Now().Format(TimeFormat), v)
	logFile.Println(str)
	log.Println(str)
}

func ForError(v ...interface{}) {
	str := fmt.Sprintf("%s |ERROR|: %v\n", time.Now().Format(TimeFormat), v)
	logFile.Println(str)
	log.Fatalln(str)
}
