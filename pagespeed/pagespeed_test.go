package pagespeed

import (
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestDesktop(t *testing.T) {
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

	m := r.Desktop()

	if m.ID != "https://www.google.com/" {
		t.Errorf("We are testing %s, but %s returned", "https://www.google.com/", m.ID)
	}

	r.WebToBeTested = "asdf"
	m = r.Desktop()

	if m.ID != "" {
		t.Errorf("It should be error but found %s", m.ID)
	}
}
