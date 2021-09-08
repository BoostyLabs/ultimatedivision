// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package templatemath

// Iter function for creating a slice of iterable values from minimum and maximum values.
func Iter(start, end int) []int {
	var slice []int
	for i := start; i <= end; i++ {
		slice = append(slice, i)
	}
	return slice
}

// Inc function - unary operator that add one, to their operand.
func Inc(i int) int {
	return i + 1
}

// Dec function - unary operator that subtract one, from their operand.
func Dec(i int) int {
	return i - 1
}
