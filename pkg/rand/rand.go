// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package rand

import (
	"math/rand"
	"time"
)

// Minute generates minute in given interval.
func Minute(begin, end int) int {
	rand.Seed(time.Now().UnixNano())
	minute := begin + rand.Intn(end-begin+1)

	return minute
}
