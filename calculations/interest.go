package calculations

import (
	"math"
)

// MonthlyInterest takes a starting value and computes interest over some number of months
//	startAmt: number of cents
//	numMonths: number of months
//	interest: interest % to use in returns calculation
func MonthlyInterest(startAmt int64, numMonths int32, interest float64) []int64 {
	var result []int64 = make([]int64, numMonths)

	result[0] = startAmt

	for i := 1; i < len(result); i++ {
		result[i] = result[i-1] + int64(math.Round(float64(result[i-1])*(interest)/12))
	}

	return result
}

// YearlyInterest takes a starting value and computes interest over some number of months
//	startAmt: number of cents
//	numMonths: number of months
//	interest: interest % to use in returns calculation
func YearlyInterest(startAmt int64, numYears int32, interest float64) []int64 {
	var result []int64 = make([]int64, numYears)

	result[0] = startAmt

	for i := 1; i < len(result); i++ {
		result[i] = result[i-1] + int64(math.Round(float64(result[i-1])*(interest)))
	}

	return result
}
