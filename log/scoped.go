/*
Copyright © 2022 Aaqa Ishtyaq aaqaishtyaq@gmail.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package log

import (
	"sync/atomic"
)

// LogLevel represents the level at which the logger will emit log messages
type LogLevel int32

// Set updates the LogLevel to the supplied value
func (ll *LogLevel) Set(newLevel LogLevel) {
	atomic.StoreInt32((*int32)(ll), int32(newLevel))
}

// Get retrieves the current LogLevel value
func (ll *LogLevel) Get() LogLevel {
	return LogLevel(atomic.LoadInt32((*int32)(ll)))
}

func (ll LogLevel) String() string {
	switch ll {
	case LogLevelDisabled:
		return "Disabled"
	case LogLevelError:
		return "Error"
	case LogLevelWarn:
		return "Warn"
	case LogLevelInfo:
		return "Info"
	case LogLevelDebug:
		return "Debug"
	case LogLevelTrace:
		return "Trace"
	default:
		return "UNKNOWN"
	}
}

const (
	// LogLevelDisabled completely disables logging of any events
	LogLevelDisabled LogLevel = iota
	// LogLevelError is for fatal errors which should be handled by user code,
	// but are logged to ensure that they are seen
	LogLevelError
	// LogLevelWarn is for logging abnormal, but non-fatal library operation
	LogLevelWarn
	// LogLevelInfo is for logging normal library operation (e.g. state transitions, etc.)
	LogLevelInfo
	// LogLevelDebug is for logging low-level library information (e.g. internal operations)
	LogLevelDebug
	// LogLevelTrace is for logging very low-level library information (e.g. network traces)
	LogLevelTrace
)

// LeveledLogger is the basic Logger interface
type LeveledLogger interface {
	Trace(msg string)
	Tracef(format string, args ...interface{})
	Debug(msg string)
	Debugf(format string, args ...interface{})
	Info(msg string)
	Infof(format string, args ...interface{})
	Warn(msg string)
	Warnf(format string, args ...interface{})
	Error(msg string)
	Errorf(format string, args ...interface{})
}

// LoggerFactory is the basic LoggerFactory interface
type LoggerFactory interface {
	NewLogger(scope string) LeveledLogger
}
