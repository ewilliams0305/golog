package golog

import (
	"errors"
	"fmt"
	"testing"
)

func Test_CreateLogger_ReturnsLogger(t *testing.T) {

	// ARRANGE
	logger := LoggingConfiguration().
		Configure(Verbose, "[%l %t] %m").
		WriteTo(&testsink{}).
		CreateLogger()

	// ACT

	// ASSERT
	if logger == nil {
		t.Errorf("LOGGER IS NIL")
	}
}

func Test_CreateLogger_ReturnsLogger_WithVerbosityVerbose(t *testing.T) {

	// ARRANGE
	logger := LoggingConfiguration().
		Configure(Verbose, "[%l %t] %m").
		WriteTo(&testsink{}).
		CreateLogger()

	// ACT

	// ASSERT
	if logger.CurrentLevel() != Verbose {
		t.Errorf("GOT %q, EXPECTED %q", logger.CurrentLevel(), Verbose)
	}
}

func Test_CreateLogger_ReturnsLogger_WithVerbosityDebug(t *testing.T) {

	// ARRANGE
	logger := LoggingConfiguration().
		Configure(Debug, "[%l %t] %m").
		WriteTo(&testsink{}).
		CreateLogger()

	// ACT

	// ASSERT
	if logger.CurrentLevel() != Debug {
		t.Errorf("GOT %q, EXPECTED %q", logger.CurrentLevel(), Debug)
	}
}

func Test_CreateLogger_ReturnsLogger_WithVerbosityInformation(t *testing.T) {

	// ARRANGE
	logger := LoggingConfiguration().
		Configure(Information, "[%l %t] %m").
		WriteTo(&testsink{}).
		CreateLogger()

	// ACT

	// ASSERT
	if logger.CurrentLevel() != Information {
		t.Errorf("GOT %q, EXPECTED %q", logger.CurrentLevel(), Information)
	}
}

func Test_CreateLogger_ReturnsLogger_WithVerbosityWarn(t *testing.T) {

	// ARRANGE
	logger := LoggingConfiguration().
		Configure(Warn, "[%l %t] %m").
		WriteTo(&testsink{}).
		CreateLogger()

	// ACT

	// ASSERT
	if logger.CurrentLevel() != Warn {
		t.Errorf("GOT %q, EXPECTED %q", logger.CurrentLevel(), Warn)
	}
}

func Test_CreateLogger_ReturnsLogger_WithVerbosityError(t *testing.T) {

	// ARRANGE
	logger := LoggingConfiguration().
		Configure(Error, "[%l %t] %m").
		WriteTo(&testsink{}).
		CreateLogger()

	// ACT

	// ASSERT
	if logger.CurrentLevel() != Error {
		t.Errorf("GOT %q, EXPECTED %q", logger.CurrentLevel(), Error)
	}
}

func Test_CreateLogger_ReturnsLogger_WithVerbosityFatal(t *testing.T) {

	// ARRANGE
	logger := LoggingConfiguration().
		Configure(Fatal, "[%l %t] %m").
		WriteTo(&testsink{}).
		CreateLogger()

	// ACT

	// ASSERT
	if logger.CurrentLevel() != Fatal {
		t.Errorf("GOT %q, EXPECTED %q", logger.CurrentLevel(), Fatal)
	}
}

func Test_CreateLogger_ReturnsLogger_DoesntWriteDebug(t *testing.T) {

	// ARRANGE
	logger := LoggingConfiguration().
		Configure(Information, "[%l %t] %m").
		WriteTo(&testsink{}).
		CreateLogger()

	// ACT

	logger.Debug("Writing Debug message should not write as the verbosity is set to debug")

	// ASSERT
	if wrote {
		t.Errorf("Debug message was written with higher verbostity")
	}

	wrote = false
}

func Test_CreateLogger_ReturnsLogger_DoesntWriteInformation(t *testing.T) {

	// ARRANGE
	logger := LoggingConfiguration().
		Configure(Warn, "[%l %t] %m").
		WriteTo(&testsink{}).
		CreateLogger()

	// ACT

	logger.Information("Writing Information message should not write as the verbosity is set to debug")

	// ASSERT
	if wrote {
		t.Errorf("Information message was written with higher verbostity")
	}

	wrote = false
}

func Test_CreateLogger_ReturnsLogger_DoesntWriteWarn(t *testing.T) {

	// ARRANGE
	logger := LoggingConfiguration().
		Configure(Error, "[%l %t] %m").
		WriteTo(&testsink{}).
		CreateLogger()

	// ACT

	logger.Warn("Writing Warn message should not write as the verbosity is set to debug")

	// ASSERT
	if wrote {
		t.Errorf("Warn message was written with higher verbostity")
	}

	wrote = false
}

func Test_CreateLogger_ReturnsLogger_DoesntWriteError(t *testing.T) {

	// ARRANGE
	logger := LoggingConfiguration().
		Configure(Fatal, "[%l %t] %m").
		WriteTo(&testsink{}).
		CreateLogger()

	// ACT

	logger.Error("Writing Error message should not write as the verbosity is set to debug", errors.New("ERROR"))

	// ASSERT
	if wrote {
		t.Errorf("Error message was written with higher verbostity")
	}

	wrote = false
}

func Test_CreateLogger_ReturnsLogger_ThatWritesVerbose(t *testing.T) {

	// ARRANGE
	logger := LoggingConfiguration().
		Configure(Verbose, "[%l %t] %m").
		WriteTo(&testsink{}).
		CreateLogger()

	// ACT

	logger.Verbose("Writing Debug message should not write as the verbosity is set to debug")

	// ASSERT
	if !wrote {
		t.Errorf("Verbose message was not written with = || > verbostity")
	}

	wrote = false
}

func Test_CreateLogger_ReturnsLogger_ThatWritesDebug(t *testing.T) {

	// ARRANGE
	logger := LoggingConfiguration().
		Configure(Debug, "[%l %t] %m").
		WriteTo(&testsink{}).
		CreateLogger()

	// ACT

	logger.Debug("Writing Debug message should not write as the verbosity is set to debug")

	// ASSERT
	if !wrote {
		t.Errorf("Debug message was not written with = || > verbostity")
	}

	wrote = false
}

func Test_CreateLogger_ReturnsLogger_ThatWritesInformation(t *testing.T) {

	// ARRANGE
	logger := LoggingConfiguration().
		Configure(Information, "[%l %t] %m").
		WriteTo(&testsink{}).
		CreateLogger()

	// ACT

	logger.Information("Writing Debug message should not write as the verbosity is set to debug")

	// ASSERT
	if !wrote {
		t.Errorf("Information message was not written with = || > verbostity")
	}

	wrote = false
}

func Test_CreateLogger_ReturnsLogger_ThatWritesWarn(t *testing.T) {

	// ARRANGE
	logger := LoggingConfiguration().
		Configure(Warn, "[%l %t] %m").
		WriteTo(&testsink{}).
		CreateLogger()

	// ACT

	logger.Warn("Writing Debug message should not write as the verbosity is set to debug")

	// ASSERT
	if !wrote {
		t.Errorf("Warn message was not written with = || > verbostity")
	}

	wrote = false
}

func Test_CreateLogger_ReturnsLogger_ThatWritesError(t *testing.T) {

	// ARRANGE
	logger := LoggingConfiguration().
		Configure(Error, "[%l %t] %m").
		WriteTo(&testsink{}).
		CreateLogger()

	// ACT

	logger.Error("Writing Error message should not write as the verbosity is set to debug", errors.New("ERROR"))

	// ASSERT
	if !wrote {
		t.Errorf("Error message was not written with = || > verbostity")
	}

	wrote = false
}

func Test_CreateLogger_ReturnsLogger_ThatWritesFatal(t *testing.T) {

	// ARRANGE
	logger := LoggingConfiguration().
		Configure(Fatal, "[%l %t] %m").
		WriteTo(&testsink{}).
		CreateLogger()

	// ACT

	logger.Fatal("Writing Debug message should not write as the verbosity is set to debug", errors.New("ERROR"))

	// ASSERT
	if !wrote {
		t.Errorf("Fatal message was not written with = || > verbostity")
	}

	wrote = false
}

var wrote = false

type testsink struct {
}

func (s *testsink) WriteTo(message LogEvent) error {
	wrote = true
	fmt.Printf("%s", message.Message)
	return fmt.Errorf("%d MESSAGE WAS WRITTEN ", message.Level)
}
