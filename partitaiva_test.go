package partitaiva

import (
	"fmt"
	"testing"
)

func TestItPartitaiva(t *testing.T) {
	pivaOk := []string{"12345678903","00000000000", "44444444440"}
	for _, v := range pivaOk {
		res, err := ItPartitaIva(v)
		if err != nil {
			t.Fatal("Error, ", v, " should be valid", err)
		}
		if res {
			fmt.Println(v," Ok")
		}
	}

	res, err := ItPartitaIva("123456789012")
	if err == nil {
		t.Fatal("Error, should be invalid")
	}
	if !res {
		fmt.Println("Ok", err)
	}
	res, err = ItPartitaIva("12345678901")
	if err == nil {
		t.Fatal("Error, should be invalid (length)")
	}
	if !res {
		fmt.Println("Ok", err)
	}
}
