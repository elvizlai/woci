/**
 * Copyright 2015-2016, Wothing Co., Ltd.
 * All rights reserved.
 *
 * Created by Elvizlai on 2016/04/23 10:42
 */

package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/urfave/cli"

	"github.com/wothing/woci/ci"
	"github.com/wothing/woci/conf"
	"github.com/wothing/woci/plugin"
	"github.com/wothing/woci/util/log"
)

var configFile = "/woci.json"

func main() {
	app := cli.NewApp()
	app.Name = "woci"
	app.HelpName = app.Name
	app.Usage = "make coding joyful!"
	app.Version = "0.3"
	app.Author = "elvizlai"
	app.EnableBashCompletion = true

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "config, c",
			Value:       "/woci.json",
			Usage:       "load config from `file`",
			Destination: &configFile,
		},
		cli.BoolFlag{
			Name:        "debug, d",
			Usage:       "debug mode",
			Destination: &log.DebugMode,
		},
		cli.BoolFlag{
			Name:        "force, f",
			Usage:       "force mode",
			Destination: &log.ForceMode,
		},
	}

	app.Before = func(c *cli.Context) error {
		log.Initial()
		if c.NArg() > 0 {
			conf.ParseConfig(configFile)
		}
		return nil
	}

	app.Commands = []cli.Command{
		cli.Command{
			Name:    "all",
			Aliases: []string{"a"},
			Usage:   "Evil Mode.👿",
			Action: func(c *cli.Context) error {
				conf.GenUUID()
				lifecycle()
				return nil
			},
		},
		cli.Command{
			Name:      "init",
			ShortName: "i",
			Usage:     "Initial stage",
			Action: func(c *cli.Context) error {
				conf.GenUUID()
				ci.Initial()
				return nil
			},
		},
		cli.Command{
			Name:    "build",
			Aliases: []string{"b"},
			Usage:   "Build stage",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "only, o",
					Value: "",
					Usage: "only build app by `name`",
				},
			},
			Action: func(c *cli.Context) error {
				if arg := c.String("only"); arg != "" {
					temp := []conf.Module{}
					for _, m := range strings.Split(arg, ",") {
						for _, v := range conf.Config.Modules {
							if m == v.Name {
								temp = append(temp, v)
							}
						}
					}
					conf.Config.Modules = temp
					ci.Build()
					return nil
				}
				conf.RestoreUUID()
				ci.Build()
				ci.Start()
				return nil
			},
		},
		cli.Command{
			Name:    "rebuild",
			Aliases: []string{"r", "re"},
			Usage:   "Rebuild stage",
			Action: func(c *cli.Context) error {
				conf.RestoreUUID()
				n := c.NArg()

				if n == 0 {
					return cli.NewExitError("please specify app name", 88)
				}

				if c.Args().First() != "all" {
					temp := []conf.Module{}
					for i := 0; i < n; i++ {
						for _, m := range strings.Split(c.Args().Get(i), ",") {
							for _, v := range conf.Config.Modules {

								if m == v.Name {
									temp = append(temp, v)
								}
							}
						}
					}
					conf.Config.Modules = temp
				}

				ci.Clean()
				ci.Build()
				ci.Start()
				return nil
			},
		},
		cli.Command{
			Name:    "test",
			Aliases: []string{"t"},
			Usage:   "Run all test case if no app name specified",
			Action: func(c *cli.Context) error {
				conf.RestoreUUID()
				args := []string{c.Args().First()}
				args = append(args, c.Args().Tail()...)
				ci.Test(strings.Split(strings.Join(args, ","), ",")...)
				return nil
			},
		},
		cli.Command{
			Name:    "clean",
			Aliases: []string{"c"},
			Usage:   "Clean all app and data",
			Action: func(c *cli.Context) error {
				conf.RestoreUUID()
				ci.After()
				return nil
			},
		},
		cli.Command{
			Name:  "plugin",
			Usage: "Plugin center",
			Subcommands: []cli.Command{
				cli.Command{
					Name:  "postgres",
					Usage: "Postgres data importer",
					Action: func(c *cli.Context) error {
						plugin.Postgres(c.Args().First(), c.Args().Tail()...)
						return nil
					},
				},
			},
		},
		cli.Command{
			Name:    "logs",
			Aliases: []string{"l"},
			Usage:   "Fetch the logs of specified app",
			Action: func(c *cli.Context) error {
				if c.NArg() == 0 {
					return cli.NewExitError("please provide app name", 88)
				}
				conf.RestoreUUID()

				//args := c.Args().First()
				//if args == "all" {
				//	for _, v := range conf.Services {
				//		data, err := base.CMD("docker logs " + conf.Tracer + "-" + v.Name)
				//		if err != nil {
				//			log.Terrorf(conf.Tracer, data)
				//			log.Tfatal(conf.Tracer, err)
				//		}
				//		fmt.Fprintf(app.Writer, data)
				//	}
				//} else {
				//	appList := strings.Split(args, ",")
				//	for _, v := range appList {
				//		data, err := base.CMD("docker logs " + conf.Tracer + "-" + v)
				//		if err != nil {
				//			log.Terrorf(conf.Tracer, data)
				//			log.Tfatal(conf.Tracer, err)
				//		}
				//		fmt.Fprintf(app.Writer, data)
				//	}
				//}
				return nil
			},
		},
	}

	app.CommandNotFound = func(c *cli.Context, s string) {
		fmt.Println("COMMANDS '" + s + "' NOT FOUND")
	}

	app.Run(os.Args)
}

func lifecycle() {
	ci.Initial()
	ci.Build()
	ci.Start()
	ci.Test()
}
