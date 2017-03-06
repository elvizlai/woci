package logger

import "testing"

func TestLog(t *testing.T) {
	Debugf("debug")
	SetDebugMode()
	Debugf("debug")
	Infof("info")
	Warnf("warn")
	Errorf("error")
	Fatalf("fatal")
}
