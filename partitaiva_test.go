package partitaiva

import (
	"fmt"
	"testing"
)

func TestItPartitaiva(t *testing.T) {
	pivaOk := []string{"12345678903", "00000000000", "44444444440",""}
	pivaKO := []string{"xxxxxxxxxxx", "123456789012", "44444444444", "12345678901", "44444+44440", "corto", "123456789.3",
	"4444X444440"}
	for _, v := range pivaOk {
		res, err := ItPartitaIva(v)
		if err != nil {
			t.Fatal("Ko. Error, ", v, " should be valid", err)
		}
		if res {
			fmt.Printf("Ok (valid) \"%s\"\n", v)
		}
	}

	for _, v := range pivaKO {
		res, err := ItPartitaIva(v)
		if err == nil {
			t.Fatal("Ko. Error, ", v, " should be invalid")
		}
		if !res {
			fmt.Printf("Ok (invalid) \"%s\",%s\n", v, err)
		}
	}
}
