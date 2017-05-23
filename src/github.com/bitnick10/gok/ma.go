package gok

type MA struct {
	N  int
	Ma []float64
}

func NewMA(chart *Chart, n int) *MA {
	ret := &MA{}
	ret.N = n
	ret.Ma = make([]float64, len(chart.Sticks()), len(chart.Sticks()))
	for data_i := 0; data_i < len(chart.Sticks()); data_i++ {
		begin_i := data_i - n + 1
		if begin_i < 0 {
			begin_i = 0
		}
		sum := 0.0
		sum_n := 0.0
		for i := begin_i; i <= data_i; i++ {
			sum += chart.Sticks()[i].Close
			sum_n += 1
		}
		ret.Ma[data_i] = sum / sum_n
	}
	return ret
}
