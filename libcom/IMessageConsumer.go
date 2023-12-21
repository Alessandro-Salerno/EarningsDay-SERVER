package libcom;


// Le classi che implementano quest'interfaccia
// descrivono il comportamento del server
// all'arrivo di un messaggio
type IMessageConsumer interface {
  ConsumeMessage(message string);
  OnClose();
}
