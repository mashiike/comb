package comb

import "reflect"

func Slice(slice interface{}, groupCount int, energy func(i [][]int) float64, options ...Option) [][]int {
	value := reflect.ValueOf(slice)
	length := value.Len()
	alg := newSimulatedAnneling(length, groupCount, energy, options...)
	return alg.solve()
}
