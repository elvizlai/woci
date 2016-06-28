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
	"sync"
	"time"

	"github.com/urfave/cli"
	"github.com/wothing/log"

	"github.com/wothing/woci/base"
	"github.com/wothing/woci/ci"
	"github.com/wothing/woci/conf"
)

var wg = &sync.WaitGroup{}

var gofunc = func(foo func()) {
	defer wg.Done()
	foo()
}

var debug = false
var configFile = "/woci.json"

func init() {
	log.SetFlags(log.LstdFlags | log.Llevel)
}

func main() {
	app := cli.NewApp()
	app.Name = "woci"
	app.HelpName = app.Name
	app.Usage = "make coding joyful!"
	app.Version = "0.1"
	app.Author = "elvizlai"
	app.EnableBashCompletion = true

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:        "debug, d",
			Usage:       "debug mode",
			Destination: &debug,
		},
		cli.StringFlag{
			Name:        "config, c",
			Value:       "/woci.json",
			Usage:       "load config from `file`",
			Destination: &configFile,
		},
	}

	app.Before = func(c *cli.Context) error {
		if debug {
			log.SetOutputLevel(log.Ldebug)
		}
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
				ci.Etcd()
				wg.Add(3)
				go gofunc(ci.Pgsql)
				go gofunc(ci.Redis)
				go gofunc(ci.Nsq)
				wg.Wait()

				ci.AppBuild()

				start := time.Now()
				ci.AppStart()
				<-time.After(time.Now().Add(time.Second * 5).Sub(start))

				ci.AppTest()

				return nil
			},
		},
		cli.Command{
			Name:      "init",
			ShortName: "i",
			Usage:     "Init etcd, postgres, redis, nsq",
			Action: func(c *cli.Context) error {
				conf.GenUUID()
				ci.Etcd()
				wg.Add(3)
				go gofunc(ci.Pgsql)
				go gofunc(ci.Redis)
				go gofunc(ci.Nsq)
				wg.Wait()
				return nil
			},
		},
		cli.Command{
			Name:    "build",
			Aliases: []string{"b"},
			Usage:   "Build all app and start",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "only, o",
					Value: "",
					Usage: "only build app by `name`",
				},
			},
			Action: func(c *cli.Context) error {
				if arg := c.String("only"); arg != "" {
					temp := []conf.Service{}
					buildList := strings.Split(arg, ",")
					for _, x := range buildList {
						for _, y := range conf.Services {
							if x == y.Name {
								temp = append(temp, y)
							}
						}
					}
					conf.Services = temp
					ci.AppBuild()
					return nil
				}
				conf.RestoreUUID()
				ci.AppBuild()
				ci.AppStart()
				return nil
			},
		},
		cli.Command{
			Name:    "rebuild",
			Aliases: []string{"r", "re"},
			Usage:   "Rebuild app. split by ',' if multiple",
			Action: func(c *cli.Context) error {
				conf.RestoreUUID()
				args := c.Args().First()
				switch args {
				case "":
					return cli.NewExitError("please provide app name", 88)
				case "all":
				default:
					temp := []conf.Service{}
					buildList := strings.Split(args, ",")
					for _, x := range buildList {
						for _, y := range conf.Services {
							if x == y.Name {
								temp = append(temp, y)
							}
						}
					}
					conf.Services = temp
				}
				ci.AppClean()
				ci.AppBuild()
				ci.AppStart()
				return nil
			},
		},
		cli.Command{
			Name:    "test",
			Aliases: []string{"t"},
			Usage:   "Run test case",
			Action: func(c *cli.Context) error {
				conf.RestoreUUID()
				ci.AppTest()
				return nil
			},
		},
		cli.Command{
			Name:    "clean",
			Aliases: []string{"c"},
			Usage:   "Clean all app and data",
			Action: func(c *cli.Context) error {
				conf.RestoreUUID()
				ci.AppClean()
				ci.DataClean()
				return nil
			},
		},
		cli.Command{
			Name:    "logs",
			Aliases: []string{"l"},
			Usage:   "Log by app name",
			Action: func(c *cli.Context) error {
				if c.NArg() == 0 {
					return cli.NewExitError("please provide app name", 88)
				}
				conf.RestoreUUID()

				args := c.Args().First()
				if args == "all" {
					for _, v := range conf.Services {
						data, err := base.CMD("docker logs " + conf.Tracer + "-" + v.Name)
						if err != nil {
							log.Terrorf(conf.Tracer, data)
							log.Tfatal(conf.Tracer, err)
						}
						fmt.Fprintf(app.Writer, data)
					}
				} else {
					appList := strings.Split(args, ",")
					for _, v := range appList {
						data, err := base.CMD("docker logs " + conf.Tracer + "-" + v)
						if err != nil {
							log.Terrorf(conf.Tracer, data)
							log.Tfatal(conf.Tracer, err)
						}
						fmt.Fprintf(app.Writer, data)
					}
				}
				return nil
			},
		},
	}

	app.CommandNotFound = func(c *cli.Context, s string) {
		fmt.Println("COMMANDS '" + s + "' NOT FOUND")
	}

	app.Run(os.Args)
}
