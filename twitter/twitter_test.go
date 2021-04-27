package twitter

import (
    "testing"
	"fmt"
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

func TestYmlElementRetrievalInt(t *testing.T) {

	result := GetElementInt("num")
	fmt.Println(result)
	if result != 42069 {
		t.Errorf("Twitter element.yml file is not being parsed properly")
	}
}

func TestYmlElementRetrievalString(t *testing.T) {

	result := GetElementString("post")
	fmt.Println(result)
	if result != "a post from a Twitter element" {
		t.Errorf("Twitter element.yml file is not being parsed properly")
	}
}