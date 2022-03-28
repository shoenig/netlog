package main

import (
	"context"
	"flag"
	"os"

	"github.com/google/subcommands"
	"github.com/hashicorp/go-hclog"
	"gophers.dev/pkgs/netlog/cmds/netlog/commands"
)

func main() {
	fs := flag.NewFlagSet("", flag.ContinueOnError)
	subs := subcommands.NewCommander(fs, "")
	subs.Register(commands.ListenCmd())
	subs.Register(commands.LogCmd(hclog.Trace))
	subs.Register(commands.LogCmd(hclog.Debug))
	subs.Register(commands.LogCmd(hclog.Info))
	subs.Register(commands.LogCmd(hclog.Warn))
	subs.Register(commands.LogCmd(hclog.Error))

	if err := fs.Parse(os.Args[1:]); err != nil {
		panic(err)
	}

	ctx := context.Background()
	rc := subs.Execute(ctx, fs.Args())
	os.Exit(int(rc))
}
