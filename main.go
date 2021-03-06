package main

import (
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Name = "Read Files"
	app.Usage = "Read multiple files in a directory"
	app.Version = "1.0.0"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "dir",
			Value: "",
			Usage: "dir for files",
		},
		cli.StringFlag{
			Name:  "output",
			Value: "",
			Usage: "Name file output",
		},
	}

	app.Action = func(c *cli.Context) error {
		dir := c.String("dir")
		output := c.String("output")

		return concatFiles(dir, output)

	}

	app.Run(os.Args)
}
