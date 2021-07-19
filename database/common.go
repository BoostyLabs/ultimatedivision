// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package database

import (
	"strings"
)

// Action defines the list of possible filter actions.
type Action string

const (
	// EQ - equal to value.
	EQ Action = "="
	// GTE - greater than or equal to value.
	GTE Action = ">="
	// LTE - less than or equal to value.
	LTE Action = "<="
	// LIKE - like to value.
	LIKE Action = "LIKE"
)

// normalizeEmail brings the email to UpperCase.
func normalizeEmail(email string) string {
	return strings.ToUpper(email)
}
