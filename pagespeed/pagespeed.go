package pagespeed

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// Metric will collect the result of test from Google API
type Metric struct {
	ID               string `json:"id"`
	Strategy         string
	LighthouseResult lighthouseResult `json:"LighthouseResult"`
}

type lighthouseResult struct {
	Audits     audits     `json:"audits"`
	Categories categories `json:"categories"`
}

type audits struct {
	FirstContentfulPaint score `json:"first-contentful-paint"`
	SpeedIndex           score `json:"speed-index"`
	TimeToInteractive    score `json:"interactive"`
}

type categories struct {
	Performance performance `json:"performance"`
}

type performance struct {
	Score float32 `json:"score"`
}

// Score is for showing the result in number
type score struct {
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

const baseURL = "https://www.googleapis.com/pagespeedonline/v5/runPagespeed"

// Run is for calculating page speed
func (r RequestAttrs) Run() (Metric, error) {
	m := Metric{
		Strategy: r.Strategy,
	}

	u, err := url.ParseRequestURI(r.WebToBeTested)
	if err != nil {
		return Metric{}, err
	}

	url := fmt.Sprintf("%s?url=%s&strategy=%s", baseURL, u, r.Strategy)

	if errHTTP := getJSON(url, &m); errHTTP != nil {
		return Metric{}, errHTTP
	}

	return m, nil
}

// ScorePercentum will convert score to percentum
func (m Metric) ScorePercentum() float32 {
	return m.LighthouseResult.Categories.Performance.Score * 100
}

var myClient = &http.Client{Timeout: 30 * time.Second}

func getJSON(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}
