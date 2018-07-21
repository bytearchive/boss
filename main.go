package main

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "boss"
	app.Version = "1"
	app.Description = "run containers like a boss or rick ross"
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "debug",
			Usage: "enable debug output",
		},
		cli.StringFlag{
			Name:  "namespace,n",
			Usage: "containerd namespace",
			Value: "default",
		},
		cli.StringFlag{
			Name:  "log-path",
			Usage: "set the log path",
			Value: "/var/log/boss",
		},
	}
	app.Before = func(clix *cli.Context) error {
		if clix.GlobalBool("debug") {
			logrus.SetLevel(logrus.DebugLevel)
		}
		if err := os.Mkdir(clix.GlobalString("log-path"), 0755); err != nil {
			return err
		}
		return nil
	}
	if err := app.Run(os.Args); err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
}