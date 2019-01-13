package amount

import (
	"fmt"

	"github.com/astromechza/coinagelib/internal/decext"

	"github.com/astromechza/coinagelib/core/commodity"

	"github.com/ericlagergren/decimal"
)

// Amount is a value that is tagged with a currency or commodity name.
// The value can be nil if it needs to be assigned by another system.
type Amount struct {
	Value     *decimal.Big
	Commodity commodity.Commodity
}

func (a Amount) IsEmpty() bool {
	return a.Value == nil
}

func (a Amount) NotEmpty() bool {
	return !a.IsEmpty()
}

func (a Amount) IsZero() bool {
	return a.NotEmpty() && decext.IsZero(a.Value)
}

func (a Amount) IsNotZero() bool {
	return a.IsEmpty() || !decext.IsZero(a.Value)
}

func (a Amount) IsValid(allowEmpty bool) (ok bool, err error) {
	if ok, err := a.Commodity.IsValid(); !ok {
		return false, fmt.Errorf("invalid commodity: %s", err)
	}
	if a.IsEmpty() {
		if !allowEmpty {
			return false, fmt.Errorf("value is nil")
		}
	} else if !a.Value.IsFinite() {
		return false, fmt.Errorf("value %s is not finite", a.Value)
	}
	return true, nil
}

func (a Amount) Copy() *Amount {
	return New(a.Commodity, decext.Copy(a.Value))
}

func New(c commodity.Commodity, v *decimal.Big) *Amount {
	if v == nil || !v.IsFinite() {
		panic(fmt.Errorf("refusing to create non finite amount %s", v))
	}
	return &Amount{
		Commodity: c,
		Value:     v,
	}
}

func NewZero(c commodity.Commodity) *Amount {
	return &Amount{
		Commodity: c,
		Value:     new(decimal.Big),
	}
}

func NewEmpty(c commodity.Commodity) *Amount {
	return &Amount{
		Commodity: c,
		Value:     nil,
	}
}
