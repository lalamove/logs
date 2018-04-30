# Objective
Offer a Golang logger based on Lalamove k8s logging format.

# Install
```
go get -u github.com/lalamove-go/logs
```

# Importance
The backtrace will only show while the level => Error
``` go
// Reference : https://github.com/uber-go/zap/blob/master/config.go#L63
// By default, stacktraces are captured for WarnLevel and above logs in
// development and ErrorLevel and above in production.
```

# Performance
```go
// Reference : https://github.com/uber-go/zap/blob/master/field.go#L182
// Stack constructs a field that stores a stacktrace of the current goroutine
// under provided key. Keep in mind that taking a stacktrace is eager and
// expensive (relatively speaking); this function both makes an allocation and
// takes about two microseconds.
func Stack(key string) zapcore.Field {
	// Returning the stacktrace as a string costs an allocation, but saves us
	// from expanding the zapcore.Field union struct to include a byte slice. Since
	// taking a stacktrace is already so expensive (~10us), the extra allocation
	// is okay.
	return String(key, takeStacktrace())
}
```

# Usage
```go
import lalamove "github.com/lalamove-go/logs"

func main(){
    lalamove.Logger().Debug("I am a Debug")
    // {"level":"debug","time":"2018-03-22T02:18:43.428168664Z","src_file":"/home/alpha/works/src/github.com/lalamove-go/logs/logs_test.go","message":"I am a Debug","src_line":"12","context":{}}

    lalamove.Logger().Info("I am an Info")
    // {"level":"info","time":"2018-03-22T02:18:43.428218327Z","src_file":"/home/alpha/works/src/github.com/lalamove-go/logs/logs_test.go","message":"I am an Info","src_line":"13","context":{}}

    lalamove.Logger().Warn("I am a Warn")
    // {"level":"warning","time":"2018-03-22T02:18:43.428246914Z","src_file":"/home/alpha/works/src/github.com/lalamove-go/logs/logs_test.go","message":"I am a Warn","src_line":"14","context":{}}

    lalamove.Logger().Error("I am an Error")
    // {"level":"error","time":"2018-03-22T02:18:43.428294281Z","src_file":"/home/alpha/works/src/github.com/lalamove-go/logs/logs_test.go","message":"I am an Error","src_line":"15","context":{},"backtrace":"github.com/lalamove-go/logs_test.TestGetLalamoveLoggerPassDebug\n\t/home/alpha/works/src/github.com/lalamove-go/logs/logs_test.go:15\ntesting.tRunner\n\t/home/alpha/go/src/testing/testing.go:777"}

    lalamove.Logger().Fatal("I am a Fatal")
    // {"level":"fatal","time":"2018-03-22T02:18:43.428246914Z","src_file":"/home/alpha/works/src/github.com/lalamove-go/logs/logs_test.go","message":"I am a Fatal","src_line":"14","context":{},"backtrace":"github.com/lalamove-go/logs_test.TestGetLalamoveLoggerPassDebug\n\t/home/alpha/works/src/github.com/lalamove-go/logs/logs_test.go:15\ntesting.tRunner\n\t/home/alpha/go/src/testing/testing.go:777"}

    // Testing with extra fields
    lalamove.Logger().Debug("I am a Debug", zap.String("f0", "I go to school by bus"),zap.String("f1", "Goodest english"))
    // {"level":"error","time":"2018-03-22T02:18:43.428491016Z","src_file":"/home/alpha/works/src/github.com/lalamove-go/logs/logs_test.go","message":"I am a Debug","src_line":"40","context":{"f2":"I go to school by bus","f3":"Goodest english"},"backtrace":"github.com/lalamove-go/logs_test.TestGetLalamoveLoggerPassErrorWithRootLevelNamespace\n\t/home/alpha/works/src/github.com/lalamove-go/logs/logs_test.go:40\ntesting.tRunner\n\t/home/alpha/go/src/testing/testing.go:777"}

    // Remember to close the logger
    defer lalamove.Logger().Sync()
}

```
# Run test
```sh
go test . -v
```

# Run benchmark
```sh
go test -bench=. -benchmem
```

# Benchmark
```
lalamove-go/logs
1000000	      1812 ns/op	    1542 B/op	       8 allocs/op

go.uber.org/zap
1000000	      1385 ns/op	    1479 B/op	       8 allocs/op
```

# Report issue
alpha.wong@lalamove.com

# Credit
- francois.parquet@lalamove.com
- mikael.knutsson@lalamove.com
- milan.r@lalamove.com
- simon.tse@lalamove.com

# License
Released under the MIT License.
