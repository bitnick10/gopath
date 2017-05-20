package quant

import (
	"github.com/bitnick10/gok"
	"time"
)

type PositionType int

const (
	Long  PositionType = 1
	Short PositionType = -1
)

type Position struct {
	cc      *gok.Chart
	account *Account
	Type    PositionType
	// CreateI     int
	// CloseI      int
	CreateTime  time.Time
	CloseTime   time.Time
	CreatePrice float64
	ClosePrice  float64
}

func (self *Position) ProfitPercent() float64 {
	profit := float64(self.Type) * (self.ClosePrice - self.CreatePrice) / self.CreatePrice
	return profit
}

func (self *Position) GetProfitPercent(price float64) float64 {
	profit := float64(self.Type) * (price - self.CreatePrice) / self.CreatePrice
	return profit
}

func (self *Position) TypeInShortString() string {
	if self.Type == Long {
		return "L"
	} else {
		return "S"
	}
}
