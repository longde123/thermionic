package main

import (
	"flag"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/pborman/uuid"
	"github.com/pkg/errors"
	"github.com/spoke-d/clui"
	"github.com/spoke-d/clui/flagset"
	"github.com/spoke-d/thermionic/internal/exec"
)

type discoverServiceListCmd struct {
	ui      clui.UI
	flagset *flagset.FlagSet

	debug      bool
	address    string
	format     string
	serverCert string
	clientCert string
	clientKey  string
}

// NewDiscoverServiceListCmd creates a Command with sane defaults
func NewDiscoverServiceListCmd(ui clui.UI) clui.Command {
	c := &discoverServiceListCmd{
		ui:      ui,
		flagset: flagset.NewFlagSet("discover", flag.ExitOnError),
	}
	c.init()
	return c
}

func (c *discoverServiceListCmd) init() {
	c.flagset.BoolVar(&c.debug, "debug", false, "debug logging")
	c.flagset.StringVar(&c.address, "address", "127.0.0.1:8080", "address of the api server")
	c.flagset.StringVar(&c.format, "format", "yaml", "format to output the information json|yaml|tabular")
	c.flagset.StringVar(&c.serverCert, "server-cert", "", "server certificate")
	c.flagset.StringVar(&c.clientCert, "client-cert", "", "client certificate")
	c.flagset.StringVar(&c.clientKey, "client-key", "", "client key")
}

// UI returns a UI for interaction.
func (c *discoverServiceListCmd) UI() clui.UI {
	return c.ui
}

// FlagSet returns the FlagSet associated with the command. All the flags are
// parsed before running the command.
func (c *discoverServiceListCmd) FlagSet() *flagset.FlagSet {
	return c.flagset
}

// Help should return a long-form help text that includes the command-line
// usage. A brief few sentences explaining the function of the command, and
// the complete list of flags the command accepts.
func (c *discoverServiceListCmd) Help() string {
	return `
Usage: 

  services [flags] list

Description:

  List all discovery services, which are currently interned.

Example:

  therm services list
`
}

// Synopsis should return a one-line, short synopsis of the command.
// This should be short (50 characters of less ideally).
func (c *discoverServiceListCmd) Synopsis() string {
	return "List all services on the cluster."
}

// Run should run the actual command with the given CLI instance and
// command-line arguments. It should return the exit status when it is
// finished.
//
// There are a handful of special exit codes that can return documented
// behavioral changes.
func (c *discoverServiceListCmd) Run() clui.ExitCode {
	// Logging.
	var logger log.Logger
	{
		logLevel := level.AllowInfo()
		if c.debug {
			logLevel = level.AllowAll()
		}
		logger = NewLogCluiFormatter(c.UI())
		logger = log.With(logger,
			"ts", log.DefaultTimestampUTC,
			"uid", uuid.NewRandom().String(),
		)
		logger = level.NewFilter(logger, logLevel)
	}

	outputFormat := c.format
	if !contains([]string{"json", "yaml"}, outputFormat) {
		return exit(c.ui, "invalid format type (expected: json|yaml)")
	}

	client, err := getClient(c.address, certs{
		serverCert: c.serverCert,
		clientCert: c.clientCert,
		clientKey:  c.clientKey,
	}, askPassword(c.UI()), logger)
	if err != nil {
		return exit(c.ui, err.Error())
	}

	g := exec.NewGroup()
	exec.Block(g)
	{
		g.Add(func() error {
			nodes, err := client.Discovery().Services().List()
			if err != nil {
				return errors.WithStack(err)
			}
			content, err := outputContent(outputFormat, nodes)
			if err != nil {
				return errors.WithStack(err)
			}

			c.ui.Output(string(content))

			return nil
		}, func(err error) {
			// ignore
		})
	}
	exec.Interrupt(g)
	if err := g.Run(); err != nil {
		return exit(c.ui, err.Error())
	}

	return clui.ExitCode{}
}
