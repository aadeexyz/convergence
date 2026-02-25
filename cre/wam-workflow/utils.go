package main

import (
	"math"
	"math/big"
	"sort"
)

const halfLife = 6.0
const dt = 0.25
const sigmoidMidpoint = 5.0
const sigmoidScale = 2.0

func ComputeRawIndex(g, y, x float64) float64 {
	social := 0.50*y + 0.50*x
	return 10 * (0.30*g + 0.70*social)
}

func Sigmoid(x float64) float64 {
	return 1 / (1 + math.Exp(-(x-sigmoidMidpoint)/sigmoidScale))
}

func EMA(rawIndex float64, prevEMA *float64) float64 {
	if prevEMA == nil {
		return rawIndex
	}
	alpha := 1 - math.Pow(2, -(dt/halfLife))
	return *prevEMA + alpha*(rawIndex-*prevEMA)
}

// MedianNonZero computes the median of non-zero values from a []*big.Int slice.
// Zero values are ignored (uninitialized oracle slots). Returns nil if all values are zero.
func MedianNonZero(window []*big.Int) *big.Int {
	nonZero := make([]*big.Int, 0, len(window))
	for _, v := range window {
		if v != nil && v.Sign() > 0 {
			nonZero = append(nonZero, new(big.Int).Set(v))
		}
	}

	if len(nonZero) == 0 {
		return nil
	}

	sort.Slice(nonZero, func(i, j int) bool {
		return nonZero[i].Cmp(nonZero[j]) < 0
	})

	mid := len(nonZero) / 2
	if len(nonZero)%2 == 0 {
		sum := new(big.Int).Add(nonZero[mid-1], nonZero[mid])
		return sum.Div(sum, big.NewInt(2))
	}
	return nonZero[mid]
}
