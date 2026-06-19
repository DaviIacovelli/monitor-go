package config

import "testing"

func TestInitLogger(t *testing.T) {
	InitLogger()

	if Log == nil {
		t.Error("O Logger global não foi instanciado corretamente")
	}
}