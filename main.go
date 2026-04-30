package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/DWoodhouse22/gototp/cmd"
)

type Config struct {
}

func main() {
	registerCmd := flag.String("register", "", "Register secret")
	generateCmd := flag.Bool("generate", false, "Generate TOTP")
	flag.Parse()

	switch {
	case *registerCmd != "":
		cmd.Register(*registerCmd)
	case *generateCmd:
		cmd.Generate()
	default:
		fmt.Println("no valid command provided")
		flag.Usage()
		os.Exit(1)
	}
}
