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
		Commands: []*cli.Command{&snap, &sound, &play},
	}

	snap = cli.Command{
		Name:      "snap",
		Usage:     "Snap that shot",
		Args:      true,
		ArgsUsage: "<SOURCE>",
		Action:    snapAction,
	}

	sound = cli.Command{
		Name:      "sound",
		Usage:     "Make it yours",
		Args:      true,
		ArgsUsage: "<SOURCE>",
		Action:    soundAction,
	}

	play = cli.Command{
		Name:      "play",
		Usage:     "Play it now",
		Args:      true,
		ArgsUsage: "<SOURCE>",
		Action:    playAction,
	}
)

func snapAction(ctx *cli.Context) error {
	name := ctx.Args().First()
	bytes, err := readFile(name)
	if err != nil {
		return err
	}
	return saveImage(trimExtension(name)+".png", encodeBytesAsImage(bytes))
}

func soundAction(ctx *cli.Context) error {
	name := ctx.Args().First()
	bytes, err := originalBytesFromFile(name)
	if err != nil {
		return err
	}
	return saveFile(trimExtension(name)+" (Reverted).mp3", bytes)
}

func playAction(ctx *cli.Context) error {
	name := ctx.Args().First()
	bytes, err := originalBytesFromFile(name)
	if err != nil {
		return err
	}
	return playBytes(bytes)
}

func main() {
	if err := app.Run(os.Args); err != nil {
		log.Fatalln(err)
	}
}
