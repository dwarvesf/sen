package main

import (
	"os"

	"github.com/dwarvesf/sen/cmd"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "sen"
	app.Version = "1.0"
	app.Usage = "A small cli written in Go to help automation test"
	app.Email = "dev@dwarvesf.com"
	app.Action = cmd.Action
	app.Flags = cmd.Flags
	app.Run(os.Args)
}
