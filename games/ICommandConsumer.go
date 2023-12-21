package main


// Definisce l'interfaccia per gestire
// i singoli comandi
// È dichiarato un metodo per ogni comando client
type ICommandConsumer interface {
  GetUsers();
  Login(username string)
  Logout()
  QuitGame()
  Invite(username string)
  AcceptGame()
  RejectGame()
  Chat(message string)
  Sell(units string)
  Buy(units string)
  // AcceptOffer()
  // RejectOffer()
  Error(error string)
}
