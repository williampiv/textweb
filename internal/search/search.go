package search

import (
	"fmt"
	"github.com/anaskhan96/soup"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// ResultURL is the url & title that is then presented as links
type ResultURL struct {
	URL   string
	Title string
	Short string
}

// GetSearchResults retrieves search results from DuckDuckGo and then returns them formatted for textWeb
func GetSearchResults(searchString string) ([]ResultURL, error) {
	finalResults := make([]ResultURL, 0)
	searchString = url.QueryEscape(searchString)
	// TODO: Break out request to DDG into own function, then process
	req, err := http.NewRequest("GET", fmt.Sprintf("https://html.duckduckgo.com/html?q=%s", searchString), nil)
	if err != nil {
		log.Println(err)
		return []ResultURL{ResultURL{URL: "/", Title: "No results found", Short: "No results found"}}, err
	}
	// Set User-Agent since DDG sometimes doesn't want to provide results to Golang default agent
	req.Header.Set("User-Agent",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.114 Safari/537.36",
	)
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return []ResultURL{ResultURL{URL: "/", Title: "No results found", Short: "No results found"}}, err
	}
	responseData, _ := ioutil.ReadAll(resp.Body)
	if err = resp.Body.Close(); err != nil {
		log.Println(err)
	}
	pageDOM := soup.HTMLParse(string(responseData))
	results := pageDOM.FindAll("div", "class", "result__body")
	for _, y := range results {
		// TODO: This could use some better error handling in the event an element cannot be found
		link := y.Find("a", "class", "result__a")
		linkShort := y.Find("a", "class", "result__snippet")
		log.Printf("%v", linkShort)
		if strings.Contains(link.Attrs()["href"], "https://duckduckgo.com/y.js") {
			continue
		}
		href := strings.Replace(link.Attrs()["href"], "//duckduckgo.com/l/?uddg=", "", -1)
		formedURL := ResultURL{URL: href, Title: link.Text(), Short: linkShort.Text()}
		finalResults = append(finalResults, formedURL)
	}
	return finalResults, nil
}
