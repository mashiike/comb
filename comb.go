package comb

import "reflect"

func Slice(slice interface{}, groupCount int, score func(i [][]int) float64) [][]int {
	value := reflect.ValueOf(slice)
	length := value.Len()
	alg := &simulatedAnnealing{
		length:     length,
		groupCount: groupCount,
		score:      score,
	}
	return alg.solve()
}
