/*
 * Copyright (c) 2020-present unTill Pro, Ltd.
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package logger

import (
	"testing"
	"time"

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

func Test_CheckSetLevels(t *testing.T) {
	// 1. Info LogLevel = LogLevelInfo
	SetLogLevel(LogLevelInfo)
	assert.False(t, IsDebug())
	assert.True(t, IsEnabled(LogLevelInfo))
	assert.True(t, IsEnabled(LogLevelWarning))
	assert.True(t, IsEnabled(LogLevelError))

	// 2. Debug LogLevel = LogLevelDebug
	SetLogLevel(LogLevelDebug)
	assert.True(t, IsDebug())
	assert.True(t, IsEnabled(LogLevelInfo))
	assert.True(t, IsEnabled(LogLevelWarning))
	assert.True(t, IsEnabled(LogLevelError))

	// 3. Warning LogLevel = LogLevelWarning
	SetLogLevel(LogLevelWarning)
	assert.False(t, IsDebug())
	assert.False(t, IsEnabled(LogLevelInfo))
	assert.True(t, IsEnabled(LogLevelWarning))
	assert.True(t, IsEnabled(LogLevelError))

	// 4. Error LogLevel = LogLevelError
	SetLogLevel(LogLevelError)
	assert.False(t, IsDebug())
	assert.False(t, IsEnabled(LogLevelInfo))
	assert.False(t, IsEnabled(LogLevelWarning))
	assert.True(t, IsEnabled(LogLevelError))
}

func Test_CheckRightPrefix(t *testing.T) {
	// 1. Info LogLevel = LogLevelInfo
	SetLogLevel(LogLevelInfo)
	assert.Equal(t, getLevelPrefix(globalLogPrinter.logLevel), infoPrefix)

	// 2. Debug LogLevel = LogLevelDebug
	SetLogLevel(LogLevelDebug)
	assert.Equal(t, getLevelPrefix(globalLogPrinter.logLevel), debugPrefix)

	// 3. Warning LogLevel = LogLevelWarning
	SetLogLevel(LogLevelWarning)
	assert.Equal(t, getLevelPrefix(globalLogPrinter.logLevel), warningPrefix)

	// 4. Error LogLevel = LogLevelError
	SetLogLevel(LogLevelError)
	assert.Equal(t, getLevelPrefix(globalLogPrinter.logLevel), errorPrefix)

	// 5. Unexisting level
	SetLogLevel(7)
	assert.Equal(t, getLevelPrefix(globalLogPrinter.logLevel), "")

	SetLogLevel(LogLevelInfo)
}

type mystruct struct {
}

func (m *mystruct) iWantToLog() {
	Error("OOPS")
}

func Benchmark_FuncForPC(b *testing.B) {
	var funcName string

	start := time.Now()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		funcName, _ = globalLogPrinter.getFuncName()
	}
	assert.Equal(b, funcName, "testing.(*B).runN")

	elapsed := time.Since(start).Seconds()
	b.ReportMetric(float64(b.N)/elapsed, "rps")
}
