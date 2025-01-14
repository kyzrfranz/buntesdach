package upstream

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/stretchr/testify/assert"
	"net/url"
	"strings"
	"testing"
)

const rawUrl = "https://www.bundestag.de/abgeordnete/biografien/A/abdi_sanae-861028"

func TestScraper(t *testing.T) {
	url, err := url.Parse(rawUrl)
	assert.NoError(t, err)
	scraper, err := NewProfileScraper(url)
	assert.NoError(t, err)

	profil := scraper.QuerySelect("div.bt-profil")
	assert.NotEmpty(t, profil)
}

func TestScraperLinks(t *testing.T) {
	url, err := url.Parse(rawUrl)
	assert.NoError(t, err)
	scraper, err := NewProfileScraper(url)
	assert.NoError(t, err)

	links := scraper.QuerySelect("div.bt-profil-kontakt ul.bt-linkliste a")
	assert.NotEmpty(t, links)

	links.Each(func(i int, selection *goquery.Selection) {
		href, exists := selection.Attr("href")
		if !exists {
			return
		}

		// Print parsed URL details
		fmt.Printf("Link %s: %s \n", strings.Trim(selection.Text(), "\n"), href)
	})
}
