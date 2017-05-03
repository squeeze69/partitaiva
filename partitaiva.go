package partitaiva

import "strconv"

/*
Verifica partita IVA - 2017 - Squeeze69
Licenza: LGPL
Porting basato sulle informazioni pubblicate da Umberto Salsi su icosaedro:
sito web: http://www.icosaedro.it/cf-pi/index.html

package: https://github.com/squeeze69/partitaiva
con go: go get github.com/squeeze69/partitaiva
*/

//PIVAError : Partita IVA - error
type PIVAError struct {
	msg string
}

func (r *PIVAError) Error() string {
	return r.msg
}

//ItPartitaIva controlla se è una partita IVA valida nel formato
//se piva è vuota, viene considerata valida, per questo caso, il controllo dovrebbe essere altrove
func ItPartitaIva(piva string) (bool, *PIVAError) {
	if len(piva) == 0 {
		return true,nil
	}
	if len(piva) != 11 {
		er := new(PIVAError)
		er.msg = "Wrong Length"
		return false, er
	}

	var primo, secondo, i int
	for i = 0; i <= 9; i += 2 {
		v,err := strconv.Atoi(string(piva[i]))
		if err != nil {
			er := new(PIVAError)
			er.msg = "Invalid chars"
			return false, er
		}
		primo += v
	}
	for i = 1; i <= 9; i += 2 {
		v,err := strconv.Atoi(string(piva[i]))
		if err != nil {
			er := new(PIVAError)
			er.msg = "Invalid chars"
			return false, er
		}
		secondo = 2 * v
		if secondo > 9 {
			secondo = secondo - 9
		}
		primo += secondo
	}
	if v,_ := strconv.Atoi(string(piva[10])); (10-primo%10)%10 == v {
		return true, nil
	}

	er := new(PIVAError)
	er.msg = "Wrong Check Code"
	return false, er

}
