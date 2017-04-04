package partitaiva

import (
	"fmt"
	"testing"
)

func TESTPartitaiva(t *testing.T) {
	res, err := ItPartitaIva("00468990015")
	if err != nil {
		t.Fatal("Error, should be valid")
	}
	if res {
		fmt.Println("Ok")
	}
	res, err = ItPartitaIva("00468990014")
	if err == nil {
		t.Fatal("Error, should be invalid")
	}
	if !res {
		fmt.Println("Ok", err)
	}
	res, err = ItPartitaIva("004689900")
	if err == nil {
		t.Fatal("Error, should be invalid (length)")
	}
	if !res {
		fmt.Println("Ok", err)
	}
}
