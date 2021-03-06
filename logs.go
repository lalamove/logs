// MIT License
//
// Copyright (c) 2017 Lalamove.com
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.
//
// Reference: https://lalamove.atlassian.net/wiki/spaces/TECH/pages/82149406/Kubernetes
// Lalamove kubernetes logging format
// {
// 		"message": "", // string describing what happened
// 		"src_file": "", // file path
// 		"src_line": "", // line number
// 		"fields": {}, // custom field here
// 		"level": "", // debug/info/warning/error/fatal
// 		"time": "", // ISO8601.nanoseconds+TZ (in node only support precision up to milliseconds)
// 		"backtrace": "" // err stack
// }
//
package logs

import (
	"runtime"
	"strconv"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	ISO8601 = "2006-01-02T15:04:05.000000000Z0700"

	Debug   = "debug"
	Info    = "info"
	Warning = "warning"
	Error   = "error"
	Fatal   = "fatal"

	TimeKey       = "time"
	LevelKey      = "level"
	CallerKey     = "src_file"
	SourceLineKey = "src_line"
	MessageKey    = "message"
	StacktraceKey = "backtrace"

	EncodingType = "json"
)

// NewLalamoveEncoderConfig will create an EncoderConfig
func NewLalamoveEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:        TimeKey,
		LevelKey:       LevelKey,
		CallerKey:      CallerKey,
		MessageKey:     MessageKey,
		StacktraceKey:  StacktraceKey,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    LalamoveLevelEncoder,
		EncodeTime:     LalamoveISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

// NewLalamoveZapConfig will create a config for zap
func NewLalamoveZapConfig() *zap.Config {
	return &zap.Config{
		Level:            zap.NewAtomicLevelAt(zapcore.DebugLevel),
		Development:      true,
		Encoding:         EncodingType,
		EncoderConfig:    NewLalamoveEncoderConfig(),
		OutputPaths:      []string{"stdout", "/tmp/logs"},
		ErrorOutputPaths: []string{"stderr"},
	}
}

// LalamoveLevelEncoder will convert the warn display string to warning
func LalamoveLevelEncoder(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	if l.String() == zapcore.WarnLevel.String() {
		// Set warn label to warning
		enc.AppendString(Warning)
	} else {
		enc.AppendString(l.String())
	}
}

// LalamoveISO8601TimeEncoder will convert the time to ISO8601 based on Lalamove k8s logging format
func LalamoveISO8601TimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.UTC().Format(ISO8601))
}

// Logger will create a zap based logger
// return a *zap.Logger for logging
func Logger() *zap.Logger {
	// Skip this function
	_, _, fl, _ := runtime.Caller(1)

	cfg := NewLalamoveZapConfig()

	showSourceLine := zap.WrapCore(func(c zapcore.Core) zapcore.Core {
		return zapcore.NewTee(
			c.With([]zapcore.Field{
				{
					Key:    SourceLineKey,
					Type:   zapcore.StringType,
					String: strconv.Itoa(fl),
				},
			}),
		)
	})

	Logger, _ := cfg.Build()
	defer Logger.Sync()
	return Logger.WithOptions(showSourceLine).With(zap.Namespace("fields"))
}
