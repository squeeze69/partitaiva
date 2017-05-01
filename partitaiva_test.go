package partitaiva

import (
	"fmt"
	"testing"
)

func TestItPartitaiva(t *testing.T) {
	pivaOk := []string{"12345678903", "00000000000", "44444444440"}
	pivaKO := []string{"xxxxxxxxxxx", "123456789012", "44444444444", "12345678901", "44444+44440", "corto"}
	for _, v := range pivaOk {
		res, err := ItPartitaIva(v)
		if err != nil {
			t.Fatal("Error, ", v, " should be valid", err)
		}
		if res {
			fmt.Println("Ok (valid)", v)
		}
	}

	for _, v := range pivaKO {
		res, err := ItPartitaIva(v)
		if err == nil {
			t.Fatal("Error, ", v, " should be invalid")
		}
		if !res {
			fmt.Println("Ok (invalid)", v, err)
		}
	}
}
