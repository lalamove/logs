package logs_test

import (
	"testing"

	lalamove "github.com/logs"
	"github.com/stretchr/testify/assert"
)

func TestGetLalamoveLoggerPassDebug(t *testing.T) {
	lalamove.Logger().Debug("I am a Debug")
	lalamove.Logger().Info("I am an Info")
	lalamove.Logger().Warn("I am a Warn")
	lalamove.Logger().Error("I am an Error")
	// It should not be called as the it will return exit code 3
	// Logger().Fatal("I am not a Fatal")
	// By default, loggers are unbuffered. However, since zap's low-level APIs allow buffering,
	// calling Sync before letting your process exit is a good habit.
	defer lalamove.Logger().Sync()
	assert.True(t, true)
}
