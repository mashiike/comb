# comb

[![CircleCI](https://circleci.com/gh/mashiike/rating.svg?style=svg)](https://circleci.com/gh/mashiike/rating)

This library provides the ability to split Slice elements into multiple groups.  
The approach is to solve combinatorial optimization problems by minimizing the energy function.  

## Usage
The interface is similar to sort.Slice. It can be used as the following example.
```go

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
```

