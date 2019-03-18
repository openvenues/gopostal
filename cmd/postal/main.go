package main

import (
	"fmt"
	"os"

	"github.com/skyline-ai/postal"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "postal"
	app.Version = "0.0.1"
	app.Usage = "postal cli"

	app.Commands = []cli.Command{
		{
			Name:      "comps",
			Aliases:   []string{"c"},
			Usage:     "parse address components",
			UsageText: "postal comps <address>",
			Action: func(c *cli.Context) error {
				address := c.Args().Get(0)

				comps := postal.ParseAddress(address, postal.DefaultParserOptions())
				for _, comp := range comps {
					fmt.Printf("%s: '%s'\n", comp.Label, comp.Value)
				}

				return nil
			},
		},
	}

	app.Run(os.Args)
}
