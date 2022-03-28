package netlog

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/hashicorp/go-cleanhttp"
	"github.com/hashicorp/go-hclog"
)

var _ hclog.Logger = (*Log)(nil)

type Level = hclog.Level

const (
	DefaultUseTLS  = false
	DefaultAddress = "127.0.0.1"
	DefaultPort    = 9999
	DefaultTimeout = 1 * time.Second
	EnvTLS         = "NETLOG_TLS"
	EnvAddress     = "NETLOG_ADDRESS"
	EnvPort        = "NETLOG_PORT"
)

type Log struct {
	tls     bool
	address string
	port    int
	timeout time.Duration
	client  *http.Client
}

type Option func(*Log)

func New(opts ...Option) *Log {
	l := &Log{
		tls:     DefaultUseTLS,
		address: DefaultAddress,
		port:    DefaultPort,
		timeout: DefaultTimeout,
	}
	for _, opt := range opts {
		opt(l)
	}
	l.client = cleanhttp.DefaultClient()
	l.client.Timeout = l.timeout
	return l
}

func WithTimeout(timeout time.Duration) Option {
	return func(l *Log) {
		l.timeout = timeout
	}
}

func WithAddress(address string) Option {
	return func(l *Log) {
		l.address = address
	}
}

func WithPort(port int) Option {
	return func(l *Log) {
		l.port = port
	}
}

func (l *Log) send(level Level, s string) {
	address := fmt.Sprintf("http://%s:%d", l.address, l.port)

	request, err := http.NewRequest(http.MethodPost, address, strings.NewReader(s))
	if err != nil {
		panic(err)
	}

	request.Header.Set(HeaderContentType, "text/plain; charset=UTF-8")
	request.Header.Set(HeaderLevel, level.String())
	request.Header.Set(HeaderUserAgent, "netlog/v1")

	if _, err = l.client.Do(request); err != nil {
		panic(err)
	}
}

func (l *Log) Log(level hclog.Level, msg string, args ...interface{}) {
	panic("not implemented")
}

func (l *Log) Trace(msg string, args ...interface{}) {
	l.send(hclog.Trace, fmt.Sprintf(msg, args...))
}

func (l *Log) Debug(msg string, args ...interface{}) {
	l.send(hclog.Debug, fmt.Sprintf(msg, args...))
}

func (l *Log) Info(msg string, args ...interface{}) {
	l.send(hclog.Info, fmt.Sprintf(msg, args...))
}

func (l *Log) Warn(msg string, args ...interface{}) {
	l.send(hclog.Warn, fmt.Sprintf(msg, args...))
}

func (l *Log) Error(msg string, args ...interface{}) {
	l.send(hclog.Error, fmt.Sprintf(msg, args...))
}

func (l *Log) IsTrace() bool {
	panic("not implemented")
}

func (l *Log) IsDebug() bool {
	panic("not implemented")
}

func (l *Log) IsInfo() bool {
	panic("not implemented")
}

func (l *Log) IsWarn() bool {
	panic("not implemented")
}

func (l *Log) IsError() bool {
	panic("not implemented")
}

func (l *Log) ImpliedArgs() []interface{} {
	panic("not implemented")
}

func (l *Log) With(args ...interface{}) hclog.Logger {
	panic("not implemented")
}

func (l *Log) Name() string {
	panic("not implemented")
}

func (l *Log) Named(name string) hclog.Logger {
	panic("not implemented")
}

func (l *Log) ResetNamed(name string) hclog.Logger {
	panic("not implemented")
}

func (l *Log) SetLevel(level hclog.Level) {
	panic("not implemented")
}

func (l *Log) StandardLogger(opts *hclog.StandardLoggerOptions) *log.Logger {
	panic("not implemented")
}

func (l *Log) StandardWriter(opts *hclog.StandardLoggerOptions) io.Writer {
	panic("not implemented")
}
