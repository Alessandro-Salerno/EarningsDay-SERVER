package main

import "strconv"

// Controlla se una funzione restituisce un errore
// se lo restituisce, chiama un gestore dell'errore (fail)
// se la funzione viene eseguita con successo, viene chiamato un gestore (success)
// il risultato di ritorno di fn() vene passato a success()
func runAndCheck[T any](fn func () (T, error), success func (T), fail func (error)) {
  res, err := fn();

  if err != nil {
    fail(err);
    return;
  }

  success(res);
}

// retituisce il primo elemento non nil in un array di errori
//  se non ci sono elementi non nil, viene restiuito un nil generico
func notNil(values ...error) error {
  for _, val := range values {
    if val != nil {
      return val;
    }
  }

  return nil;
}

// Converte un intero in uan stringa
// Assicurandosi che la stringa sia allineata correttamente con li zeri
func intFormatPad(value int64, padding int64) string {
  intval := strconv.FormatInt(value, 10)
  
  for len(intval) % int(padding) != 0 {
    intval = "0" + intval
  }

  return intval
}
