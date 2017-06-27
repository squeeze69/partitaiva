package partitaiva

import (
	"fmt"
	"testing"
)

func TestItPartitaiva(t *testing.T) {
	//partite IVA valide (formalmente)
	pivaOk := []string{"12345678903", "00000000000", "44444444440", ""}
	//partite IVA non valide
	pivaKO := []string{"xxxxxxxxxxx", "123456789012", "44444444444", "12345678901", "44444+44440", "corto", "123456789.3",
		"4444X444440", "1234567890x", }
	for _, v := range pivaOk {
		res, err := ItPartitaIva(v)
		if err != nil {
			t.Fatal("Ko. Errore, ", v, " dovrebbe essere valido", err)
		}
		if res {
			fmt.Printf("Ok (valido) \"%s\"\n", v)
		}
	}

	for _, v := range pivaKO {
		res, err := ItPartitaIva(v)
		if err == nil {
			t.Fatal("Ko. Errore, ", v, " NON dovrebbe essere valido")
		}
		if !res {
			fmt.Printf("Ok (non valido) \"%s\",%s\n", v, err)
		}
	}
}
