package partitaiva

import "strconv"

/*
Verifica partita IVA
Versione: 1.0
Data: 1/5/2017
Autore: Squeeze69
Licenza: LGPL
Porting basato sulle informazioni pubblicate da Umberto Salsi su Icosaedro:
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
//Nota: se piva è vuota, viene considerata valida, per questo caso, il controllo dovrebbe essere altrove
//Ingresso: piva: stringa
//Uscita: bool:true (a posto)/false (problemi) e *PIVAError (nil (a posto)/puntatore all'errore (problemi)
func ItPartitaIva(piva string) (bool, *PIVAError) {
	if len(piva) == 0 {
		return true,nil
	}
	if len(piva) != 11 {
		er := new(PIVAError)
		er.msg = "Lunghezza Sbagliata"
		return false, er
	}

	var primo, secondo, i int
	for i = 0; i <= 9; i += 2 {
		v,err := strconv.Atoi(string(piva[i]))
		if err != nil {
			er := new(PIVAError)
			er.msg = "Caratteri Non Validi"
			return false, er
		}
		primo += v
	}
	for i = 1; i <= 9; i += 2 {
		v,err := strconv.Atoi(string(piva[i]))
		if err != nil {
			er := new(PIVAError)
			er.msg = "Caratteri Non Validi"
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
	er.msg = "Carattere Di Controllo Non Valido"
	return false, er

}
