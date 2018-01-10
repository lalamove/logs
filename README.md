# Objective
Offer a Golang logger based on Lalamove k8s logging format.

# Install
```
go get -u github.com/lalamove-go/logs
```

# Usage
```go
import lalamove "github.com/lalamove-go/logs"

func main(){
    lalamove.Logger().Debug("I am a Debug")
    // {"level":"debug","time":"2017-12-23T05:42:47.752491212Z","src_file":"logs/logs_test.go:10","message":"I am a Debug","src_line":"10"}

    lalamove.Logger().Info("I am an Info")
    // {"level":"info","time":"2017-12-23T05:42:47.752524440Z","src_file":"logs/logs_test.go:11","message":"I am an Info","src_line":"11"}

    lalamove.Logger().Warn("I am a Warn")
    // {"level":"warning","time":"2017-12-23T05:42:47.752541092Z","src_file":"logs/logs_test.go:12","message":"I am a Warn","src_line":"12","backtrace":"github.com/logs.TestGetLalamoveLoggerPassDebug\n\t/home/alpha/works/src/github.com/logs/logs_test.go:12\ntesting.tRunner\n\t/home/alpha/go/src/testing/testing.go:746"}

    lalamove.Logger().Error("I am an Error")
    // {"level":"error","time":"2017-12-23T05:42:47.752575758Z","src_file":"logs/logs_test.go:13","message":"I am an Error","src_line":"13","backtrace":"github.com/logs.TestGetLalamoveLoggerPassDebug\n\t/home/alpha/works/src/github.com/logs/logs_test.go:13\ntesting.tRunner\n\t/home/alpha/go/src/testing/testing.go:746"}

    lalamove.Logger().Fatal("I am a Fatal")
    // {"level":"fatal","time":"2017-12-23T05:30:41.901899661Z","src_file":"logs/logs_test.go:49","message":"I am a Fatal","src_line":"49","backtrace":"github.com/logs.TestGetLalamoveLoggerPassFatal\n\t/home/alpha/works/src/github.com/logs/logs_test.go:49\ntesting.tRunner\n\t/home/alpha/go/src/testing/testing.go:746"}

    // Testing with extra fields
    lalamove.Logger().Debug("I am a Debug", zap.String("f0", "I go to school by bus"),zap.String("f1", "Goodest english"))
    // {"level":"debug","time":"2018-01-03T03:42:42.145362160Z","src_file":"logs/logs_test.go:40","message":"I am a Debug","src_line":"40","context":{"f0":"I go to school by bus","f1":"Goodest english"}}

    // Remember to close the logger
   defer lalamove.Logger().Sync()
}

```
# Run test
```
go test . -v
```

# Report issue
alpha.wong@lalamove.com

# Credit
- francois.parquet@lalamove.com
- mikael.knutsson@lalamove.com
- milan.r@lalamove.com

# License
Released under the MIT License.
