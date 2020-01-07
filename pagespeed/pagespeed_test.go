package pagespeed

import (
	"testing"
)

func TestDesktop(t *testing.T) {
	r := RequestAttrs{
		URL:           "https://www.googleapis.com/pagespeedonline/v5/runPagespeed",
		WebToBeTested: "https://www.google.com/",
		Category:      "performance",
	}

	m := r.Desktop()

	if m.ID != "https://www.google.com/" {
		t.Errorf("We are testing %s, but %s returned", "https://www.google.com/", m.ID)
	}
}
