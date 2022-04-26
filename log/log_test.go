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
	"bytes"
	"os"
	"strings"
	"testing"

	"go.uber.org/zap"
)

func testNoDebugLevel(t *testing.T, logger *DefaultLeveledLogger) {
	var outBuf bytes.Buffer
	logger.WithOutput(&outBuf)

	logger.Debug("this shouldn't be logged")
	if outBuf.Len() > 0 {
		t.Error("Debug was logged when it shouldn't have been")
	}
	logger.Debugf("this shouldn't be logged")
	if outBuf.Len() > 0 {
		t.Error("Debug was logged when it shouldn't have been")
	}
}

func testDebugLevel(t *testing.T, logger *DefaultLeveledLogger) {
	var outBuf bytes.Buffer
	logger.WithOutput(&outBuf)

	dbgMsg := "this is a debug message"
	logger.Debug(dbgMsg)
	if !strings.Contains(outBuf.String(), dbgMsg) {
		t.Errorf("Expected to find %q in %q, but didn't", dbgMsg, outBuf.String())
	}
	logger.Debugf(dbgMsg)
	if !strings.Contains(outBuf.String(), dbgMsg) {
		t.Errorf("Expected to find %q in %q, but didn't", dbgMsg, outBuf.String())
	}
}

func testWarnLevel(t *testing.T, logger *DefaultLeveledLogger) {
	var outBuf bytes.Buffer
	logger.WithOutput(&outBuf)

	warnMsg := "this is a warning message"
	logger.Warn(warnMsg)
	if !strings.Contains(outBuf.String(), warnMsg) {
		t.Errorf("Expected to find %q in %q, but didn't", warnMsg, outBuf.String())
	}
	logger.Warnf(warnMsg)
	if !strings.Contains(outBuf.String(), warnMsg) {
		t.Errorf("Expected to find %q in %q, but didn't", warnMsg, outBuf.String())
	}
}

func testErrorLevel(t *testing.T, logger *DefaultLeveledLogger) {
	var outBuf bytes.Buffer
	logger.WithOutput(&outBuf)

	errMsg := "this is an error message"
	logger.Error(errMsg)
	if !strings.Contains(outBuf.String(), errMsg) {
		t.Errorf("Expected to find %q in %q, but didn't", errMsg, outBuf.String())
	}
	logger.Errorf(errMsg)
	if !strings.Contains(outBuf.String(), errMsg) {
		t.Errorf("Expected to find %q in %q, but didn't", errMsg, outBuf.String())
	}
}

func TestDefaultLoggerFactory(t *testing.T) {
	f := DefaultLoggerFactory{
		Writer:          os.Stdout,
		DefaultLogLevel: LogLevelWarn,
		ScopeLevels: map[string]LogLevel{
			"foo": LogLevelDebug,
		},
	}

	logger := f.NewLogger("baz")
	bazLogger, ok := logger.(*DefaultLeveledLogger)
	if !ok {
		t.Error("Invalid logger type")
	}

	testNoDebugLevel(t, bazLogger)
	testWarnLevel(t, bazLogger)

	logger = f.NewLogger("foo")
	fooLogger, ok := logger.(*DefaultLeveledLogger)
	if !ok {
		t.Error("Invalid logger type")
	}

	testDebugLevel(t, fooLogger)
}

func TestDefaultLogger(t *testing.T) {
	logger := NewDefaultLeveledLoggerForScope("test1", LogLevelWarn, os.Stdout)

	testNoDebugLevel(t, logger)
	testWarnLevel(t, logger)
	testErrorLevel(t, logger)
}

func TestSetLevel(t *testing.T) {
	logger := NewDefaultLeveledLoggerForScope("testSetLevel", LogLevelWarn, os.Stdout)

	testNoDebugLevel(t, logger)
	logger.SetLevel(LogLevelDebug)
	testDebugLevel(t, logger)
}

func ExampleLogger() {
	zf := &ZapFactory{BaseLogger: zap.NewExample()}
	l := zf.NewLogger("scope")

	l.Error("test")
	l.Errorf("test printf %d", 1)

	// Output:
	// {"level":"error","logger":"scope","msg":"test"}
	// {"level":"error","logger":"scope","msg":"test printf 1"}
}
