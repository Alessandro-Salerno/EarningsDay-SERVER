package main;


import (
	"fmt"
  "os"
  "net"

  "salerno.it/libcom"
)


// the main function
func main() {
	l, err := net.Listen("tcp", ":60000")

  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }

  defer l.Close()
  
  users := make(map[string]*UserSession);

  for {
    c, err := l.Accept()

    if err != nil {
      fmt.Println(err)
      continue
    }


    sender := libcom.Sender{Conn: c};
    reader := libcom.Reader{Conn: c, ForStop: "LOGOUT", Stop: false, Consumer: nil};

    senderProtocol := SenderProtocolManager{sender};
    userSession := new(UserSession)
    userSession.sender = senderProtocol
    application := Application{users: users, session: userSession}
    receiverProtocol := ReceiverProtocolManager{application}

    reader.SetConsumer(receiverProtocol);
  }
}
