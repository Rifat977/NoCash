package log

import "log"

func Info(v ...interface{}) {
	log.Println(v...)
}

func Error(v ...interface{}) {
	log.Println(v...)
}

func Fatal(v ...interface{}) {
	log.Fatal(v...)
}

func Fatalf(format string, v ...interface{}) {
	log.Fatalf(format, v...)
}

func Println(v ...interface{}) {
	log.Println(v...)
}

func Printf(format string, v ...interface{}) {
	log.Printf(format, v...)
}

func Printlnf(format string, v ...interface{}) {
	log.Printf(format, v...)
}