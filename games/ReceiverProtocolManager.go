package main

import (
	"strings"
)

type ReceiverProtocolManager struct {
	commandConsumer ICommandConsumer
}

// Esegue il parsing del messaggio
  // e lancia l'esecuzione del gestore
func (this ReceiverProtocolManager) ConsumeMessage(message string) {
  chunks := strings.Split(message, "*")

  // Se il comando ha una lunghezza insufficientePPLI
  // if len(chunks) < 1 {
  //   this.commandConsumer.Error("Invalid command syntax. Need at least one argument.")
  //   return
  // }

  // Estrapolazione del comando
  command := chunks[0]

  switch (command) {
  case "GETUSERS": 
    this.expectAndCall(0, func (args ...string) {
      this.commandConsumer.GetUsers()
    }, chunks)

  case "LOGIN": 
    this.expectAndCall(1, func (args ...string) {
      this.commandConsumer.Login(args[0])
    }, chunks)

  case "QUITGAME":
    this.expectAndCall(0, func (args ...string) {
      this.commandConsumer.QuitGame()
    }, chunks)

  case "INVITE":
    this.expectAndCall(1, func (args ...string) {
      this.commandConsumer.Invite(args[0])
    }, chunks)

  case "ACCEPTGAME":
    this.expectAndCall(0, func (args ...string) {
      this.commandConsumer.AcceptGame()
    }, chunks)

  case "REJECTGAME":
    this.expectAndCall(0, func (args ...string) {
      this.commandConsumer.RejectGame()
    }, chunks)

  case "CHAT":
    this.expectAndCall(1, func (args ...string) {
      this.commandConsumer.Chat(args[0])
    }, chunks)

  // MISSING
  // BUY
  // SELL

  case "SELL":
    this.expectAndCall(1, func (args ...string) {
      this.commandConsumer.Sell(args[0])
    }, chunks)

  case "BUY":
    this.expectAndCall(1, func (args ...string) {
      this.commandConsumer.Buy(args[0])
    }, chunks)


  default:
    this.commandConsumer.Error("Unknown command")
  }
}

// Gestisce la chiusura della connessione
  // L'implementazione si deve occupare di eventuali
  // chiusure di stream ed altre connessioni
func (this ReceiverProtocolManager) OnClose() {
  this.commandConsumer.Logout()
}

// Compara il numero di parametri forniti con quelli richiesti
// se i due sono equivalenti, chiama la funzione callback omittendo il primo parametro (nome comando)
// altrimenti manda un errore via il commandConsumer
func (this ReceiverProtocolManager) expectAndCall(argc int, callback func (...string), args []string) {
  if len(args) - 1 == argc {
    callback(args[1:]...)
    return
  }

  this.commandConsumer.Error("Incorect number of arguments")
}
