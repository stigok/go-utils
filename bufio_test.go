package utils

import (
	"bufio"
	"strings"
	"testing"
)

func TestScan2Lines(t *testing.T) {
	s := `hello my
world
!

foo
bar baz
broken

baz`
	var lines []string
	scanner := bufio.NewScanner(strings.NewReader(s))
	scanner.Split(ScanTwoConsecutiveNewlines)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if lines[0] != "hello my world !" {
		t.Errorf("wanted 'hello world !', got '%s'", lines[0])
	}
	if lines[1] != "foo bar baz broken" {
		t.Errorf("wanted 'foo bar', got '%s'", lines[1])
	}
	if lines[2] != "baz" {
		t.Errorf("wanted 'baz', got '%s'", lines[2])
	}
}
