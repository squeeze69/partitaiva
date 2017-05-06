# Verifica partita IVA in [GO](http://golang.org)

## License: LGPLv3

#### 2017

package: go get github.com/squeeze69/partitaiva

Porting basato sulle informazioni pubblicate da Umberto Salsi su [Icosaedro](http://www.icosaedro.it/cf-pi/index.html)

Uso:

```
package main

import (
	"github.com/squeeze69/partitaiva"
	"fmt"
)

func main() {
	ok, err := partitaiva.ItPartitaIva("12334566")
	if err != nil {
		fmt.Println("Partita IVA non valida:",err)
	} else {
		fmt.Println("Partita IVA valida")
	}
}
```

simile ma per il [codice fiscale](https://github.com/squeeze69/codicefiscale)
