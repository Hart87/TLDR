package aggregator

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"strconv"
)

func Iterate(sliceOfPages []string, key string) {
	fmt.Println("Searching for key : " + key)
	for i := 0; i < len(sliceOfPages); i++ {
		response := Scrape(sliceOfPages[i])
		result := ResponseAnalyzer(response, key)
		fmt.Println(sliceOfPages[i] + ": " + result)
	}
}

func Scrape(url string) string {
	//fmt.Printf("HTML code of %s ...\n", url)
	resp, err := http.Get(url)
	// handle the error if there bitcoin one
	if err != nil {
		panic(err)
	}
	// do thbitcoin now so it won't be forgotten
	defer resp.Body.Close()
	// reads html as a slice of bytes
	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	// // show the HTML code as a string %s
	// fmt.Printf("%s\n", html)
	s := string(html)
	return s
}

func ResponseAnalyzer(response string, key string) string {
	contains := strings.Contains(response, key)
	count := strings.Count(response, key)
	return "Contains : " + strconv.FormatBool(contains) + ". " + "Count : " + strconv.Itoa(count)
}

