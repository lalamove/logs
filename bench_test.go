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
// 100000	     21891 ns/op	    1585 B/op	       8 allocs/op ( Use native caller option )
// 100000	     20152 ns/op	    1554 B/op	       9 allocs/op ( Use all zap builtin caller option )
// 1000000	      1812 ns/op	    1542 B/op	       8 allocs/op ( Use zap.NewProductionConfig() )
//{
//    "level":"error",
//    "time":"2018-01-12T03:46:33.991083023Z",
//    "src_file":"logs/bench_test.go:27",
//    "message":"I am a Debug",
//    "src_line":"27",
//    "context":{
//        "f0":"I go to school by bus",
//        "f1":"Goodest english"
//    },
//    "backtrace":"github.com/lalamove-go/logs_test.BenchmarkLalamoveErrorLogger.func1\n\t/home/alpha/works/src/github.com/lalamove-go/logs/bench_test.go:27\ntesting.(*B).RunParallel.func1\n\t/home/alpha/go/src/testing/benchmark.go:625"
//}
func BenchmarkLalamoveErrorLogger(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			lalamove.Logger().Error("I am a Debug", zap.String("f0", "I go to school by bus"), zap.String("f1", "Goodest english"))
		}
		lalamove.Logger().Sync()
	})

}

//{
//   "level":"error",
//   "ts":1515729052.9642668,
//   "caller":"logs/bench_test.go:59",
//   "msg":"I am a Debug",
//   "context":{
//      "f0":"I go to school by bus",
//      "f1":"Goodest english"
//   },
//   "stacktrace":"github.com/lalamove-go/logs_test.BenchmarkUberZapErrorLogger.func1\n\t/home/alpha/works/src/github.com/lalamove-go/logs/bench_test.go:59\ntesting.(*B).RunParallel.func1\n\t/home/alpha/go/src/testing/benchmark.go:625"
//}
// 1000000	      1385 ns/op	    1479 B/op	       8 allocs/op
func BenchmarkUberZapErrorLogger(b *testing.B) {
	logger, _ := zap.NewProduction()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.With(zap.Namespace(lalamove.CustomFieldKey)).Error("I am a Debug", zap.String("f0", "I go to school by bus"), zap.String("f1", "Goodest english"))
		}
		logger.Sync()
	})
}

// 1000000	      1828 ns/op	    1537 B/op	       8 allocs/op
func BenchmarkLalamoveDebugLogger(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			lalamove.Logger().Debug("I am a Debug", zap.String("f0", "I go to school by bus"), zap.String("f1", "Goodest english"))
		}
		lalamove.Logger().Sync()
	})

}

//{
//   "level":"debug",
//   "ts":1515729585.4879098,
//   "caller":"logs/bench_test.go:84",
//   "msg":"I am a Debug",
//   "context":{
//      "f0":"I go to school by bus",
//      "f1":"Goodest english"
//   }
//}
// 1000000	      1098 ns/op	    1474 B/op	       8 allocs/op
func BenchmarkUberZapDebugLogger(b *testing.B) {
	cfg := zap.NewProductionConfig()
	cfg.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	logger, _ := cfg.Build()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.With(zap.Namespace(lalamove.CustomFieldKey)).Debug("I am a Debug", zap.String("f0", "I go to school by bus"), zap.String("f1", "Goodest english"))
		}
		logger.Sync()
	})
}
