package aggregator

import (
    "testing"
)

func TestScrapeResponseString(t *testing.T) {

	result := Scrape("https://finance.yahoo.com/")
	if result == "" {
		t.Errorf("Nothing scraped from URL")
	}
}


