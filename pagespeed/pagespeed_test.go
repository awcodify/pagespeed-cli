package pagespeed

import (
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestRun(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://www.googleapis.com/pagespeedonline/v5/runPagespeed",
		func(req *http.Request) (*http.Response, error) {
			return httpmock.NewJsonResponse(200, map[string]interface{}{
				"id": "https://www.google.com/",
			})
		},
	)

	r := RequestAttrs{
		URL:           "https://www.googleapis.com/pagespeedonline/v5/runPagespeed",
		WebToBeTested: "https://www.google.com/",
		Strategy:      "mobile",
	}

	m, err := r.Run()

	if m.ID != "https://www.google.com/" {
		t.Errorf("We are testing %s, but %s returned", "https://www.google.com/", m.ID)
	}

	if err != nil {
		t.Errorf("%s", err)
	}

	r.WebToBeTested = "asdf"
	m, err = r.Run()

	if m.ID != "" {
		t.Errorf("It should be nil but found %s", m.ID)
	}

	if err == nil {
		t.Errorf("%s", err)
	}

	r.URL = "asdf"
	r.WebToBeTested = "https://google.com"
	m, err = r.Run()

	if err == nil {
		t.Errorf("It should be invalid URI for request")
	}
}

func TestScorePercentum(t *testing.T) {
	m := Metric{
		LighthouseResult: lighthouseResult{
			Categories: categories{
				Performance: performance{
					Score: 0.9,
				},
			},
		},
	}

	s := m.ScorePercentum()

	if s != float32(90.00) {
		t.Errorf("It should 90.0, but get %f", s)
	}
}
