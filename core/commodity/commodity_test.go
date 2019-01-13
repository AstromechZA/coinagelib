package commodity

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/astromechza/coinagelib/internal/assert"
)

func TestCommodity_IsValid_good(t *testing.T) {
	for i, c := range []Commodity{
		// currency symbols
		"£", "$", "R", "kn", "Kč", "₿", "₨",
		// stock symbols
		"ORCL", "GOOG",
		// units
		"h", "g",
		// long
		"丏丏丏丏丏丏丏丏丏丏",
	} {
		t.Run(fmt.Sprintf("%02d[%s]", i+1, c), func(t *testing.T) {
			ok, err := c.IsValid()
			assert.ShouldEqual(t, ok, true)
			assert.ShouldEqual(t, err, nil)
		})
	}
}

func TestCommodity_IsValid_bad(t *testing.T) {
	for i, c := range []struct {
		c        Commodity
		expected error
	}{
		// length checks
		{"", fmt.Errorf("is shorter than 1 characters")},
		{"thisismuchtoolong", fmt.Errorf("is longer than 10 characters")},
		// bad characters
		{"no space", fmt.Errorf("contains bad character ` ` at position 2")},
		// separator
		{"-", fmt.Errorf("contains bad character `-` at position 0")},
	} {
		t.Run(fmt.Sprintf("%02d[%s]", i+1, c.c), func(t *testing.T) {
			ok, err := c.c.IsValid()
			assert.ShouldEqual(t, ok, false)
			if assert.ShouldNotEqual(t, err, nil) {
				assert.Equal(t, err, c.expected)
			}
		})
	}
}

func TestCommodity_json_marshal(t *testing.T) {
	x, err := json.Marshal([]Commodity{Commodity("₿")})
	assert.Equal(t, err, nil)
	assert.Equal(t, string(x), `["₿"]`)
}

func TestCommodity_unmarshal(t *testing.T) {
	var c struct {
		C Commodity
	}
	err := json.Unmarshal([]byte(`{"C": "£"}`), &c)
	assert.Equal(t, err, nil)
	assert.Equal(t, c.C, Commodity("£"))
}
