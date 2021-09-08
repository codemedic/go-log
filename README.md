# Go Log

[![Go Reference](https://pkg.go.dev/badge/github.com/codemedic/go-log.svg)](https://pkg.go.dev/github.com/codemedic/go-log)
[![license](https://img.shields.io/github/license/codemedic/go-log?style=flat)](https://raw.githubusercontent.com/codemedic/go-log/master/LICENSE)

GoLog adds level based logging to logger(s) from standard library.

GoLog's API is designed to be expressive with sensible configuration defaults and to be easy to use.

## Installation

    go get -u github.com/codemedic/go-log

## Getting Started

To get started, import the library and use one of the constructor functions, wrapped with `log.Must`. Make sure the
logger is closed, once you are done with it, using `defer l.Close()`. Now you are all set to start logging.

#### Example

```go
package main

import golog "github.com/codemedic/go-log"

func main() {
  // create syslog logger
  l := golog.Must(golog.NewSyslog())
  
  // make sure resources are freed up when we are done
  defer l.Close()

  // hello to the world
  l.Print("hello world!")
}
```

> **NOTE**<br/>
> Functions `log.Print` and `log.Printf` logs to `Debug` level by default. It is preferable to use a method that log to
> a specific level. The level logged to by `log.Print` and `log.Printf` can be changed using [`WithPrintLevel`](https://pkg.go.dev/github.com/codemedic/go-log#WithPrintLevel).

You can find more [examples here](https://pkg.go.dev/github.com/codemedic/go-log#pkg-examples).

## Leveled Logging

The methods below provides leveled logging. They follow the same pattern as `fmt.Print` and `fmt.Printf` and uses the
same format specification.

```go
// Log string message at specific levels
Debug(value ...interface{})
Info(value ...interface{})
Warning(value ...interface{})
Error(value ...interface{})

// Log formatted string message at specific levels, similar to log.Printf from standard library
Debugf(format string, value ...interface{})
Infof(format string, value ...interface{})
Warningf(format string, value ...interface{})
Errorf(format string, value ...interface{})
```

#### Example

```go
package main

import (
  "errors"
  golog "github.com/codemedic/go-log"
)

func main() {
  l := golog.Must(golog.NewSyslog())
  defer l.Close()

  l.Debug("debug message")
  l.Debugf("formatted %s message", "debug")

  l.Info("informational message")
  l.Infof("formatted %s message", "informational")

  l.Warning("warning message")
  l.Warningf("formatted %s message", "warning")

  l.Error("error message")
  l.Errorf("formatted %v message", errors.New("error"))
}
```

## Options / Settings

See [documentation](https://pkg.go.dev/github.com/codemedic/go-log#Option) for all available `Options`.

#### Example

```go
package main

import golog "github.com/codemedic/go-log"

func main() {
  l := golog.Must(golog.NewSyslog(
    golog.OptionsMust(golog.Options(
      golog.WithLevelFromEnv("LOG_THRESHOLD", golog.Info),
      golog.WithSourceLocationFromEnv("LOG_CALLER_LOCATION", "short"),
      golog.WithSyslogTag("my-test-app"),
    ))))
  defer l.Close()

  l.Info("hello world!")
}
```

## Standard log handler

Logging via standard logger is handled by default. This is meant for cases where the logging via the standard library is
outside your control; a library used in your project for example. Those will be logged at `Info` level, but this
behaviour can be customised using [`WithStdlogSorter`](https://pkg.go.dev/github.com/codemedic/go-log#WithStdlogSorter).

#### Example

```go
package main

import (
  "bytes"
  golog "github.com/codemedic/go-log"
)

func sortStdlog(b []byte) golog.Level {
  switch {
  case bytes.HasPrefix(b, []byte("WARNING")):
    fallthrough
  case bytes.HasPrefix(b, []byte("ERROR")):
    return golog.Warning
  case bytes.HasPrefix(b, []byte("DEBUG")):
    return golog.Disabled
  default:
    return golog.Info
  }
}

func main() {
  l, _ := golog.NewSyslog(golog.WithStdlogSorter(sortStdlog))
  defer l.Close()

  l.Info("hello world!")
}
```