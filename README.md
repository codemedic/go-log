# Go Log

[![godoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/codemedic/go-log)
[![license](https://img.shields.io/github/license/codemedic/go-log?style=flat-square)](https://raw.githubusercontent.com/codemedic/go-log/master/LICENSE)

GoLog adds level based logging to logger(s) from standard library.

GoLog's API is designed be expressive for configuration with sensible defaults and to be easy to use.

## Installation

    go get -u github.com/codemedic/go-log

## Getting Started

### Logging Example

To get started, import the library and use one of the constructor functions, wrapped with `log.Must`.

```go
package main

import "github.com/codemedic/go-log"

func main() {
  // create syslog logger
  l := log.Must(log.NewSyslog())
  
  // make sure resources are freed up when we are done
  defer l.Close()

  // hello to the world
  l.Print("hello world!")
}

// Output: 2021/07/19 07:57:57.936834 main.go:18: INFO: hello world!
```

> **NOTE**
> 
> Functions `log.Print` and `log.Printf` logs to `Debug` level by default. It is preferable to use a method that log to
> a specific level. The level logged to by `log.Print` and `log.Printf` can be changed using `WithPrintLevel`.

### Leveled Logging

The methods below provides leveled logging with GoLog. See [examples](doc_test.go) for more.

```go
// Log string message at specific levels
Debug(message string)
Info(message string)
Warning(message string)
Error(message string)

// Log formatted string message at specific levels, similar to log.Printf from standard library
Debugf(format string, value ...interface{})
Infof(format string, value ...interface{})
Warningf(format string, value ...interface{})
Errorf(format string, value ...interface{})
```

### Settings aka Options

See [documentation](tbd) for all available settings.
