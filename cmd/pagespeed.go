package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/awcodify/pagespeed-cli/pagespeed"
	"github.com/hokaccha/go-prettyjson"
)

var (
	webToBeTested = flag.String("url", "", "Web URL being tested")
	category      = flag.String("category", "performance", "Category to be displayed")
)

func main() {
	flag.Parse()

	if *webToBeTested == "" {
		log.Fatal("Provide the URL!")
	}

	r := pagespeed.RequestAttrs{
		URL:           "https://www.googleapis.com/pagespeedonline/v5/runPagespeed",
		WebToBeTested: *webToBeTested,
		Category:      *category,
	}

	m := r.Desktop()

	s, _ := prettyjson.Marshal(m)
	fmt.Println(string(s))
}
