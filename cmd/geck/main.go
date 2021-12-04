package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("expected a subcommand. See geck help")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "init":
		initCommand := flag.NewFlagSet("init", flag.ExitOnError)
		initDirectory := initCommand.String("directory", "~/geck/", "The local directory in which geck will store its repository. This directory will be created if it does not exist.")
		initCommand.Parse(os.Args[2:])
	case "clone":
		log.Println("geck clone not implemented yet")
		os.Exit(1)
	case "set":
		log.Println("geck set not implemented yet")
		os.Exit(1)
	case "file":
		log.Println("geck file not implemented yet")
		os.Exit(1)
	case "user":
		log.Println("geck user not implemented yet")
		os.Exit(1)
	case "group":
		log.Println("geck group not implemented yet")
		os.Exit(1)
	case "plugin":
		log.Println("geck plugin not implemented yet")
		os.Exit(1)
	}
}
