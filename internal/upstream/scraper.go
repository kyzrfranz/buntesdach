package upstream

import (
	"github.com/PuerkitoBio/goquery"
	myhttp "github.com/kyzrfranz/buntesdach/internal/http"
	"net/url"
	"strings"
)

type ProfileScraper struct {
	url *url.URL
	doc *goquery.Document
}

func NewProfileScraper(docUrl *url.URL) (*ProfileScraper, error) {
	doc, err := fetchHTML(docUrl)
	if err != nil {
		return nil, err
	}
	scraper := &ProfileScraper{
		url: docUrl,
		doc: doc,
	}

	return scraper, nil
}

func fetchHTML(url *url.URL) (*goquery.Document, error) {
	data, err := myhttp.FetchUrl(url)
	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(data)))
	if err != nil {
		return nil, err
	}

	return doc, nil
}

func (s *ProfileScraper) QuerySelect(query string) *goquery.Selection {
	selection := s.doc.Find(query).Each(func(i int, selection *goquery.Selection) {
		// do nothing
	})

	return selection
}
