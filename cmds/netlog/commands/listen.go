package commands

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"io"
	"net/http"

	"github.com/google/subcommands"
	"gophers.dev/pkgs/extractors/env"
	"gophers.dev/pkgs/netlog"
)

const (
	listenCmdName     = "listen"
	listenCmdSynopsis = "Listen for log events"
	listenCmdUsage    = "listen"
)

func ListenCmd() (subcommands.Command, string) {
	return new(listenCmd), "[listening]"
}

type listenCmd struct {
}

func (l *listenCmd) Name() string {
	return listenCmdName
}

func (l *listenCmd) Synopsis() string {
	return listenCmdSynopsis
}

func (l *listenCmd) Usage() string {
	return listenCmdUsage
}

func (l *listenCmd) SetFlags(set *flag.FlagSet) {
	// no flags when listening
}

func (l *listenCmd) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	if err := l.execute(); err != nil {
		fail("failed to run listen server", err)
		return subcommands.ExitFailure
	}
	return subcommands.ExitSuccess
}

func (l *listenCmd) execute() error {
	return start(env.OS)
}

func start(e env.Environment) error {
	var (
		address string
		port    int
	)

	if err := env.Parse(e, env.Schema{
		netlog.EnvAddress: env.String(&address, false),
		netlog.EnvPort:    env.Int(&port, false),
	}); err != nil {
		return err
	}

	if address == "" {
		address = netlog.DefaultAddress
	}

	if port == 0 {
		port = netlog.DefaultPort
	}

	listen := fmt.Sprintf("%s:%d", address, port)
	fmt.Printf("netlog: listening @ %s\n", listen)
	return http.ListenAndServe(listen, http.HandlerFunc(handler))
}

func handler(w http.ResponseWriter, r *http.Request) {

	level := r.Header.Get("X-NetLog-Level")

	b, _ := io.ReadAll(r.Body)
	b = bytes.TrimSpace(b)
	msg := fmt.Sprintf("[%s] %s", level, string(b))
	fmt.Println(msg)
}
