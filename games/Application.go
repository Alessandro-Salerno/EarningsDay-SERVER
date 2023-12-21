package main

import (
	_ "fmt"
	"strconv"
	"strings"
)


type Application struct {
  users map[string]*UserSession
  session *UserSession
}

func (this Application) GetUsers() {
  this.session.sender.SendUsers(this.users)
}

func (this Application) Login(username string) {
  _, found := this.users[username]

  if found {
    this.session.sender.Error("Username already in use")
    return
  }
  
  if strings.Contains(username, "*") {
    this.session.sender.Error("invalid username")
    return
  }

  this.broadcastJoined(username)
  this.session.username = username
  this.users[username] = this.session
  this.session.sender.SendLoginConfirm()
}

func (this Application) Logout() {
  if this.session.playing {
    this.session.playing = false
    this.session.game.ForceEnd()
  }

  _, found := this.users[this.session.username]
  if found {
    delete(this.users, this.session.username)
  }

  this.session.sender.SendCloseAck()
  this.broadcastLeft(this.session.username)
}

func (this Application) QuitGame() {
  if !this.session.playing {
    this.session.sender.Error("You're not playing")
    return
  }

  this.session.playing = false
  this.session.game.ForceEnd()
  this.session.vsUsername = ""
}

func (this Application) Invite(username string) {
  user, found := this.users[username]
  if !found {
    this.session.sender.Error("No such user")
    return
  }

  if user.playing {
    this.session.sender.Error("User is already playing")
    return
  }

  this.session.vsUsername = username
  user.vsUsername = this.session.username
  user.sender.SendRequest(this.session.username)
}

func (this Application) AcceptGame() {
  if this.session.vsUsername == "" {
    this.session.sender.Error("Noone asked you")
    return
  }

  user, found := this.users[this.session.vsUsername]
  if !found {
    this.session.sender.Error("Could not find user")
    return
  }

  user.playing = true
  user.vsUsername = this.session.username
  this.session.playing = true

  gameusers := make(map[string]*UserSession)
  gameusers[this.session.username] = this.session
  gameusers[user.username] = user

  gameSession := new(GameSession)
  gameSession.users = gameusers

  user.game = gameSession
  this.session.game = gameSession
  this.session.game.Start()
}

func (this Application) RejectGame() {
  if this.session.vsUsername == "" {
    this.session.sender.Error("Noone asked you")
    return
  }

  user, found := this.users[this.session.vsUsername]
  if !found {
    this.session.sender.Error("Could not find user")
    return
  }

  user.playing = false
  user.vsUsername = ""
  user.sender.SendRejected(this.session.username)
}

func (this Application) Chat(message string) {
  if !this.session.playing {
    this.session.sender.Error("You're not playing")
    return
  }

  user, found := this.users[this.session.vsUsername]
  if !found {
    // end game
    return
  }

  user.sender.SendChatMessage(message)
  this.session.sender.SendChatMessage("-- YOU SENT A MESSAGE --")
}

func (this Application) Sell(units string) {
  if !this.session.playing {
    this.Error("You're not playing")
    return
  }

  val, err := strconv.Atoi(units)
  if err != nil {
    this.Error("Invalid number format")
    return
  }

  this.session.game.Sell(this.session.username, int64(val))
}

func (this Application) Buy(units string) {
  if !this.session.playing {
    this.Error("You're not playing")
    return
  }

  val, err := strconv.Atoi(units)
  if err != nil {
    this.Error("Invalid number format")
    return
  }

  this.session.game.Buy(this.session.username, int64(val))

}

func (this Application) Error(error string) {
  this.session.sender.Error(error)
}

func (this Application) broadcastLeft(username string) {
  for _, sender := range this.users {
    sender.sender.SendLeft(username)
  }
}

func (this Application) broadcastJoined(username string) {
  for _, sender := range this.users {
    sender.sender.SendJoined(username)
  }
}
