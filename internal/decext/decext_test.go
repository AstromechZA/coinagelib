package decext

import (
	"testing"

	"github.com/astromechza/coinagelib/internal/assert"
	"github.com/ericlagergren/decimal"
)

func TestCopy(t *testing.T) {
	a := decimal.New(1, 0)
	b := Copy(a)
	a.Add(a, a)
	assert.Equal(t, a.String(), "2")
	assert.Equal(t, b.String(), "1")
}

func TestIsZero(t *testing.T) {
	assert.ShouldBeFalse(t, IsZero(decimal.New(1, 0)))
	assert.ShouldBeFalse(t, IsZero(nil))
	assert.ShouldBeTrue(t, IsZero(decimal.New(0, 0)))
}

func TestInf(t *testing.T) {
	assert.True(t, Inf().IsInf(0))
}
