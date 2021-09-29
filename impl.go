/*
 * Copyright (c) 2020-present unTill Pro, Ltd.
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package logger

import (
	"fmt"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

// TLogLevel s.e.
type TLogLevel int32

// Log Levels
const (
	LogLevelNone = TLogLevel(iota)
	LogLevelError
	LogLevelWarning
	LogLevelInfo
	LogLevelDebug
)

var globalLogLevel = LogLevelInfo

func SetLogLevel(logLevel TLogLevel) {
	atomic.StoreInt32((*int32)(&globalLogLevel), int32(logLevel))
}

func IsEnabled(logLevel TLogLevel) bool {
	curLogLevel := TLogLevel(atomic.LoadInt32((*int32)(&globalLogLevel)))
	return curLogLevel >= logLevel
}

func IsDebug() bool {
	return IsEnabled(LogLevelDebug)
}

var m sync.Mutex

func print(msgType string, args ...interface{}) {
	m.Lock()
	defer m.Unlock()
	var funcName string
	pc, _, line, ok := runtime.Caller(2)
	details := runtime.FuncForPC(pc)
	if ok && details != nil {
		elems := strings.Split(details.Name(), "/")
		if len(elems) > 1 {
			funcName = elems[len(elems)-1]
		} else {
			funcName = details.Name()
		}
	} else {
		funcName = ""
	}
	t := time.Now()
	out := fmt.Sprint(t.Format("01/02 15:04:05.000"))
	out += fmt.Sprint(": " + msgType)
	out += fmt.Sprintf(": [%v:%v]:", funcName, line)
	if len(args) > 0 {
		var s string
		for _, arg := range args {
			s = s + fmt.Sprint(" ", arg)
		}
		for i := len(s); i < 60; i++ {
			s = s + " "
		}
		out += fmt.Sprint(s)
	}
	fmt.Println(out)
}

// Error s.e.
func Error(args ...interface{}) {
	print("*****", args...)

}

// Warning s.e.
func Warning(args ...interface{}) {
	if IsEnabled(LogLevelWarning) {
		print("!!!", args...)
	}
}

// Info s.e.
func Info(args ...interface{}) {
	if IsEnabled(LogLevelInfo) {
		print("===", args...)
	}
}

// Debug s.e.
func Debug(args ...interface{}) {
	if IsDebug() {
		print("---", args...)
	}
}
