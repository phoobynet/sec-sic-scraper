package sec_sic_scraper

import "testing"

func TestGet(t *testing.T) {
	actual, err := Get(nil)

	if err != nil {
		t.Errorf("failed to get SICs: %v", err)
	}

	if actual == nil {
		t.Errorf("expected SICs to be not nil")
	}

	if len(actual) == 0 {
		t.Errorf("expected SICs to be not empty")
	}

	t.Log("SICs count:", len(actual))
}
