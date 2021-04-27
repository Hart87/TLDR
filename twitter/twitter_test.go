package twitter

import (
    "testing"
)

func TestHelloMatch(t *testing.T ) {
	
	result := TwitterHello()
	if result != "Doing stuff from twitter..." {
		t.Errorf("Twitter hello function returned an accurate string. supposed to be %v", result)
	}

}

func TestHelloNotNull(t *testing.T) {

	result := TwitterHello()
	if result == "" {
		t.Errorf("Twitter hello function returned an empty string. supposed to be %v", result)
	}

}