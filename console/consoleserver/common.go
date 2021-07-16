// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package consoleserver

import (
	"fmt"
	"strings"
)

// ValidateKey checks the validity of the key filter.
func ValidateKey(key string) error {
	for _, v := range Keys {
		if key == v {
			return nil
		}
	}
	return fmt.Errorf(key + " key does not exist")
}

// Keys defines the list of possible card key filters.
var Keys = []string{
	"dominant_foot",
	"tactics",
	"physique",
	"player_name",
}

// ValidateValue returns a copy of the string without of invalid UTF-8 bytes.
func ValidateValue(value string) string {
	return strings.ToValidUTF8(value, "")
}
