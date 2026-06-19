package main

import (
	"monitor-go/models"
	"testing"
)

type MockScraper struct{}

func (m *MockScraper) Execute(url string) (*models.Product, error) {
	return &models.Product{
		URL:   url,
		Title: "Mock Product",
		Price: "$9.99",
		Availability: "In Stock",
	}, nil
}

func TestScrapepComMock(t *testing.T) {
	mock := &MockScraper{}

	res, err := mock.Execute("http://example.com/product")

	if err != nil {
		t.Fatalf("Não espera erro,mas veio: %v", err)
	}

	if res.Title != "Mock Product" {
		t.Errorf("Esperava título 'Mock Product', mas veio '%s'", res.Title)
	}
}