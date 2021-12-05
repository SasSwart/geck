package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	git "github.com/libgit2/git2go/v33"
)

func main() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Could not determine user home directory.")
		fmt.Println(err)
		os.Exit(1)
	}

	if len(os.Args) < 2 {
		fmt.Println("Geck expects a subcommand. See geck help")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "init":
		initCommand := flag.NewFlagSet("init", flag.ExitOnError)
		initPath := initCommand.String("path", homeDir+"/geck/", "The local path in which geck will store its repository. This directory will be created if it does not exist.")
		initCommand.Parse(os.Args[2:])

		_, err = git.InitRepository(*initPath+".git", true)
		if err != nil {
			fmt.Println("Could not create git repository")
			fmt.Println(err)
			os.Exit(1)
		}
	case "clone":
		cloneCommand := flag.NewFlagSet("init", flag.ExitOnError)
		cloneSource := cloneCommand.String("repo", "https://example.com", "The git repository to clone")
		clonePath := cloneCommand.String("path", homeDir+"/geck/", "The local path in which geck will store its repository. This directory will be created if it does not exist.")
		cloneCommand.Parse(os.Args[2:])

		// For now, this only works for HTTPS
		// TODO: implement host key checking callback for ssh
		_, err := git.Clone(*cloneSource, *clonePath, &git.CloneOptions{
			CheckoutBranch: "main",
			CheckoutOptions: git.CheckoutOptions{
				Strategy: git.CheckoutForce,
			},
		})

		if err != nil {
			fmt.Println("Could not clone git repository")
			fmt.Println(err)
			os.Exit(1)
		}
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
