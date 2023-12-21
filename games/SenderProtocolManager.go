package main

import (
	"strconv"

	"salerno.it/libcom"
)

type SenderProtocolManager struct {
  sender libcom.Sender
}

func (this SenderProtocolManager) SendTime(time string, price int64) {
  this.sender.Send("TIME*" + time + "*" + strconv.FormatInt(price, 10))
}

func (this SenderProtocolManager) SendLoginConfirm() {
  this.sender.Send("LOGINCONFIRM")
}

func (this SenderProtocolManager) SendUsers(users map[string]*UserSession) {
  userList := ""
  for username, _ := range users {
    userList += username + ","
  }
  this.sender.Send("USERS*" + userList)
}

func (this SenderProtocolManager) SendJoined(username string) {
  this.sender.Send("JOINED*" + username)
}

func (this SenderProtocolManager) SendLeft(username string) {
  this.sender.Send("LEFT*" + username)
}

func (this SenderProtocolManager) SendStart(units int64, balance int64) {
  this.sender.Send("START*" + strconv.FormatInt(units, 10) + "*" + strconv.FormatInt(balance, 10))
}

func (this SenderProtocolManager) SendRequest(username string) {
  this.sender.Send("REQUEST*" + username)
}

func (this SenderProtocolManager) SendRejected(username string) {
  this.sender.Send("REJECTED*" + username)
}

func (this SenderProtocolManager) SendChatMessage(message string) {
  this.sender.Send("CHAT*" + message)
}

func (this SenderProtocolManager) SendConfirm(balance int64, units int64) {
  this.sender.Send("CONFIRM*" + strconv.FormatInt(balance, 10) + "* " + strconv.FormatInt(units, 10))
}

func (this SenderProtocolManager) SendOffer(amount int64) {
  this.sender.Send("OFFER*" + strconv.FormatInt(amount, 10))
}

func (this SenderProtocolManager) SendNews(tone string, news string) {
  this.sender.Send("NEWS*" + tone + "*" + news)
}

func (this SenderProtocolManager) SendWon(yournw int64, othernw int64, yield float64) {
  this.sender.Send("WON*" + strconv.FormatInt(yournw, 10) + "* " + strconv.FormatInt(othernw, 10) + "*" + strconv.FormatFloat(yield, 'f', 0, 64))
}

func (this SenderProtocolManager) SendLost(yournw int64, othernw int64, yield float64) {
  this.sender.Send("LOST*" + strconv.FormatInt(yournw, 10) + "* " + strconv.FormatInt(othernw, 10) + "*" + strconv.FormatFloat(yield, 'f', 0, 64))
}

func (this SenderProtocolManager) SendEnd(networth int64, yield float64) {
  this.sender.Send("END*" + strconv.FormatInt(networth, 10) + "*" + strconv.FormatFloat(yield, 'f', 0, 64))
}

func (this SenderProtocolManager) SendCloseAck() {
  this.sender.Send("CLOSEACK")
}

func (this SenderProtocolManager) Error(message string) {
  this.sender.Send("ERR*" + message)
}
