package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
  "github.com/abadojack/whatlanggo"
)

func main() {
	someText, err := GetSomeText("https://www3.nhk.or.jp/news/")
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Text:", someText)
  info := whatlanggo.Detect(someText)
  fmt.Println("Language:", info.Lang.String(), "|", "ISO:", info.Lang.Iso6391(), "|", "Confidence:", info.Confidence)
}

// GetSomeText gets text using a selector and returns it.
func GetSomeText(url string) (string, error) {

	// Get the HTML
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	// Convert HTML into goquery document
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", err
	}

	// Gets first text from given element/selector
  text := doc.Find("a").First().Text()
	return text, nil
}
