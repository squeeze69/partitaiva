package partitaiva
//Verifica partita IVA - 2017 - Squeeze69

//PIVAError : Partita IVA - error
type PIVAError struct {
	msg string
}

func (r *PIVAError) Error() string {
	return r.msg
}

//ItPartitaIva check if it"s a valid partita IVA (VAT)
func ItPartitaIva(piva string) (bool, *PIVAError) {
	ordv := map[string]int{
		"0": 0, "1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9,
		"A": 0, "B": 1, "C": 2, "D": 3, "E": 4, "F": 5, "G": 6, "H": 7,
		"I": 8, "J": 9, "K": 10, "L": 11, "M": 12, "N": 13, "O": 14, "P": 15, "Q": 16, "R": 17,
		"S": 18, "T": 19, "U": 20, "V": 21, "W": 22, "X": 23, "Y": 24, "Z": 25,
	}
	if len(piva) != 11 {
		er := new(PIVAError)
		er.msg = "Wrong Length"
		return false, er
	}
	var first, secondo, i int
	for i = 0; i <= 9; i += 2 {
		first += ordv[string(piva[i])] - ordv["0"]
	}
	for i = 1; i <= 9; i += 2 {
		secondo = 2 * (ordv[string(piva[i])] - ordv["0"])
		if secondo > 9 {
			secondo = secondo - 9
		}
		first += secondo
	}
	if 10-first%10 != ordv[string(piva[10])]-ordv["0"] {
		er := new(PIVAError)
		er.msg = "Check error"
		return false, er
	}
	return true, nil
}
