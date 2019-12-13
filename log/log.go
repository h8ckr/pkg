package log

import (
	"fmt"
	"time"
)

var (
	runtimeLogger logger
)

func init() {
	runtimeLogger.level.validDepths = struct {
		debugDepth   *int
		infoDepth    *int
		errOnlyDepth *int
		silentDepth  *int
	}{debugDepth: &debugDepth, infoDepth: &infoDepth, errOnlyDepth: &errOnlyDepth, silentDepth: &silentDepth}
	runtimeLogger.level.setMaxDepth()
	runtimeLogger.level.setMinDepths()
	runtimeLogger.level.set(&errOnlyDepth)
	runtimeLogger.prefix = time.Now().String()
}

func stringFormatToBytes(format string, args ...interface{}) []byte {
	return []byte(fmt.Sprintf(format, args...))
}

func Info(arg interface{}) {
	runtimeLogger.info(stringFormatToBytes(runtimeLogger.prefix+" %v", arg))
}

func Infof(format string, args ...interface{}) {
	runtimeLogger.info(stringFormatToBytes(runtimeLogger.prefix+" "+format, args...))
}

func Infoln(arg interface{}) {
	runtimeLogger.info(stringFormatToBytes(runtimeLogger.prefix+" %v %v", arg, "\n"))
}