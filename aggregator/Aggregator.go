package aggregator

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	//"github.com/spf13/viper"
)

func Recon (key string) {
	result := ScrapeStuff("https://finance.yahoo.com/")
	//fmt.Println(result)
	fmt.Println(strings.Contains(result, key))
	fmt.Println(strings.Count(result, key))
	
}

func ScrapeStuff(urlString string) string {

	url := urlString
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