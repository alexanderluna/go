package pocketlog_test

import (
	"cli_logger/pocketlog"
	"testing"
)

func ExampleLogger_Debugf() {
	debugLogger := pocketlog.New(pocketlog.LevelDebug)
	debugLogger.Debugf("Hello, %s", "world")
	// Output: Hello, world
}

const (
	debugMessage = "Here is some helpful information"
	infoMessage  = "Here is something you might not have known"
	warnMessage  = "This is causing some odd behavior"
	errorMessage = "This is fatal to the current operation"
	fatalMessage = "This caused the app or service to shut down"
)

func TestLogger_DebugInfoWarnErrorFatal(t *testing.T) {
	testCases := map[string]struct {
		level  pocketlog.Level
		output string
	}{
		"debug": {
			level: pocketlog.LevelDebug,
			output: debugMessage + "\n" +
				infoMessage + "\n" +
				warnMessage + "\n" +
				errorMessage + "\n" +
				fatalMessage + "\n",
		},
		"info": {
			level: pocketlog.LevelInfo,
			output: infoMessage + "\n" +
				warnMessage + "\n" +
				errorMessage + "\n" +
				fatalMessage + "\n",
		},
		"warn": {
			level: pocketlog.LevelWarn,
			output: warnMessage + "\n" +
				errorMessage + "\n" +
				fatalMessage + "\n",
		},
		"error": {
			level: pocketlog.LevelError,
			output: errorMessage + "\n" +
				fatalMessage + "\n",
		},
		"fatal": {
			level:  pocketlog.LevelFatal,
			output: fatalMessage + "\n",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			// use our test writer instead of Stdout
			tw := &testWriter{}
			tl := pocketlog.New(tc.level, pocketlog.WithOutput(tw))

			tl.Debugf(debugMessage)
			tl.Infof(infoMessage)
			tl.Warnf(warnMessage)
			tl.Errorf(errorMessage)
			tl.Fatalf(fatalMessage)

			if tw.contents != tc.output {
				t.Errorf("Wrong output, expect: %q, got: %q", tc.output, tw.contents)
			}
		})
	}
}

// testWriter is a struct that implements io.Writer.
// this tests that we can write to a custom output instead of Stdout
type testWriter struct {
	contents string
}

// Write implements the io.Writer interface
func (tw *testWriter) Write(p []byte) (n int, err error) {
	tw.contents = tw.contents + string(p)
	return len(p), nil
}
