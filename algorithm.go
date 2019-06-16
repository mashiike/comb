package comb

import (
	crand "crypto/rand"
	"math"
	"math/big"
	"math/rand"
	"sort"

	"github.com/seehuhn/mt19937"
)

type Config struct {
	maxIter int
	rng     *rand.Rand
}

type Option func(*Config)

type simulatedAnnealing struct {
	length     int
	groupCount int
	energy     func([][]int) float64
	config     *Config
}

func newSimulatedAnneling(length, groupCount int, energy func([][]int) float64, options ...Option) *simulatedAnnealing {
	seed, _ := crand.Int(crand.Reader, big.NewInt(math.MaxInt64))
	rng := rand.New(mt19937.New())
	rng.Seed(seed.Int64())
	config := &Config{
		maxIter: 10000,
		rng:     rng,
	}
	for _, option := range options {
		option(config)
	}

	return &simulatedAnnealing{
		length:     length,
		groupCount: groupCount,
		energy:     energy,
		config:     config,
	}
}

func (alg *simulatedAnnealing) solve() [][]int {
	current := alg.initState()
	minimum := deepcopy(current)
	minimumEnergy := alg.energy(minimum)

	for k := 0; k < alg.config.maxIter; k++ {
		i1, j1, i2, j2 := genSwapIndex(current, alg.config.rng)
		doswap(current, i1, j1, i2, j2)
		currentEnergy := alg.energy(current)
		if currentEnergy <= minimumEnergy {
			minimum = deepcopy(current)
			minimumEnergy = currentEnergy
			continue
		}
		if alg.config.rng.Float64() > math.Exp((minimumEnergy-currentEnergy)/(float64(alg.config.maxIter-k)+0.01)) {
			doswap(current, i2, j2, i1, j1)
		}
	}
	return groupSort(minimum)
}

func (alg *simulatedAnnealing) initState() [][]int {
	ret := make([][]int, 0, alg.groupCount)
	for i := 0; i < alg.groupCount; i++ {
		ret = append(ret, make([]int, 0, alg.length/alg.groupCount+1))
	}

	for k := 0; k < alg.length; {
		for i := 0; i < alg.groupCount; i++ {
			ret[i] = append(ret[i], k)
			k++
			if k >= alg.length {
				break
			}
		}
	}
	return ret
}

func genSwapIndex(current [][]int, rng *rand.Rand) (i1, j1, i2, j2 int) {
	switch len(current) {
	case 2:
		i1 = 0
		i2 = 1
	default:
		i1 = rng.Intn(len(current))
		i2 = i1
		for i1 == i2 {
			i2 = rng.Intn(len(current))
		}
	}
	j1 = rng.Intn(len(current[i1]))
	j2 = rng.Intn(len(current[i2]))
	return
}

func doswap(current [][]int, i1, j1, i2, j2 int) {
	tmp := current[i1][j1]
	current[i1][j1] = current[i2][j2]
	current[i2][j2] = tmp
}

func deepcopy(ind [][]int) [][]int {
	ret := make([][]int, len(ind))
	for i, inner := range ind {
		ret[i] = make([]int, len(inner))
		copy(ret[i], inner)
	}
	return ret
}

func groupSort(ind [][]int) [][]int {
	for _, slice := range ind {
		sort.Slice(slice, func(i, j int) bool { return slice[i] < slice[j] })
	}
	sort.Slice(ind, func(i, j int) bool { return ind[i][0] < ind[j][0] })
	return ind
}

func MaxIter(n int) Option {
	return func(config *Config) {
		config.maxIter = n
	}
}

func Seed(n int64) Option {
	return func(config *Config) {
		config.rng.Seed(n)
	}
}
