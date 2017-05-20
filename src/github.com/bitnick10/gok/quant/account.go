package quant

import (
	"fmt"
	"time"
)

type Account struct {
	ClosedPositions []Position
	Positions       []Position
}

func NewAccount() *Account {
	ret := &Account{}
	ret.Positions = make([]Position, 0)
	ret.ClosedPositions = make([]Position, 0)
	return ret
}

func (self *Account) OpenPosition(p Position) {
	// position := Position{CreateI: i, CreatePirce: createPrice}
	self.Positions = append(self.Positions, p)
}

func (self *Account) Print() {
	for _, p := range self.ClosedPositions {
		fmt.Printf("%s %s %.0f %5.2f C at %s\n",
			p.CreateTime.Format("2006-01-02 15:04"), p.TypeInShortString(), p.CreatePrice, p.ProfitPercent()*100, p.CloseTime.Format("2006-01-02 15:04"))
		// fmt.Println(p.CreateTime.Format("2006-01-02 15:04") )
	}
	for _, p := range self.Positions {
		fmt.Println(p.CreateTime.Format("2006-01-02 15:04"))
	}
	// position := Position{CreateI: i, CreatePirce: createPrice}
	// self.Positions = append(self.Positions, position)
}

func (self *Account) ClosePosition(i int, closePrice float64, closeTime time.Time) {
	self.Positions[i].ClosePrice = closePrice
	self.Positions[i].CloseTime = closeTime
	self.ClosedPositions = append(self.ClosedPositions, self.Positions[i])
	self.Positions = append(self.Positions[:i], self.Positions[i+1:]...)
}

func (self *Account) WinRate() float64 {
	n := 0.0
	for _, p := range self.ClosedPositions {
		if p.ProfitPercent() > 0 {
			n += 1
		}
	}
	return n / float64(len(self.ClosedPositions))
}

func (self *Account) ProfitPercent() float64 {
	profit := 0.0
	for _, p := range self.ClosedPositions {
		profit += p.ProfitPercent()
	}
	return profit
}

func (self *Account) AverageProfit() float64 {
	return self.ProfitPercent() / float64(len(self.ClosedPositions))
}

// func (self *Account) CloseAllPositions() {
// 	for _, p := range self.Positions {
// 		self.ClosedPositions = append(self.ClosedPositions, p)
// 	}
// 	self.Positions = self.Positions[:0]
// }
