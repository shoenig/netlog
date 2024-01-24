# netlog

[![GoDoc](https://godoc.org/github.com/shoenig/netlog?status.svg)](https://godoc.org/github.com/shoenig/netlog)
[![GitHub](https://img.shields.io/github/license/shoenig/netlog.svg)](LICENSE)

## Project Overview

Package `netlog` provides an implementation of `go-hclog.Logger` that sends messages over the network to a listener.
Useful for debug logging in situations where `os.Stdout`/`os.Stderr` are not available.

Also implements colorized local logging helpers, e.g. `netlog.Yellow("message", "key", "value", ...)`

## Getting Started

The `netlog` command can be installed by running:

```bash
➜ go install github.com/shoenig/netlog/cmds/netlog@latest
```

Example usage:

#### help

```bash
➜ netlog help
Subcommands for [listening]:
	listen           Listen for log events

Subcommands for [logging]:
	debug            Send debug message to listener
	error            Send error message to listener
	info             Send info message to listener
	trace            Send trace message to listener
	warn             Send warn message to listener
```

#### listen

```bash
➜ netlog listen
netlog: listening @ 127.0.0.1:9999
```

#### logging

For convenience, the `netlog` command can directly emit log messages to a listener.

```bash
➜ netlog info This is a message!
```

It supports sending `trace`, `debug`, `info`, `warn`, `error` messages.

## API

The utility of the `netlog` package is that it implements HashiCorp's `go-hclog.Logger` interface, and can be used as a drop-in replacement in situations where terminal output is not available.

```go
import "github.com/shoenig/netlog"

// ... 

log := netlog.New("log-name")
log.Info("this is a message")
```

## Contributing

Contributions are welcome! Feel free to help make `netlog` better.

## License

The `github.com/shoenig/netlog` module is open source under the [BSD-3-Clause](LICENSE) license.
