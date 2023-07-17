package ego

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"path"
	"runtime"
	"strings"
)

func CallerInfo(skip int) string {
	rpc := make([]uintptr, 1)
	//skip 1:本方法信息；2:本方法的调用者方法信息；3:本方法的调用者方法的调用者方法信息
	n := runtime.Callers(skip, rpc[:])
	if n < 1 {
		return "-"
	}
	frame, _ := runtime.CallersFrames(rpc).Next()
	filePath := strings.ReplaceAll(frame.File, projectRootPath(), "")
	funcName := strings.Split(frame.Function, ".")[1]
	return fmt.Sprintf("%s:%d method:%s()", filePath, frame.Line, funcName)
}

func CallerInfoL(skip int) string {
	rpc := make([]uintptr, 1)
	//skip 1:本方法信息；2:本方法的调用者方法信息；3:本方法的调用者方法的调用者方法信息
	n := runtime.Callers(skip, rpc[:])
	if n < 1 {
		return "-"
	}
	frame, _ := runtime.CallersFrames(rpc).Next()
	filePath := strings.ReplaceAll(frame.File, projectRootPath(), "")
	return fmt.Sprintf("%s:%d", filePath, frame.Line)
}

func projectRootPath() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir+"/", "\\", "/", -1)
}

func getCallerInfo(skip int) (info string) {
	pc, file, lineNo, ok := runtime.Caller(skip)
	if !ok {

		info = "runtime.Caller() failed"
		return
	}
	funcName := runtime.FuncForPC(pc).Name()
	fileName := path.Base(file) // Base函数返回路径的最后一个元素
	return fmt.Sprintf("FuncName:%s, file:%s, line:%d ", funcName, fileName, lineNo)
}
