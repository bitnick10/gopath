package gok

type KDJ struct {
	N   int
	Sn1 int
	Sn2 int
	K   []float64
	D   []float64
	J   []float64
}

func (self *KDJ) IsKGTD(dataI int) bool {
	return self.K[dataI] > self.D[dataI]
}

func (self *KDJ) IsKUp(dataI int) bool {
	return self.K[dataI] > self.K[dataI-1]
}

func (self *KDJ) IsJUp(dataI int) bool {
	return self.J[dataI] > self.J[dataI-1]
}

func NewKDJ(chart *Chart, n, sn1, sn2 int) *KDJ {
	ret := &KDJ{N: n, Sn1: sn1, Sn2: sn2}
	rsvs := chart.RawStochasticValues(n)
	ret.K = SMA(rsvs, sn1, 1.0)
	ret.D = SMA(ret.K, sn2, 1.0)
	ret.J = make([]float64, len(ret.D))
	for i := 0; i < len(ret.D); i++ {
		ret.J[i] = ret.K[i]*3 - ret.D[i]*2
	}
	return ret
}
