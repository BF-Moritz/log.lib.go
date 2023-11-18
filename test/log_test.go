package loglibgo_test

import (
	"os"
	"testing"

	loglibgo "github.com/BF-Moritz/log.lib.go"
	"github.com/BF-Moritz/log.lib.go/consts"
	"github.com/BF-Moritz/log.lib.go/enum"
	"github.com/BF-Moritz/log.lib.go/test/daos/testdao"
)

var logger *loglibgo.Logger

type TestStruct struct {
	Number int
}

func (t TestStruct) ToString() string {
	logger.LogDebug("b")
	return "TestStruct"
}

func TestLog(t *testing.T) {
	levels := []enum.LogLevel{enum.LogLevelDebug, enum.LogLevelInfo, enum.LogLevelError, enum.LogLevelNone}

	file, err := os.Create("test.log")
	if err != nil {
		t.Error(err)
	}
	defer file.Close()

	logger = loglibgo.NewLogger(enum.LogLevelDebug, true, true, file)
	logger.SetTimeFormatString(consts.ExactDateTimeFormat)
	d := testdao.NewDAO()

	for _, level := range levels {
		t.Log("Testing level", level.ToString())
		logger = loglibgo.NewLogger(level, true, true, file)
		logger.SetTimeFormatString(consts.ExactDateTimeFormat)

		testStruct := TestStruct{Number: 1}
		testStruct.ToString()
		d.Test(logger)

		logger.LogDebug("Test Debug")
		logger.LogInfo("Test Info")
		logger.LogError("Test Error")
	}
}
