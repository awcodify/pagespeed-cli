package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/awcodify/pagespeed-cli/pagespeed"
	"github.com/hokaccha/go-prettyjson"
	"github.com/urfave/cli/v2"
)

func main() {
	var webToBeTested, strategy, format, key string
	var threshold int

	app := cli.NewApp()

	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:        "strategy",
			Value:       "desktop",
			Usage:       "Strategy to use when analyzing the page: mobile|desktop",
			Destination: &strategy,
		},
		&cli.StringFlag{
			Name:        "format",
			Value:       "json",
			Usage:       "Output format: cli|json",
			Destination: &format,
		},
		&cli.StringFlag{
			Name:        "key",
			Usage:       "Google API Key. By default the free tier is used",
			Destination: &key,
		},

		&cli.IntFlag{
			Name:        "threshold",
			Value:       80,
			Usage:       "Threshold score to pass the PageSpeed test",
			Destination: &threshold,
		},
	}

	app.Action = func(c *cli.Context) error {
		if c.NArg() > 0 {
			webToBeTested = c.Args().Get(0)
		} else {
			fmt.Println("Please provide the URL to be tested!")
			cli.ShowAppHelp(c)
			return nil
		}

		if format != "json" || threshold != 80 || key != "" {
			log.Fatal("This feature is under development")
		}

		r := pagespeed.RequestAttrs{
			URL:           "https://www.googleapis.com/pagespeedonline/v5/runPagespeed",
			WebToBeTested: webToBeTested,
			Strategy:      strategy,
		}

		m, err := r.Run()
		if err != nil {
			return err
		}

		s, _ := prettyjson.Marshal(m)
		fmt.Println(string(s))

		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func printUsage() {
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println(" ", "$", os.Args[0], "URL", "OPTIONS")
	fmt.Println()

	fmt.Println("Options:")
	flag.PrintDefaults()

	fmt.Println("Example")
	fmt.Println(" ", "$", os.Args[0], "--strategy=mobile")
	os.Exit(1)
}
