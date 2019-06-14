package comb_test

import (
	"fmt"

	"github.com/mashiike/comb"
)

func init() {
	comb.SetRngSeed(1)
}

func ExampleSlice() {
	target := []float64{1.0, 2.0, 3.0, 3.5, 4.0, 3.8}
	score := func(ind []int) float64 {
		total := 0.0
		for _, i := range ind {
			total += target[i]
		}
		return total
	}
	energy := func(inds [][]int) float64 {
		var min, max float64
		for _, ind := range inds {
			s := score(ind)
			if min > s {
				min = s
			}
			if max < s {
				max = s
			}
		}
		return (max - min) * (max - min)
	}
	groups := comb.Slice(target, 4, energy)

	for _, group := range groups {
		fmt.Printf("[")
		for _, i := range group {
			fmt.Printf(" %f ", target[i])
		}
		fmt.Printf("] => %f\n", score(group))
	}
	//Output:
	//[ 1.000000  3.800000 ] => 4.800000
	//[ 2.000000  3.000000 ] => 5.000000
	//[ 3.500000 ] => 3.500000
	//[ 4.000000 ] => 4.000000
}
