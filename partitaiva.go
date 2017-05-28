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

//errPIVA : genera un *PIVAError, non esportata
func errPIVA(s string) *PIVAError {
	err := new(PIVAError)
	err.msg = s
	return err
}

//ItPartitaIva controlla se è una partita IVA valida nel formato
//Nota: se piva è vuota, viene considerata valida, per questo caso, il controllo dovrebbe essere altrove
//Ingresso: piva: stringa
//Uscita: bool:true (a posto)/false (problemi) e *PIVAError (nil (a posto)/puntatore all'errore (problemi)
func ItPartitaIva(piva string) (bool, *PIVAError) {
	if len(piva) == 0 {	//convenzione da Icosaedro
		return true, nil
	}
	if len(piva) != 11 {
		return false, errPIVA("Lunghezza Sbagliata")
	}

	var primo, secondo, i int
	for i = 0; i <= 9; i += 2 {
		v, err := strconv.Atoi(string(piva[i]))
		if err != nil {
			return false, errPIVA("Caratteri Non Validi")
		}
		primo += v
		v, err = strconv.Atoi(string(piva[i+1]))
		if err != nil {
			return false, errPIVA("Caratteri Non Validi")
		}
		secondo = 2 * v
		if secondo > 9 {
			secondo = secondo - 9
		}
		primo += secondo
	}
	//corrisponde?
	if v, _ := strconv.Atoi(string(piva[10])); (10-primo%10)%10 == v {
		return true, nil
	}
	//non corrisponde
	return false, errPIVA("Carattere Di Controllo Non Valido")
}
