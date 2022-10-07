package cmd

import (
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

// LogLevelArg is the CLI argument for the logging level
const LogLevelArg string = "log-level"

// LogLevelFlag is the urfave/cli Flag configuration for the logging level
var LogLevelFlag = &cli.StringFlag{
	Name:    LogLevelArg,
	Usage:   "Log level (trace, debug, info, warn, error)",
	EnvVars: []string{"LOG_LEVEL"},
	Value:   "info",
}

// LogFormatArg is the CLI argument for the logging format
const LogFormatArg string = "log-format"

// LogFormatFlag is the urfave/cli Flag configuration for the logging format
var LogFormatFlag = &cli.StringFlag{
	Name:    LogFormatArg,
	Usage:   "Log format (plain, json)",
	EnvVars: []string{"LOG_FORMAT"},
	Value:   "plain",
}

// PrepareLogger prepares the global logger with options from CLI args
func PrepareLogger(ctx *cli.Context) {
	if logLevel, err := log.ParseLevel(ctx.String(LogLevelArg)); err == nil {
		log.SetLevel(logLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}

	if strings.ToUpper(strings.TrimSpace(ctx.String(LogFormatArg))) == "JSON" {
		log.SetFormatter(&log.JSONFormatter{})
	}
}
