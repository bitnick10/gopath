package gok

import (
	"time"
)

type Candlestick struct {
	Prev                   *Candlestick
	OpenTime               time.Time
	CloseTime              time.Time
	Open, Close, High, Low float64
	Volume                 float64
}

func (self *Candlestick) IsStickUp() bool {
	return self.Close > self.Open
}

func (self *Candlestick) IsPriceIn(price float64) bool {
	return (self.Low <= price && price <= self.High)
}

func (self *Candlestick) UpPercent() float64 {
	return (self.Close - self.Prev.Close) / self.Prev.Close
}

func (self *Candlestick) Amplitude() float64 {
	return (self.High - self.Low) / self.Open
}

func CombineCandlesticks(sticks ...Candlestick) Candlestick {
	ret := sticks[0]
	l := len(sticks)
	ret.Close = sticks[l-1].Close
	ret.CloseTime = sticks[l-1].CloseTime

	for i := 1; i < len(sticks); i++ {
		if ret.High < sticks[i].High {
			ret.High = sticks[i].High
		}
		if ret.Low > sticks[i].Low {
			ret.Low = sticks[i].Low
		}
		ret.Volume += sticks[i].Volume
	}
	return ret
}
