package reddit

import (
    "testing"
	"fmt"
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

func TestYmlElementRetrievalInt(t *testing.T) {

	result := GetElementInt("num")
	fmt.Println(result)
	if result != 33567 {
		t.Errorf("Reddit element.yml file is not being parsed properly")
	}
}

func TestYmlElementRetrievalString(t *testing.T) {

	result := GetElementString("post")
	fmt.Println(result)
	if result != "a post from an element" {
		t.Errorf("Reddit element.yml file is not being parsed properly")
	}
}