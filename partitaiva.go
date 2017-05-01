package partitaiva

import "unicode"

//Verifica partita IVA - 2017 - Squeeze69
//Porting (più o meno) dalla versione in PHP presente sul sito Icosaedro:
//un grazie ad Umberto Salsi
//http://www.icosaedro.it/cf-pi/

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
	}
	if len(piva) != 11 {
		er := new(PIVAError)
		er.msg = "Wrong Length"
		return false, er
	}

	//verifica se è composto da soli digit
	for _, c := range piva {
		if !unicode.IsDigit(c) {
			er := new(PIVAError)
			er.msg = "Invalid chars"
			return false, er
		}
	}
	var primo, secondo, i int
	for i = 0; i <= 9; i += 2 {
		primo += ordv[string(piva[i])]
	}
	for i = 1; i <= 9; i += 2 {
		secondo = 2 * ordv[string(piva[i])]
		if secondo > 9 {
			secondo = secondo - 9
		}
		primo += secondo
	}
	if (10-primo%10)%10 == ordv[string(piva[10])] {
		return true, nil
	}

	er := new(PIVAError)
	er.msg = "Check error"
	return false, er

}
