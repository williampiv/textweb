package textify

import (
	"strings"
	"time"

	"github.com/go-shiori/go-readability"
)

// HTMLReplacement contains the string to be replaced, with its replacement string
type HTMLReplacement struct {
	ReplaceString     string
	ReplacementString string
}

// ByURL takes in a URL and returns the textified string data for the page
func ByURL(url string) (string, string, error) {
	content, err := readability.FromURL(url, time.Second*30)
	if err != nil {
		return "Unreadable", "Unable to Read Page", err
	}
	return content.Title, content.Content, nil
}

// ReplaceHTMLContent takes in a pointer to web content (string) and replaces the
// ReplaceString from an HTMLReplacement struct with its ReplacementString
func ReplaceHTMLContent(webContent *string) {
	replacements := []HTMLReplacement{
		{"href=\"", "href=\"/text?url="},
		{"src=", "width=\"25%\" src="},
		{" target=\"_blank\"", ""},
	}
	for _, i := range replacements {
		*webContent = strings.Replace(*webContent, i.ReplaceString, i.ReplacementString, -1)
	}
}
