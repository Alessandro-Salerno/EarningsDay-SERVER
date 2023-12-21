package main

type ISenderProtocol interface {
  SendTime(time string, price int64)
  SendLoginConfirm()
  SendUsers(users map[string]*UserSession)
  SendJoined(username string)
  SendLeft(username string)
  SendStart(units int64, balance int64)
  SendRequest(username string)
  SendRejected(username string)
  SendChatMessage(message string)
  SendConfirm(balance int64, units int64)
  SendOffer(amount int64)
  SendNews(tone string, news string)
  SendWon(yournw int64, othernw int64, yield float64)
  SendLost(yournw int64, othernw int64, yield float64)
  SendEnd(networth int64, yield float64)
  SendCloseAck()
  Error(message string)
}
