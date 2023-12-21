package libcom

import (
	"net"
)


// Rappresenta il gestore
// delle comunicaizoni con
// il client
type Sender struct {
  Conn net.Conn;
}

// Invia un messaggio attraverso u ncanale TCP
// il messaggio è terminato da LF (L\n)
func (this Sender) Send(message string) {
  this.Conn.Write([]byte(message + "\n"));
}
