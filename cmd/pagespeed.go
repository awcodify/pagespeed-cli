package main

import (
	"fmt"
	"log"
	"os"

	"github.com/awcodify/pagespeed-cli/pagespeed"
	"github.com/fatih/color"
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

		if format != "json" || key != "" {
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

		score := m.ScorePercentum()
		thr := float32(threshold)
		edge := ((100 - thr) / 4) + thr
		var msg string

		fmt.Println(thr)
		if score < thr {
			msg = fmt.Sprintf("FAILED: Your score is %.2f, threshold %.2f!", score, thr)
			color.Red(msg)
		} else if score < edge {
			msg = fmt.Sprintf("PASSED: Your score is %.2f, threshold %.2f!", score, thr)
			color.Yellow(msg)
		} else {
			msg = fmt.Sprintf("PASSED: Your score is %.2f, threshold %.2f!", score, thr)
			color.Green(msg)
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
