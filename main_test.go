package main

import (
    "testing"

	"github.com/hart87/tldr/aggregator"
)

func TestYmlElementRetrievalString(t *testing.T) {

	result := aggregator.GetElementString("app")
	if result != "TLDR" {
		t.Errorf("element.yml file is not being parsed properly")
	}
}

func TestYmlElementRetrievalSliceStringForTwitter(t *testing.T) {

	result := aggregator.GetElementSliceString("urls.twitter")
	if len(result) < 1 {
		t.Errorf("element.yml file is not being parsed properly")
	}
}

func TestYmlElementRetrievalSliceStringForReddit(t *testing.T) {

	result := aggregator.GetElementSliceString("urls.reddit")
	if len(result) < 1 {
		t.Errorf("element.yml file is not being parsed properly")
	}
}
