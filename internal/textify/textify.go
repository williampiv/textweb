package textify

import (
	"github.com/go-shiori/go-readability"
	"log"
	"time"
)

// ByURL takes in a URL and returns the textified string data for the page
func ByURL(url string) string {
	content, err := readability.FromURL(url, time.Second*30)
	if err != nil {
		log.Fatalln(err)
	}
	return content.Content
}
