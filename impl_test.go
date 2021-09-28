/*
 * Copyright (c) 2020-present unTill Pro, Ltd.
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package logger

import (
	"context"
	"runtime"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_BasicUsage(t *testing.T) {

	ctx := context.Background()

	// "Hello world"
	{
		// Use context.Background() if you do not have a context
		Error(context.Background(), "Hello world", "arg1", "arg2")
		Warning(ctx, "My warning")
		Info(ctx, "My info")

		// IsDebug() is used to avoid unnecessary calculations
		if IsDebug(ctx) {
			Debug(ctx, "!!! You should NOT see it since default level is INFO")
		}
	}
	// Changing LogLevel
	{
		SetLogLevel(LogLevelDebug)
		if IsDebug(ctx) {
			Debug(ctx, "Now you should see my Debug")
		}
		SetLogLevel(LogLevelError)
		Debug(ctx, "!!! You should NOT see my Debug")
		Warning(ctx, "!!! You should NOT see my warning")
		SetLogLevel(LogLevelInfo)
		Warning(ctx, "You should see my warning")
		Warning(ctx, "You should see my info")
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
	Error(context.Background(), "OOPS")
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
