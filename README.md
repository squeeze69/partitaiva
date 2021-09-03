# Verifica partita IVA in [GO](http://golang.org)

[![Build Status](https://travis-ci.com/squeeze69/partitaiva.svg?branch=master)](https://travis-ci.com/squeeze69/partitaiva)

## Licenza: LGPLv3

**Package**: go get github.com/squeeze69/partitaiva

**No Davvero Nota Bene**: Il codice fiscale temporaneo ha le stesse caratteristiche della Partita IVA, questo modulo può essere usato per tali verifiche.

Porting basato sulle informazioni pubblicate da Umberto Salsi su [Icosaedro](http://www.icosaedro.it/cf-pi/index.html)

Uso:

```go
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

Se cercaste il [codice fiscale](https://github.com/squeeze69/codicefiscale)
