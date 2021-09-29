/*
 * Copyright (c) 2020-present unTill Pro, Ltd.
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package logger

import (
	"runtime"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_BasicUsage(t *testing.T) {

	// "Hello world"
	{
		Error("Hello world", "arg1", "arg2")
		Warning("My warning")
		Info("My info")

		// IsDebug() is used to avoid unnecessary calculations
		if IsDebug() {
			Debug("!!! You should NOT see it since default level is INFO")
		}
	}
	// Changing LogLevel
	{
		SetLogLevel(LogLevelDebug)
		if IsDebug() {
			Debug("Now you should see my Debug")
		}
		SetLogLevel(LogLevelError)
		Debug("!!! You should NOT see my Debug")
		Warning("!!! You should NOT see my warning")
		SetLogLevel(LogLevelInfo)
		Warning("You should see my warning")
		Warning("You should see my info")
	}
	// Let see how it looks when using from methods
	{
		m := mystruct{}
		m.iWantToLog()
	}
}

type mystruct struct {
}

func (m *mystruct) iWantToLog() {
	Error("OOPS")
}

func Benchmark_FuncForPC(b *testing.B) {
	var funcName string
	for i := 0; i < b.N; i++ {
		pc, _, _, ok := runtime.Caller(2)
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
	}
	assert.True(b, len(funcName) > 0)
}
