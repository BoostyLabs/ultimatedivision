// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package consoleserver

import (
	"fmt"
	"strings"

	"ultimatedivision/cards"
)

// ValidateKey checks the existence and validity of the key from the url request.
func ValidateKey(key string) error {
	for _, v := range cards.Keys {
		if key == v {
			return nil
		}
	}
	return fmt.Errorf(key + " key does not exist")
}

// ValidateValue returns a copy of the string without of invalid UTF-8 bytes.
func ValidateValue(value string) string {
	return strings.ToValidUTF8(value, "")
}
