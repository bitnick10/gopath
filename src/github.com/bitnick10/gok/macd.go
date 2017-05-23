package gok

type MACD struct {
	ShortN int
	LongN  int
	DeaN   int
	Diff   []float64
	Dea    []float64
	Bar    []float64
}

func (self *MACD) IsDiffGT0(dataI int) bool {
	return self.Diff[dataI] > 0
}

func NewMACD(chart *Chart, shortN, longN, deaN int) *MACD {
	ret := &MACD{ShortN: shortN, LongN: longN, DeaN: deaN}
	closes := make([]float64, len(chart.Sticks()))
	for i := 0; i < len(closes); i++ {
		closes[i] = chart.Sticks()[i].Close
	}

	shortEMAs := EMA(closes, shortN)
	longEMAs := EMA(closes, longN)
	ret.Diff = make([]float64, len(chart.Sticks()))
	ret.Bar = make([]float64, len(chart.Sticks()))
	for i := 0; i < len(ret.Diff); i++ {
		ret.Diff[i] = shortEMAs[i] - longEMAs[i]
	}
	ret.Dea = EMA(ret.Diff, deaN)
	for i := 0; i < len(ret.Diff); i++ {
		ret.Bar[i] = 2 * (ret.Diff[i] - ret.Dea[i])
	}
	return ret
}
