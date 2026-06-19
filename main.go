package main

import (
	"encoding/json"
	"monitor-go/config"
	"monitor-go/database"
	"monitor-go/models"
	"monitor-go/repository"
	"monitor-go/service"
	"net/http"
)

func main() {
	config.InitLogger()
	log := config.Log

	db, err := database.InitDB()
	if err != nil {
		log.Error("Failed to initialize database", "error", err)
		return
	}


	productRepo := repository.NewRepository[models.Product](db)
	scraper := &service.HttpScraper{}

	http.HandleFunc("GET /products", func(w http.ResponseWriter, r *http.Request) {
		products, err := productRepo.GetAll()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(products)
	})
	
	http.HandleFunc("POST /scrape", func(w http.ResponseWriter, r *http.Request){
		var body struct{URL string}
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			http.Error(w, "JSON Invalido", http.StatusBadRequest)
			return
		}

		log.Info("Iniciando scraping requisitado via API", "url", body.URL)

		product, err := scraper.Execute(body.URL)
		if err != nil {
			log.Error("Error no scraping", "url", body.URL, "erro", err)
			http.Error(w, "Falha ao raspar página", http.StatusInternalServerError)
			return
		}

		if err := productRepo.Create(product); err != nil {
			log.Error("Error ao salvar no banco", "url", body.URL, "erro", err)
			http.Error(w, "Erro ao salvar no banco", http.StatusInternalServerError)
			return
		}
		
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(product)
	})

	log.Info("Servidor rodando na porta :8080")
	http.ListenAndServe(":8080", nil)
}