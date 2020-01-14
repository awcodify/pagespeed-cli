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

var (
	threshold = flag.Int("threshold", 80, "")
	key       = flag.String("key", "", "")
)

func main() {
	var WebToBeTested, strategy, format, key string
	var threshold int

	app := &cli.App{
		Action: func(c *cli.Context) error {
			WebToBeTested = c.Args().Get(0)
			return nil
		},
		Flags: []cli.Flag{
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
		},
	}

	if format != "json" || threshold != 80 || key != "" {
		log.Fatal("This feature is under development")
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

	r := pagespeed.RequestAttrs{
		URL:           "https://www.googleapis.com/pagespeedonline/v5/runPagespeed",
		WebToBeTested: WebToBeTested,
		Strategy:      strategy,
	}

	m := r.Desktop()

	s, _ := prettyjson.Marshal(m)
	fmt.Println(string(s))
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
