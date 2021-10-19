// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package search_test

import (
	"testing"

	"ultimatedivision/pkg/search"
)

func TestBinarySearch(t *testing.T) {
	testData := []struct {
		dataset    []string
		keyElement string
		result     bool
	}{
		{
			[]string{"a", "b", "c", "d"},
			"a",
			true,
		},
		{
			[]string{"a", "b", "c", "d"},
			"e",
			false,
		},
	}

	for _, value := range testData {
		result := search.BinarySearch(value.dataset, value.keyElement)
		if result != value.result {
			t.Error(
				"For", value.dataset, "with key element", value.keyElement,
				"expected", value.result, "got", result,
			)
		}
	}
}
