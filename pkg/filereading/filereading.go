// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package filereading

import (
	"bufio"
	"io"
	"os"
	"strings"
)

// ReadCertainLine reads certain line from the file.
func ReadCertainLine(file *os.File, lineNum int) (string, error) {


	var line int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line++

		if line == lineNum {
			return strings.TrimSpace(scanner.Text()), scanner.Err()
		}
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}

	return "", io.EOF
}

// CountLines counts lines from file.
func CountLines(r io.Reader) (int, error) {
	var line int
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line++
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return line, nil
}
