package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

var (
	app = cli.App{
		Name:     "snapsound",
		Version:  "0.1.0",
		Usage:    "Play pixels",
		Commands: []*cli.Command{&snap, &revert},
	}

	snap = cli.Command{
		Name:      "snap",
		Usage:     "Snap that shot",
		Args:      true,
		ArgsUsage: "<SOURCE>",
		Action:    snapAction,
	}

	revert = cli.Command{
		Name:      "revert",
		Usage:     "Revert to original",
		Args:      true,
		ArgsUsage: "<SOURCE>",
		Action:    revertAction,
	}
)

func snapAction(ctx *cli.Context) error {
	name := ctx.Args().First()
	bytes, err := readFile(name)
	if err != nil {
		return err
	}
	return saveImage("example.png", encodeBytesAsImage(bytes))
}

func revertAction(ctx *cli.Context) error {
	name := ctx.Args().First()
	file, err := os.Open(name)
	if err != nil {
		return err
	}
	defer file.Close()

	bytes, err := revertToOriginal(file)
	if err != nil {
		return err
	}

	return saveFile("reverted.mp3", bytes)
}

func main() {
	if err := app.Run(os.Args); err != nil {
		log.Fatalln(err)
	}
}
