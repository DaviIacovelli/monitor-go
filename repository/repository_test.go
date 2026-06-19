package repository

import (
	"monitor-go/models"
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestGenericRepository(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Erro ao abrir banco em memória: %v", err)
	}

	db.AutoMigrate(&models.Product{})


	repo := NewRepository[models.Product](db)


	product := &models.Product{
		URL:          "http://teste.com",
		Title:        "Produto Teste",
		Price:        "£9.99",
		Availability: "Em estoque",
	}

	err = repo.Create(product)
	if err != nil {
		t.Errorf("Não esperava erro ao criar, mas veio: %v", err)
	}

	
	products, err := repo.GetAll()
	if err != nil {
		t.Errorf("Não esperava erro ao buscar, mas veio: %v", err)
	}

	if len(products) != 1 {
		t.Errorf("Esperava 1 produto no banco, encontrei %d", len(products))
	}

	if products[0].Title != "Produto Teste" {
		t.Errorf("Esperava o título 'Produto Teste', veio '%s'", products[0].Title)
	}
}