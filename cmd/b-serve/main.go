// Command line utility for b.
package main

import (
	"flag"
	"fmt"
	"syscall"

	"github.com/slogsdon/b"
	"github.com/slogsdon/b/util"
)

var (
	helpFlag   bool
	configFlag string
)

func main() {
	flag.BoolVar(&helpFlag, "h", false, "")
	flag.StringVar(&configFlag, "c", "", "")
	flag.Parse()

	command := "serve"

	args := flag.Args()
	if len(args) == 1 {
		command = args[0]
	}

	if helpFlag {
		showHelp()
		return
	}

	if configFlag != "" {
		util.ConfigPath = configFlag
	}

	switch command {
	case "help":
		showHelp()
	case "serve":
		port, ok := syscall.Getenv("PORT")
		if !ok {
			port = "3000"
		}
		options := b.Options{
			Port: port,
		}
		b.Start(options)
	case "version":
		fmt.Println("b version", b.VERSION)
	default:
		showHelp()
	}
}

func showHelp() {
	fmt.Printf(`usage: b [-h] [-c "path/to/app.config"] [command]

OPTIONS
  -h
    Show this help message
  -c
    Define the location of the app.config file.
    Defaults to './_private/app.config'.

COMMANDS
  serve
    Start the b server

  help
    Show this help message.

  version
    Show version (current version: %s)
`, b.VERSION)
}
