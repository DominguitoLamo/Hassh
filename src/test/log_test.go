package test

import (
	"hassh/src/logger"
	"testing"
)

func TestLog(t *testing.T) {
	logger.DebugLog("debug: %s", "hello")
	logger.ErrorLog("error: %s", "error")
}