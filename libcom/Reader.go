package libcom

import (
	"errors"
	"io"
	"net"
)

// Gestisce il processo di lettura della stream
// di input che scontiene i messaggi inviati
// dal client
// La lettura avviene in modo asincrono
type Reader struct {
	Conn net.Conn;
  ForStop string;
  Stop bool;
  Consumer IMessageConsumer;
}

// Imposta il consumatore dei messaggi ed avvia la lettura
// Sarà invocato il metodo ConsumeMessage di un IMessageConsumer in caso
// di comando client.
// Nel caso di un messaggio di QUIT, verrà invocato il metodo OnClose
func (this Reader) SetConsumer(consumer IMessageConsumer) {
  this.Consumer = consumer;
  go this.runReader();
}

// Metodo privato
// Effettua la lettura vera e propria
func (this Reader) runReader() {
  var message string;
  tmp := make([]byte, 1); // buffer temporaneo per tenere un carattere

  // codice molto brutto
  // perdete ogni speranza o voi che entrate
  for (!this.Stop) {
    for {
      // Legge un carattere
      _, err := this.Conn.Read(tmp);

      // Controllo separatore
      // Controllo anche per EOF
      // (Fine del messaggio)
      if tmp[0] < 32 || errors.Is(err, io.EOF) {
        // controllo comando QUIT
        if (message == this.ForStop) {
          this.Stop = true;
          this.Consumer.OnClose()
          this.Conn.Close();
        }

        // Esecuzione del 
        if len(message) != 0 {
          this.Consumer.ConsumeMessage(message);
        }

        message = "";
        tmp[0] = 0;
        break;
      }

      // Errore di connessione
      if err != nil {
        this.Consumer.OnClose()
        this.Conn.Close();
        return;
      }

      // In tutti gli altri casi
      // il buffer viene concatenato
      message += string(tmp[0]);
    }
  }
}
