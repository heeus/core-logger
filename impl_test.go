/*
 * Copyright (c) 2020-present unTill Pro, Ltd. and Contributors
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
	// Hello world
	{
		// nil canbe used as ctx
		Error(nil, "My error")
		Warning(nil, "My warning")
		Info(nil, "My info")

		// IsVerbose() is used to avoid unnecessary calculations
		if IsVerbose(nil) {
			Verbose(nil, "!!! You should NOT see it since default level is INFO")
		}
	}
	// Changing LogLevel
	{
		SetLogLevel(LogLevelVerbose)
		if IsVerbose(nil) {
			Verbose(nil, "Now you should see my verbose")
		}
		SetLogLevel(LogLevelError)
		Verbose(nil, "!!! You should NOT see my verbose")
		Warning(nil, "!!! You should NOT see my warning")
		SetLogLevel(LogLevelInfo)
		Warning(nil, "You should see my warning")
		Warning(nil, "You should see my info")
	}
	// From struct
	{
		m := mystruct{}
		m.logMe()
	}
}

type mystruct struct {
}

func (m *mystruct) logMe() {
	Error(nil, "OOPS")
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
