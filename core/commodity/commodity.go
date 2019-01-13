package commodity

import (
	"fmt"
	"regexp"
	"unicode/utf8"
)

// Commodity represents a currency, stock, or unit of measurement
type Commodity string

const (
	// CharacterSet represents the valid characters that can be used
	CharacterSet = `\p{Sc}\p{L}`
	// BadCharacterSet can be used to search for bad characters
	BadCharacterSet = `[^` + CharacterSet + `]`

	MinCommodityLen = 1
	MaxCommodityLen = 10
)

// IsValid checks the validity of the commodity name and returns any error it finds
func (c Commodity) IsValid() (ok bool, err error) {
	rl := utf8.RuneCountInString(string(c))
	if rl < MinCommodityLen {
		return false, fmt.Errorf("is shorter than %d characters", MinCommodityLen)
	} else if rl > MaxCommodityLen {
		return false, fmt.Errorf("is longer than %d characters", MaxCommodityLen)
	}

	if m := regexp.MustCompile(BadCharacterSet).FindStringIndex(string(c)); m != nil {
		return false, fmt.Errorf("contains bad character `%s` at position %d", c[m[0]:m[1]], m[0])
	}
	return true, nil
}
