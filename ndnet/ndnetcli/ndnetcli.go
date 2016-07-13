package ndnetcli

import (
	"fmt"
	"github.com/codegangsta/cli"
)


func NdnetCmdNotFound(c *cli.Context, command string) {
	fmt.Println(command, " not found ");

}

func NdnetInitialize(c *cli.Context) error {

	cfgFile := c.GlobalString("config")
	fmt.Println(cfgFile)
	if cfgFile != "" {
//		fmt.Println("Found config: ", cfgFile);
	}
	return nil
}

func NewCli(version string) *cli.App {
	app := cli.NewApp()
	app.Name = "ndnet"
	app.Version = version
	app.Author = "nexentaedge@nexenta.com"
	app.Usage = "CLI for nexenta edge networking"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "loglevel",
			Value:  "info",
			Usage:  "Specifies the logging level (debug|warning|error)",
			EnvVar: "LogLevel",
		},
	}
	app.CommandNotFound = NdnetCmdNotFound
	app.Before = NdnetInitialize
	app.Commands = []cli.Command{
		DaemonCmd,
		NetworkCmd,
	}
	return app
}
