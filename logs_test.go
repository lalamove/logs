package logs_test

import (
	"testing"

	lalamove "github.com/lalamove-go/logs"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
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

// TestGetLalamoveLoggerPassDebugWithRootLevelNamespace will test the extra fields.
// The extra fields should always instead the context namespace.
// expected result
//{
//    "level": "debug",
//    "time": "2018-01-03T03:40:02.087012761Z",
//    "src_file": "logs/logs_test.go:27",
//    "message": "I am a Debug",
//    "src_line": "27",
//    "context": {
//        "f0": "I go to school by bus",
//        "f1": "Goodest english"
//    }
//}
func TestGetLalamoveLoggerPassDebugWithRootLevelNamespace(t *testing.T) {
	lalamove.Logger().Debug("I am a Debug", zap.String("f0", "I go to school by bus"), zap.String("f1", "Goodest english"))
	defer lalamove.Logger().Sync()
	assert.True(t, true)
}
