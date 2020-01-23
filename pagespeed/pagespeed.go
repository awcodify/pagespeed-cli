package pagespeed

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// Metric is metric for page speed
type Metric struct {
	ID               string `json:"id"`
	Strategy         string
	LighthouseResult LighthouseResult `json:"LighthouseResult"`
}

// LighthouseResult is
type LighthouseResult struct {
	Audits Audits `json:"audits"`
}

// Audits is
type Audits struct {
	FirstContentfulPaint Score `json:"first-contentful-paint"`
	SpeedIndex           Score `json:"speed-index"`
	TimeToInteractive    Score `json:"interactive"`
}

// Score is
type Score struct {
	Title            string
	Description      string
	ScoreDisplayMode string
	Score            float32
	DisplayValue     string
	numericValue     int32
}

// RequestAttrs is collection of attributes for http call to google api
type RequestAttrs struct {
	URL           string
	WebToBeTested string
	Strategy      string
}

// Run is for calculating page speed
func (r RequestAttrs) Run() (Metric, error) {
	m := Metric{
		Strategy: r.Strategy,
	}

	u, err := url.ParseRequestURI(r.WebToBeTested)
	if err != nil {
		return Metric{}, err
	}

	url := fmt.Sprintf("%s?url=%s&strategy=%s", r.URL, u, r.Strategy)

	if errHTTP := getJSON(url, &m); errHTTP != nil {
		return Metric{}, errHTTP
	}

	return m, nil
}

var myClient = &http.Client{Timeout: 10 * time.Second}

func getJSON(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}
