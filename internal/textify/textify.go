package textify

import (
	"github.com/go-shiori/go-readability"
	"time"
)

// ByURL takes in a URL and returns the textified string data for the page
func ByURL(url string) (string, string, error) {
	content, err := readability.FromURL(url, time.Second*30)
	if err != nil {
		return "Unreadable", "Unable to Read Page", err
	}
	return content.Title, content.Content, nil
}
