// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package search

import "sort"

// BinarySearch searches for elements in the slice,
// it return true if element present in given data, and false otherwise.
func BinarySearch(allData []string, element string) bool {
	sort.Strings(allData)

	startIndex := 0
	endIndex := len(allData) - 1

	for startIndex <= endIndex {
		median := (startIndex + endIndex) / 2

		switch {
		case allData[median] == element:
			return true
		case allData[median] < element:
			startIndex = median + 1
		case allData[median] > element:
			endIndex = median - 1
		}
	}

	return false
}
