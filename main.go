package main

import (
	"os"
	"runtime"

	"github.com/urfave/cli"

	"github.com/elvizlai/woci/ci"
	"github.com/elvizlai/woci/config"
	"github.com/elvizlai/woci/logger"
	"github.com/elvizlai/woci/plugin/postgres"
)

const version = "2.0"

var app *cli.App

func init() {
	app = cli.NewApp()
	app.Name = "woci"
	app.HelpName = app.Name
	app.Usage = "make coding joyful!"
	app.Version = version + ", " + runtime.Version()
	app.Author = "elvizlai"
	app.EnableBashCompletion = true
}

var configFile = "/woci.yaml"
var debugMode = false
var forceMode = false

func main() {
	// global config
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "c",
			Value:       configFile,
			Usage:       "load config from `file`",
			Destination: &configFile,
		},
		cli.BoolFlag{
			Name:        "f",
			Usage:       "run in force mode",
			Destination: &forceMode,
		},
		cli.BoolFlag{
			Name:        "d",
			Usage:       "run in debug mode",
			Destination: &debugMode,
		},
	}

	app.Before = func(c *cli.Context) error {
		if debugMode {
			logger.SetDebugMode()
			logger.Debugf("Debug Mode")
		}

		if forceMode {
			logger.SetForceMode()
			logger.Warnf("Force Mode")
		}

		if c.NArg() > 0 && c.Args().First() != "plugin" {
			config.Initialize(configFile)
		}
		return nil
	}

	app.Commands = []cli.Command{
		{
			Name:    "all",
			Aliases: []string{"a"},
			Usage:   "Evil Mode.👿",
			Action: func(c *cli.Context) error {
				if !forceMode {
					logger.SetForceMode()
					logger.Warnf("Force Mode")
				}
				ci.INIT()
				ci.BEFORE()
				ci.BUILD()
				ci.START()
				ci.TEST()
				ci.CLEAN()
				return nil
			},
		},
		{
			Name:    "init",
			Aliases: []string{"i"},
			Usage:   "initial stage",
			Action: func(c *cli.Context) error {
				ci.INIT()
				return nil
			},
		},
		{
			Name:    "build",
			Aliases: []string{"b"},
			Usage:   "build stage",
			Action: func(c *cli.Context) error {
				if n := c.NArg(); n == 0 {
					return cli.NewExitError("must specify alias name or using all", 88)
				}

				ci.BEFORE()
				if c.Args().First() == "all" {
					ci.BUILD()
				} else {
					args := append([]string{c.Args().First()}, c.Args().Tail()...)
					ci.BUILD(args...)
				}
				return nil
			},
		},
		{
			Name:    "start",
			Aliases: []string{"s"},
			Usage:   "start stage",
			Action: func(c *cli.Context) error {
				if n := c.NArg(); n == 0 {
					return cli.NewExitError("must specify alias name or using all", 88)
				}

				if c.Args().First() == "all" {
					ci.START()
				} else {
					args := append([]string{c.Args().First()}, c.Args().Tail()...)
					ci.START(args...)
				}
				return nil
			},
		},
		{
			Name:    "test",
			Aliases: []string{"t"},
			Usage:   "test stage",
			Action: func(c *cli.Context) error {
				args := append([]string{c.Args().First()}, c.Args().Tail()...)
				ci.TEST(args...)
				return nil
			},
		},
		{
			Name:    "clean",
			Aliases: []string{"c"},
			Usage:   "clean stage",
			Action: func(c *cli.Context) error {
				ci.CLEAN()
				return nil
			},
		},
		{
			Name:  "plugin",
			Usage: "plugin center",
			Subcommands: []cli.Command{
				{
					Name:  "postgres",
					Usage: "postgres data importer",
					Action: func(c *cli.Context) error {
						postgres.Postgres(c.Args().First(), c.Args().Tail()...)
						return nil
					},
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		logger.Fatalf(err.Error())
	}
}
