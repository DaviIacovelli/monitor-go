package database

import (
	"os"
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestInitDB(t *testing.T) {
	tempDB := "monitor_test.db"
	defer os.Remove(tempDB)

	db, err := gorm.Open(sqlite.Open(tempDB), &gorm.Config{})
	if err != nil {
		t.Fatalf("Erro ao criar banco de teste: %v", err)
	}
	
	db, err = InitDB()
	if err != nil {
		t.Fatalf("Falha ao rodar InitDB: %v", err)
	}

	if db == nil {
		t.Error("Instância do DB retornou nula")
	}
}