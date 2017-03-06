package logger

import (
	"log"
	"os"
)

var debug = false
var force = false

var std = log.New(os.Stdout, "", log.LstdFlags)

func SetDebugMode() {
	debug = true
}

func SetForceMode() {
	force = true
}

func Debugf(format string, para ...interface{}) {
	if debug {
		std.Printf("\033[33m[DBUG] "+format+"\033[0m", para...)
	}
}

func Infof(format string, para ...interface{}) {
	std.Printf("\033[32m[INFO] "+format+"\033[0m", para...)
}

func Warnf(format string, para ...interface{}) {
	std.Printf("\033[35m[WARN] "+format+"\033[0m", para...)
}

func Errorf(format string, para ...interface{}) {
	if force {
		std.Fatalf("\033[31m[FATL] "+format+"\033[0m", para...)
	} else {
		std.Printf("\033[31m[ERRO] "+format+"\033[0m", para...)
	}
}

func Fatalf(format string, para ...interface{}) {
	std.Fatalf("\033[31m[FATL] "+format+"\033[0m", para...)
}
