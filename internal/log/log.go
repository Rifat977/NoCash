package log

import (
	"fmt"
	"github.com/fatih/color"
	"log"
	"os"
)

var (
	infoColor    = color.New(color.FgHiCyan).SprintFunc()
	warningColor = color.New(color.FgHiYellow).SprintFunc()
	errorColor   = color.New(color.FgHiRed).SprintFunc()
	successColor = color.New(color.FgHiGreen).SprintFunc()
	debugColor   = color.New(color.FgHiMagenta).SprintFunc()
)

func Info(v ...interface{}) {
	log.Println(infoColor("[INFO]"), fmt.Sprint(v...))
}

func Infof(format string, v ...interface{}) {
	log.Println(infoColor("[INFO]"), fmt.Sprintf(format, v...))
}

func Success(v ...interface{}) {
	log.Println(successColor("[SUCCESS]"), fmt.Sprint(v...))
}

func Successf(format string, v ...interface{}) {
	log.Println(successColor("[SUCCESS]"), fmt.Sprintf(format, v...))
}

func Warning(v ...interface{}) {
	log.Println(warningColor("[WARNING]"), fmt.Sprint(v...))
}

func Warningf(format string, v ...interface{}) {
	log.Println(warningColor("[WARNING]"), fmt.Sprintf(format, v...))
}

func Error(v ...interface{}) {
	log.Println(errorColor("[ERROR]"), fmt.Sprint(v...))
}

func Errorf(format string, v ...interface{}) {
	log.Println(errorColor("[ERROR]"), fmt.Sprintf(format, v...))
}

func Debug(v ...interface{}) {
	if os.Getenv("DEBUG") == "true" {
		log.Println(debugColor("[DEBUG]"), fmt.Sprint(v...))
	}
}

func Debugf(format string, v ...interface{}) {
	if os.Getenv("DEBUG") == "true" {
		log.Println(debugColor("[DEBUG]"), fmt.Sprintf(format, v...))
	}
}

func Fatal(v ...interface{}) {
	log.Fatal(errorColor("[FATAL]"), fmt.Sprint(v...))
}

func Fatalf(format string, v ...interface{}) {
	log.Fatalf(errorColor("[FATAL]")+" "+format, v...)
}
