package utils

import (
	"bytes"
	"regexp"
)

var (
	patEols  = regexp.MustCompile(`[\r\n]+`)
	pat2Eols = regexp.MustCompile(`[\r\n]{2}`)
)

// Modified version of Go's builtin bufio.ScanLines to return strings separated by
// two newlines (instead of one). Returns a string without newlines in it, and trims
// spaces from start and end.
// https://github.com/golang/go/blob/master/src/bufio/scan.go#L344-L364
func ScanTwoConsecutiveNewlines(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	if loc := pat2Eols.FindIndex(data); loc != nil && loc[0] >= 0 {
		// Replace newlines within string with a space
		s := patEols.ReplaceAll(data[0:loc[0]+1], []byte(" "))
		// Trim spaces and newlines from string
		s = bytes.Trim(s, "\n ")
		return loc[1], s, nil
	}

	if atEOF {
		// Replace newlines within string with a space
		s := patEols.ReplaceAll(data, []byte(" "))
		// Trim spaces and newlines from string
		s = bytes.Trim(s, "\r\n ")
		return len(data), s, nil
	}

	// Request more data.
	return 0, nil, nil
}
