package strings

import "testing"

const msgIndex = "%s (parte: %s) - Índices:  esperado (%d) <> encontrado (%d)."

func TestIndex(t *testing.T) {
	tests := []struct {
		texto    string
		parte    string
		esperado int
	}{
		{"Show", "é Show", 0},
		{"", "", 0},
		{"Eita", "eita", -1},
		{"leonardo", "o", 2},
	}

	for _, teste := range tests {
		t.Logf("Massa: %v", teste)
		atual := teste
	}
}
