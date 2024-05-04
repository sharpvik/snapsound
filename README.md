# `snapsound`

Break out of the Matrix - play audio from PNG files!

Don't think that's possible? Try `snapsound`!

## Install

```bash
go install github.com/sharpvik/snapsound
```

## Grab an Image

```bash
wget https://raw.githubusercontent.com/sharpvik/snapsound/main/gallery/Never%20Gonna%20Give%20You%20Up.png
```

... or make your own:

```bash
snapsound snap some.mp3
```

You should see a PNG image appear in the same folder (e.g. `some.png`).

## Play it!

```bash
snapsound play some.png
```

## `snapsound help`

```text
NAME:
   snapsound - Play pixels

USAGE:
   snapsound [global options] command [command options]

VERSION:
   0.1.0

COMMANDS:
   snap     Snap that shot
   sound    Make it yours
   play     Play it now
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```
