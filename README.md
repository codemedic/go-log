# Go Log

[![Go Reference](https://pkg.go.dev/badge/github.com/codemedic/go-log.svg)](https://pkg.go.dev/github.com/codemedic/go-log)
[![license](https://img.shields.io/github/license/codemedic/go-log?style=flat)](https://raw.githubusercontent.com/codemedic/go-log/master/LICENSE)

GoLog adds level-based logging to logger(s) from the standard library and offers means to integrate with other logging libraries. Its API is designed to be expressive with sensible configuration defaults and easy to use.

## Features

- **Leveled Logging**: Provides `Debug`, `Info`, `Warning`, and `Error` levels for categorized logging.
- **Customizable Print Level**: Default `log.Print` logs to `Debug` but can be customized with `WithPrintLevel`.
- **Syslog Support**: Includes a `NewSyslog` constructor for syslog-based logging.
- **Standard Logger Integration**: Handles logging via the standard library logger with customizable behavior.
- **Sublogger Customizations**: Supports creating subloggers with `WithLevel`, `WithPrefix`, and `WithRateLimit` for fine-grained control.
- **Options and Configuration**: Supports environment-based configuration like `WithLevelFromEnv` and `WithSourceLocationFromEnv`.
- **Resource Management**: Ensures proper cleanup with `Close` methods.

## Examples

See the documentation for [examples](https://pkg.go.dev/github.com/codemedic/go-log#pkg-examples) of how to use the library.

## Advanced Features

### Seamless Standard Log Integration

This library ensures compatibility with the standard Go logger by providing `log.Print` and `log.Printf` methods that replicate the standard logger's behavior.

While `log.Print` and `log.Printf` log at the `PrintLevel`, their use is discouraged as they lack the control and clarity of leveled logging methods. For better categorization and maintainability, use specific leveled logging methods. If needed, adjust their logging level using [`WithPrintLevel`](https://pkg.go.dev/github.com/codemedic/go-log#WithPrintLevel). Refer to [examples here](https://pkg.go.dev/github.com/codemedic/go-log#example-WithPrintLevel) for detailed usage.

### Flexible Standard Log Handling

The library automatically processes log lines from the standard Go logger, making it ideal for scenarios where standard library logging is unavoidable, such as third-party libraries or specific functionalities. By default, these logs are assigned the `Info` level, but you can customize this behavior using [`WithStdlogSorter`](https://pkg.go.dev/github.com/codemedic/go-log#WithStdlogSorter).

Explore example usage for [`WithStdlogSorter`](https://pkg.go.dev/github.com/codemedic/go-log#example-WithStdlogSorter).

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
