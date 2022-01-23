package strings

import (
	"strings"
	"testing"
)

const msgIndex = "%s (parte: %s) - Índices:  esperado (%d) <> encontrado (%d)."

func TestIndex(t *testing.T) {
	tests := []struct {
		texto    string
		parte    string
		esperado int
	}{
		{"É Show", "Show", 3},
		{"", "", 0},
		{"Eita", "eita", -1},
		{"leonardo", "o", 2},
	}

	for _, teste := range tests {
		t.Logf("Massa: %v", teste)
		atual := strings.Index(teste.texto, teste.parte)

		if atual != teste.esperado {
			t.Errorf(msgIndex, teste.texto, teste.parte, teste.esperado, atual)
		}
	}
}
