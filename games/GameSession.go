package main

import (
	"math/rand"
	"sync"
	"time"
)

type NewsPeace struct {
  content string
  weight string
}

var NEWS = map[int64][]NewsPeace {
  // US OPEN (15:30)
  390: {
    {"US STOCKS OPEN HIGHER AHEAD OF INFLATION DATA", "GOOD"},
    {"US STOCKS OPEN LOWER AFTER JOBLESS CLAIMS", "BAD"},
    {"WALL STREET MIXED AFTER FED MINUTES", "NEUTRAL"},
    {"TREASURIES QUIET AHEAD OF POWELL SPEACH", "NEUTRAL"},
  },
  
  // EUROPEAN RATES
  90: {
    {"EUROPEAN SHARES IN THE GREEN AS INTREST RATE FEARS EASE", "GOOD"},
    {"LIVE: FTSE RISES AS U.S. PRIVATE SECTOR HIRING SLOWS", "GOOD"},
    {"EURO FALLS TO 1.0793 AHEAD OF ECB RATE CALL", "BAD"},
    {"BANK SHARES TUMBLE ON ITALY WINDFALL TAX", "BAD"},
    {"ITALY BOND YIELDS SOAR ON GRIM OUTLOOK", "BAD"},
  },

  210: {
    {"POLAND: DONALD TUSK TO BE NEW P.M. AFTER FAR-RIGHT DEFEAT", "NEUTRAL"},
    {"EU: NADIA CALVINO ELECTED NEW E.I.B. CHIEF", "NEUTRAL"},
    {"EU: INSTITUTIONS STRUGGLE TO FIND DEAL ON A.I. BILL", "BAD"},
    {"ITALY: GOVERNMENT REPEALS CONTROVERSIAL PENSION BILL", "GOOD"},
  },
}

type GameSession struct {
  users map[string]*UserSession
  balances map[string]int64
  shares map[string]int64
  time int64
  price int64

  timer *time.Ticker
  random *rand.Rand
  lock sync.Mutex
}

func (this *GameSession) Start() {
  this.balances = make(map[string]int64)
  this.shares = make(map[string]int64)

  for username, _ := range this.users {
    this.balances[username] = 10000
    this.shares[username] = 1
  }

  this.price = 1000
  this.time = 0

  for username, user := range this.users {
    user.sender.SendStart(this.shares[username], this.balances[username])
  }

  this.timer = time.NewTicker(time.Millisecond * 1000)
  this.random = rand.New(rand.NewSource(time.Now().UnixNano()));

  go func() {
    for range this.timer.C {
      hours := 9 + this.time / 60
      minutes := this.time % 60

      if this.random.Intn(100) <= 20 {
        this.lock.Lock()
        this.price += int64(-10 + this.random.Intn(20))
        this.lock.Unlock()
      }

      for _, user := range this.users {
        user.sender.SendTime(intFormatPad(hours, 2) + ":" + intFormatPad(minutes, 2), this.price)
      }

      if news, found := NEWS[this.time]; found {
        newsIndex := this.random.Intn(len(news))
        newsPiece := news[newsIndex]

        for _, user := range this.users {
          user.sender.SendNews(newsPiece.weight, newsPiece.content)
        }
      }

      if this.time == 480 {
        this.timer.Stop()
        this.end()
        return
      }

      this.time++
    }
  }()
}

func (this *GameSession) ForceEnd() {
  this.timer.Stop()
  this.end()
}

func (this *GameSession) Sell(username string, units int64) {
  this.lock.Lock()
  this.balances[username] += this.price * units
  this.shares[username] -= units

  this.price += int64(-10 + this.random.Intn(9))
  this.lock.Unlock()

  this.users[username].sender.SendConfirm(this.balances[username], this.shares[username])
}

func (this *GameSession) Buy(username string, units int64) {
  this.lock.Lock()
  this.balances[username] -= this.price * units
  this.shares[username] += units

  this.price += int64(1 + this.random.Intn(9))
  this.lock.Unlock()

  this.users[username].sender.SendConfirm(this.balances[username], this.shares[username])
}

func (this *GameSession) end() {
  winner := ""
  winnernw := int64(0)
  looser := ""
  loosernw := int64(0)
  someoneLeft := false

  random := rand.New(rand.NewSource(time.Now().UnixNano()))
  yield := float64(1 + random.Float64())

  for username, user := range this.users {
    dividend := float64(this.shares[username]) * yield
    netw := int64(dividend) + this.balances[username] + this.shares[username] * this.price

    if netw > winnernw && user.playing {
      winnernw = netw
      winner = username
    } else {
      loosernw = netw
      looser = username
    }

    if !user.playing {
      someoneLeft = true
    }

    user.playing = false
    user.game = nil
  }

  if winnernw != loosernw || someoneLeft {
    this.users[winner].sender.SendWon(winnernw, loosernw, yield)
    this.users[looser].sender.SendLost(loosernw, winnernw, yield)
    return
  }

  for _, user := range this.users {
    user.sender.SendEnd(winnernw, yield)
  }
}
