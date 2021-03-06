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
	DefaultName    = "default"
	DefaultTimeout = 1 * time.Second
	EnvAddress     = "NETLOG_ADDRESS"
	EnvPort        = "NETLOG_PORT"
)

type Log struct {
	address string
	port    int
	names   []string
	timeout time.Duration
	client  *http.Client
}

func (l *Log) copy() *Log {
	c := New(
		WithAddress(l.address),
		WithPort(l.port),
	)
	c.names = make([]string, len(l.names))
	copy(c.names, l.names)
	return c
}

type Option func(*Log)

func New(opts ...Option) *Log {
	l := &Log{
		address: DefaultAddress,
		port:    DefaultPort,
		timeout: DefaultTimeout,
		names:   []string{DefaultName},
	}
	for _, opt := range opts {
		opt(l)
	}
	l.client = cleanhttp.DefaultClient()
	l.client.Timeout = l.timeout
	return l
}

func Named(name string, opts ...Option) *Log {
	return New(append(opts, WithName(name))...)
}

func WithName(name string) Option {
	return func(l *Log) {
		l.names = []string{name}
	}
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

	request.Header.Set(HeaderLevel, level.String())
	request.Header.Set(HeaderName, l.name())
	request.Header.Set(HeaderContentType, "text/plain; charset=UTF-8")
	request.Header.Set(HeaderUserAgent, "netlog/v1")

	if _, err = l.client.Do(request); err != nil {
		panic(err)
	}
}

func (l *Log) name() string {
	return strings.Join(l.names, ".")
}

func (l *Log) Log(level hclog.Level, msg string, args ...interface{}) {
	panic("not implemented")
}

func (l *Log) format(msg string, args ...interface{}) string {
	if len(args)%2 != 0 {
		return msg + " !BAD-ARGS! " + fmt.Sprintf("%v", args)
	}
	var sb strings.Builder
	sb.WriteString(msg)
	sb.WriteString(" ")
	for i := 0; i < len(args); i += 2 {
		sb.WriteString(fmt.Sprintf("%s=%v ", args[i].(string), args[i+1]))
	}
	return strings.TrimSpace(sb.String())
}

func (l *Log) Trace(msg string, args ...interface{}) {
	l.send(hclog.Trace, l.format(msg, args...))
}

func (l *Log) Debug(msg string, args ...interface{}) {
	l.send(hclog.Debug, l.format(msg, args...))
}

func (l *Log) Info(msg string, args ...interface{}) {
	l.send(hclog.Info, l.format(msg, args...))
}

func (l *Log) Warn(msg string, args ...interface{}) {
	l.send(hclog.Warn, l.format(msg, args...))
}

func (l *Log) Error(msg string, args ...interface{}) {
	l.send(hclog.Error, l.format(msg, args...))
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
	return l
}

func (l *Log) Name() string {
	return l.name()
}

func (l *Log) Named(name string) hclog.Logger {
	c := l.copy()
	c.names = append(c.names, name)
	return c
}

func (l *Log) ResetNamed(name string) hclog.Logger {
	l.names = []string{DefaultName}
	return l
}

func (l *Log) SetLevel(level hclog.Level) {
	return
}

func (l *Log) StandardLogger(opts *hclog.StandardLoggerOptions) *log.Logger {
	panic("not implemented")
}

func (l *Log) StandardWriter(opts *hclog.StandardLoggerOptions) io.Writer {
	panic("not implemented")
}
