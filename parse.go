package phpregexp

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

var (
	// ErrMalformedRegexp is returned when provided regexp string too short
	ErrMalformedRegexp          = errors.New("Malformed Regexp")
	// ErrCantFindClosingSeparator is returned when provided regexp string
	// has wrong open/close separator configuration
	ErrCantFindClosingSeparator = errors.New("Cant find closing separator")
)

// MustCompile is like Compile but panics if the expression cannot be parsed.
// It simplifies safe initialization of global variables holding compiled regular
// expressions.
func MustCompile(str string) *regexp.Regexp {
	regexp, err := Compile(str)
	if err != nil {
		panic(`regexp: Compile(` + quote(str) + `): ` + err.Error())
	}

	return regexp
}

// Compile parses a regular expression in PHP format and returns,
// if successful, a Regexp object that can be used to match against text.
func Compile(str string) (*regexp.Regexp, error) {
	if len(str) < 3 {
		return nil, ErrMalformedRegexp
	}

	// Searching for opening and closing separator
	sep := str[0:1]
	last := strings.LastIndex(str, sep)
	if last == 0 || strings.Count(str, sep) != 2 {
		return nil, ErrCantFindClosingSeparator
	}

	// Reading modifiers
	prefix := ""
	mods := str[last+1:]
	if len(mods) > 0 {
		for _, m := range mods {
			switch m {
			case 'i': // Case insensitive
				prefix += "i"
			case 'm': // Multiline
				prefix += "m"
			case 's': // Single line
				prefix += "s"
			case 'U': // Ungreedy
				prefix += "U"
			}
		}

		if prefix != "" {
			prefix = "(?" + prefix + ")"
		}
	}

	return regexp.Compile(prefix + str[1:last])
}

func quote(s string) string {
	if strconv.CanBackquote(s) {
		return "`" + s + "`"
	}
	return strconv.Quote(s)
}
