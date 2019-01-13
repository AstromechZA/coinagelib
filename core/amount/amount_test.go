package amount

import (
	"fmt"
	"testing"

	"github.com/astromechza/coinagelib/internal/decext"

	"github.com/astromechza/coinagelib/assert"

	"github.com/ericlagergren/decimal"
)

func TestAmount_IsEmpty(t *testing.T) {
	assert.ShouldBeTrue(t, Amount{}.IsEmpty())
	assert.ShouldBeFalse(t, Amount{Value: new(decimal.Big)}.IsEmpty())
}

func TestAmount_NotEmpty(t *testing.T) {
	assert.ShouldBeTrue(t, Amount{Value: new(decimal.Big)}.NotEmpty())
	assert.ShouldBeFalse(t, Amount{}.NotEmpty())
}

func TestAmount_IsZero(t *testing.T) {
	assert.ShouldBeTrue(t, Amount{Value: new(decimal.Big)}.IsZero())
	assert.ShouldBeFalse(t, Amount{}.IsZero())
	assert.ShouldBeFalse(t, Amount{Value: decimal.New(1, 1)}.IsZero())
}

func TestAmount_IsNotZero(t *testing.T) {
	assert.ShouldBeTrue(t, Amount{}.IsNotZero())
	assert.ShouldBeTrue(t, Amount{Value: decimal.New(42, 0)}.IsNotZero())
	assert.ShouldBeFalse(t, Amount{Value: new(decimal.Big)}.IsNotZero())
}

func TestAmount_IsValid(t *testing.T) {
	for i, c := range []struct {
		amount     *Amount
		allowEmpty bool
		ok         bool
		err        error
	}{
		{
			&Amount{Value: new(decimal.Big), Commodity: "$"},
			false,
			true,
			nil,
		},
		{
			&Amount{Commodity: "$"},
			false,
			false,
			fmt.Errorf("value is nil"),
		},
		{
			&Amount{Commodity: "$"},
			true,
			true,
			nil,
		},
		{
			&Amount{Value: new(decimal.Big), Commodity: "has space"},
			false,
			false,
			fmt.Errorf("invalid commodity: contains bad character ` ` at position 3"),
		},
		{
			&Amount{Value: decext.Inf(), Commodity: "£"},
			false,
			false,
			fmt.Errorf("value Infinity is not finite"),
		},
	} {
		t.Run(fmt.Sprintf("%02d", i+1), func(t *testing.T) {
			o, e := c.amount.IsValid(c.allowEmpty)
			assert.ShouldEqual(t, o, c.ok)
			assert.ShouldEqual(t, e, c.err)
		})
	}
}

func TestNew(t *testing.T) {
	a := New("$", decimal.New(42, 1))
	assert.ShouldEqual(t, string(a.Commodity), "$")
	assert.ShouldEqual(t, a.Value.String(), "4.2")
}

func TestNewZero(t *testing.T) {
	a := NewZero("$")
	assert.ShouldEqual(t, string(a.Commodity), "$")
	assert.ShouldBeTrue(t, a.IsZero())
}

func TestNewEmpty(t *testing.T) {
	a := NewEmpty("$")
	assert.ShouldEqual(t, string(a.Commodity), "$")
	assert.ShouldBeTrue(t, a.IsEmpty())
}

func TestCopy(t *testing.T) {
	a := New("$", decimal.New(42, 1))
	b := a.Copy()
	assert.Equal(t, a, b)
	a.Value.Add(a.Value, decimal.New(5, 0))
	assert.NotEqual(t, a, b)
}
