package decext

import (
	"github.com/ericlagergren/decimal"
)

func Copy(d *decimal.Big) *decimal.Big {
	return new(decimal.Big).Set(d)
}

func IsZero(d *decimal.Big) bool {
	return d != nil && d.CmpAbs(new(decimal.Big)) == 0
}

func Inf() *decimal.Big {
	return new(decimal.Big).SetInf(false)
}
