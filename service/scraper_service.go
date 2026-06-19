package service

import (
	"monitor-go/models"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Scraper interface {
	Execute(url string) (*models.Product, error)
}

type HttpScraper struct{}

func (h *HttpScraper) Execute(url string) (*models.Product, error) {
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	return &models.Product{
		URL:   url,
		Title: strings.TrimSpace(doc.Find("h1").First().Text()),
		Price: strings.TrimSpace(doc.Find(".price_color").First().Text()),
		Availability: strings.TrimSpace(doc.Find(".instock").First().Text()),
	}, nil

}