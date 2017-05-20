package gok

func SMA(data []float64, N int, M float64) []float64 {
	ret := make([]float64, len(data))
	ret[0] = data[0]
	k := M / float64(N)
	for i := 1; i < len(data); i++ {
		ret[i] = data[i]*k + ret[i-1]*(1-k)
	}
	return ret
}

func EMA(data []float64, N int) []float64 {
	return SMA(data, N, 2.0)
}
