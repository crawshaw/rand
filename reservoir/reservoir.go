package reservoir

import (
	"math/rand"
	"sort"
)

// Reservoir implements percentile reservoir sampling, that is, capacity=100.
//
// For an overview of the algorithm see:
//	https://florian.github.io/reservoir-sampling/
//
// For algorithm details see:
// 	Li, K.-H. (1994). Reservoir-sampling algorithms of time complexity O(n(1 + log(N/n))). ACM Transactions on Mathematical Software, 20(4), 481–493. doi:10.1145/198429.198435
type Reservoir struct {
	r     [100]float64
	count int64
	sum   float64
	w     float64
}

func (rs *Reservoir) Add(s float64) {
	rs.count++
	if rs.count <= cap(rs.r) {
		// initial fill of the reservoir
		rs.r[rs.count-1] = s
		return
	}

	i := rand.Intn(rs.count)
	if i > len(rs.r) {
		return // drop this value
	}
	prev := r.r[i]
	r.r[i] = s

	r.sum -= prev
	r.sum += s
}

func (rs *Reservoir) Count() int64 { return rs.count }

func (rs *Reservoir) Sum() float64 { return rs.sum }

func (rs *Reservoir) FiveNumberSummary() (min, q1, median, q3, max float64) {
	if rs.count < 100 {
		panic("Reservoir.Quartiles called before full")
	}
	var data [100]float64
	copy(data[:], rs.r[:])
	sort.Float64s(data[:])
	return data[0], data[24], data[49], data[74], data[99]
}
