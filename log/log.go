package log

import (
	"fmt"
	"runtime/debug"
)

var (
	runtimeLogger commander
)

func init() {
	r := new("", &infoDepth)
	runtimeLogger = *r
}

func Init(prefix string, depth *int) error {
	runtimeLogger.setPrefix(prefix)
	runtimeLogger.setDepth(depth)

	return nil
}

func stringFormatToBytes(format string, args ...interface{}) []byte {
	return []byte(fmt.Sprintf(format, args...))
}

func Info(arg interface{}) {
	runtimeLogger.info(stringFormatToBytes("%s %v", *runtimeLogger.getPrefix(), arg))
}

func Infof(format string, args ...interface{}) {
	argsCombined := []interface{}{*runtimeLogger.getPrefix()}
	argsCombined = append(argsCombined, args...)
	runtimeLogger.info(stringFormatToBytes("%s "+format, argsCombined...))
}

func Infoln(arg interface{}) {
	runtimeLogger.info(stringFormatToBytes("%s %v %v", *runtimeLogger.getPrefix(), arg, "\n"))
}

func Fatal(arg interface{}) {
	runtimeLogger.info(stringFormatToBytes("%s %v", *runtimeLogger.getPrefix(), arg))
}

func Fatalf(format string, args ...interface{}) {
	argsCombined := []interface{}{*runtimeLogger.getPrefix()}
	argsCombined = append(argsCombined, args...)
	runtimeLogger.info(stringFormatToBytes("%s "+format, argsCombined...))
}

func Fatalln(arg interface{}) {
	runtimeLogger.info(stringFormatToBytes("%s %v %v", *runtimeLogger.getPrefix(), arg, "\n"))
}

func Panic(err error, message ...string) {
	stackBytes := debug.Stack()
	stack := string(stackBytes)
	runtimeLogger.error(stringFormatToBytes("%s%v", stack, "\n"))
}
