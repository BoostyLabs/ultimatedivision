// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package mail

import (
	"strings"
)

// NormalizeEmail brings the email to UpperCase.
func NormalizeEmail(email string) string {
	return strings.ToUpper(email)
}
