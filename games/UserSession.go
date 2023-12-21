package main

type UserSession struct {
  sender ISenderProtocol
  username string
  playing bool
  vsUsername string
  game *GameSession
}
