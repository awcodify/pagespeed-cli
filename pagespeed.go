package main

import (
	"log"
	"os"

	"github.com/awcodify/pagespeed-cli/cli"
)

func main() {
	app := cli.NewApp()

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
