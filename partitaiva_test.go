package partitaiva

import (
	"fmt"
	"testing"
)

func TESTPartitaiva(t *testing.T) {
	res, err := ItPartitaIva("86334519757")
	if err != nil {
		t.Fatal("Error, should be valid", err)
	}
	if res {
		fmt.Println("Ok")
	}
	res, err = ItPartitaIva("86334519755")
	if err == nil {
		t.Fatal("Error, should be invalid")
	}
	if !res {
		fmt.Println("Ok", err)
	}
	res, err = ItPartitaIva("8633451975")
	if err == nil {
		t.Fatal("Error, should be invalid (length)")
	}
	if !res {
		fmt.Println("Ok", err)
	}
}
