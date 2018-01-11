package logs_test

import (
	"testing"

	lalamove "github.com/lalamove-go/logs"
	"go.uber.org/zap"
)

// 50000	     25713 ns/op	    5208 B/op	      51 allocs/op
// 50000	     21715 ns/op	    4988 B/op	      45 allocs/op ( Change flag to production mode and remove fire to /tmp/log )
// 100000	     21558 ns/op	    4987 B/op	      45 allocs/op
// 100000	     19603 ns/op	    2836 B/op	      18 allocs/op ( Init the config in the init function and pass the config via pointer )
// 100000	     18902 ns/op	    2836 B/op	      18 allocs/op ( Compare the log level without casting to string )
// 100000	     18726 ns/op	    2876 B/op	      20 allocs/op ( Change all string to byte )
// 100000	     18939 ns/op	    2836 B/op	      18 allocs/op ( Undo previous step )
// 100000	     19796 ns/op	    1602 B/op	      12 allocs/op ( Move the namespace to core level option )
// 100000	     18141 ns/op	    1569 B/op	      11 allocs/op ( Remove NewLalamoveZapConfig function and replace it by zapcore.NewCore )
// 100000	     18011 ns/op	    1521 B/op	       8 allocs/op ( Remove zap.WrapCore function and replace it by Logger.With )
func BenchmarkLalamoveLogger(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			lalamove.Logger().Debug("I am a Debug", zap.String("f0", "I go to school by bus"), zap.String("f1", "Goodest english"))
		}
		lalamove.Logger().Sync()
	})

}

// 100000	     14700 ns/op	     272 B/op	       7 allocs/op
// 100000	     14719 ns/op	     272 B/op	       7 allocs/op
// 100000	     14725 ns/op	     272 B/op	       7 allocs/op
func BenchmarkUberZapLogger(b *testing.B) {
	logger, _ := zap.NewDevelopment()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Debug("I am a Debug", zap.String("f0", "I go to school by bus"), zap.String("f1", "Goodest english"))
		}
		logger.Sync()
	})
}
