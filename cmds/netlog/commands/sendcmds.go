package commands

import (
	"context"
	"flag"
	"fmt"
	"strings"

	"github.com/google/subcommands"
	"github.com/hashicorp/go-hclog"
	"github.com/shoenig/netlog"
)

type cmd struct {
	level hclog.Level
}

func (c *cmd) Name() string {
	return c.level.String()
}

func (c *cmd) Synopsis() string {
	return fmt.Sprintf("Send %s message to listener", c.level)
}

func (c *cmd) Usage() string {
	return fmt.Sprintf("%s <message>", c.level)
}

func (c *cmd) SetFlags(*flag.FlagSet) {
	// none
}

func (c *cmd) execute(s string) error {
	log := netlog.New("")

	switch c.level {
	case hclog.Off:
	case hclog.Trace:
		log.Trace(s)
	case hclog.Debug:
		log.Debug(s)
	case hclog.Info:
		log.Info(s)
	case hclog.Warn:
		log.Warn(s)
	case hclog.Error:
		log.Error(s)
	}

	return nil
}

func (c *cmd) Execute(ctx context.Context, fs *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	s := args[0].([]string)
	msg := strings.Join(s[1:], " ")

	if err := c.execute(msg); err != nil {
		fail("failed to send log message", err)
		return subcommands.ExitFailure
	}
	return subcommands.ExitSuccess
}

func LogCmd(level hclog.Level) (subcommands.Command, string) {
	return &cmd{
		level: level,
	}, "[logging]"
}
