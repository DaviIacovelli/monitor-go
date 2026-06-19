package service

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHttpScraper_Execute(t *testing.T) {
	// Cria um servidor HTTP local simulado
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Cospe um HTML estático simulando o site de livros
		fmt.Fprintln(w, `
			<html>
				<head><title>Título de Fallback</title></head>
				<body>
					<h1>Livro Ultra Raro</h1>
					<p class="price_color">£42.00</p>
					<p class="instock">In stock (12 available)</p>
				</body>
			</html>
		`)
	}))
	defer server.Close() // Fecha o servidor no final do teste

	scraper := &HttpScraper{}

	// Passamos a URL do nosso servidor simulado (ex: http://127.0.0.1:xxxxx)
	product, err := scraper.Execute(server.URL)
	if err != nil {
		t.Fatalf("Erro inesperado no scraping: %v", err)
	}

	if product.Title != "Livro Ultra Raro" {
		t.Errorf("Esperava 'Livro Ultra Raro', veio '%s'", product.Title)
	}

	if product.Price != "£42.00" {
		t.Errorf("Esperava '£42.00', veio '%s'", product.Price)
	}
}