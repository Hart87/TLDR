package reddit

import (
    "testing"
)

func TestHelloMatch(t *testing.T ) {
	
	result := RedditHello()
	if result != "Doing stuff from reddit..." {
		t.Errorf("Reddit hello function returned an accurate string. supposed to be %v", result)
	}

}

func TestHelloNotNull(t *testing.T) {

	result := RedditHello()
	if result == "" {
		t.Errorf("Reddit hello function returned an empty string. supposed to be %v", result)
	}

}